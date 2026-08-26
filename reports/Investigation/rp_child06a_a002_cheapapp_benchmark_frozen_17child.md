# Child 06A A002 Cheapapp Frozen 17-Child Benchmark

Status: MEASUREMENT_COMPLETE.

This is measurement data for E:\cheapapp.org only. It does not make a KEEP, Supervisor, acceptance, rejection, disposition, baseline-promotion, streak, checklist, or Restaurant Manager claim.

## Result

The one authorized frozen A002 candidate launch completed with exit 0 and produced a valid same-work packet: 30/30 top-level operations, 17/17 resolution children, D001 calls=27890/files=887, exactly 2,675 exclusive intervals, zero overlaps, a nonnegative parent residual, and a nonempty resolution CPU profile.

| Boundary | Accepted A001 before | A002 frozen candidate | Candidate - before |
|---|---:|---:|---:|
| D001 resolve_calls | 25.045225300 s | 3.090914200 s | -21.954311100 s (-87.658669%) |
| B1-P1A-OP001 resolution parent | 184.481061700 s | 19.040468000 s | -165.440593700 s (-89.678904%) |
| Analyzer total | 274.474620900 s | 100.843249000 s | -173.631371900 s (-63.259536%) |
| Process wall | 279.105934600 s | 136.729876000 s | -142.376058600 s (-51.011477%) |

These candidate values are measured and valid for handoff, but remain unaccepted until Main completes the rest of the campaign sequence.

## Exact launch and frozen identity

| Field | Exact value |
|---|---|
| Target | E:\cheapapp.org |
| Target HEAD | a869876ab6262dacde6cd5d432d099a91852a646 |
| Anvien HEAD | cc420e7ad719d90dc4b2d9991be0249e8d648daa |
| Frozen packet root | E:\Anvien\.tmp\child06a_a002_a00x_benchmark_build |
| Executable | E:\Anvien\.tmp\child06a_a002_a00x_benchmark_build\anvien-a002-benchmark.exe |
| Version / bytes / SHA-256 | 1.2.8 / 73,816,576 / 0F8B8244ABC80339A73E3A29F38D32F55A7A6A65281A64C291785ACF8F4A241E |
| Adjacent DLL | E:\Anvien\.tmp\child06a_a002_a00x_benchmark_build\lbug_shared.dll |
| DLL bytes / SHA-256 | 20,230,656 / 20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7 |
| Provenance | E:\Anvien\.tmp\child06a_a002_a00x_benchmark_build\anvien-a002-benchmark.provenance.json |
| Provenance bytes / SHA-256 | 19,488 / 673FCCFEA17C48640C0AC62243479FE5EFD71E172CB1D4136D046C806E5855F1 |
| Exact command | E:\Anvien\.tmp\child06a_a002_a00x_benchmark_build\anvien-a002-benchmark.exe analyze E:\cheapapp.org --force --skip-git --json --progress --benchmark-json E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826\candidate\benchmark.json --benchmark-label child06a-a002-cheapapp-frozen-17child-20260826 |
| Launch count / PID / exit | 1 / 13232 / 0 |
| Start / end | 2026-08-26T10:58:53.0369124+07:00 / 2026-08-26T11:01:09.5419385+07:00 |
| Process wall / CPU | 136.729876000 s / 105.125000000 s |
| User / kernel CPU | 90.984375000 s / 14.140625000 s |
| Competing analyze/benchmark processes before launch | 0 |
| Target graph location | E:\cheapapp.org\.anvien |

The provenance passed schema 1, attempt A002, exit 0, exact 2/2 overlay mappings, 3/3 A002 candidate sources, and 3/3 native inputs. The overlay manifest SHA-256 is 7B138FBA06B41CBD4C6709E6DF00C80C03E4015B615D1E42F787DE2A65D378F7. Mapped resolve.go/types.go hashes are 92A89227E1B1B1C159DE8BE77A1F060361C9058C4F83C78BFC32E9AB40DADEA9 and A8DDCB1C61C312D3D3B164B1DD02BFD91C640F1FEB8B6DB511B73FD923C802B8.

Current A002 hashes remained:

| Source | SHA-256 |
|---|---|
| internal/graphhealth/diagnostics.go | AEABA8A541D1C293DFA4FE411253A85A26B6AC0DD18047AFA6F9151B4932BEE8 |
| internal/resolution/emit.go | 73AE20E8A18AC25F8019E2A3223407621AE7ABC358DBAF4F058E1D9AC3136060 |
| internal/resolution/outcome.go | 02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E |
| internal/graphhealth/diagnostics_test.go | 58E8CAE2C4EBBF4672D2D338CC5E1419B7CECA9F7389B0E24F51BF2F975A7CEA |

The process-only HOME, USERPROFILE, TEMP, TMP, APPDATA, LOCALAPPDATA, XDG_CACHE_HOME, XDG_CONFIG_HOME, and XDG_DATA_HOME were all set beneath E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826. The profile path was E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826\candidate\resolution.cpu.pprof.

## Conservation and resources

| Conservation field | Accepted A001 before | A002 candidate |
|---|---:|---:|
| Resolution parent | 184.481061700 s | 19.040468000 s |
| 17-child sum | 184.422170400 s | 19.024756200 s |
| Parent residual | 0.058891300 s | 0.015711800 s |
| Exclusive intervals | 2,675 | 2,675 |
| Overlap count | 0 | 0 |

For every candidate child, durationNs equals the exact sum of its intervalEnd - intervalStart pairs. All durations and denominators are nonnegative, every overlapGroup is exclusive_resolution_parent, and child sum + residual equals the same-run parent exactly.

| Resource field | Accepted A001 before | A002 candidate | Delta |
|---|---:|---:|---:|
| startAllocBytes | 1,309,264 | 1,315,328 | +6,064 |
| endAllocBytes | 863,629,352 | 1,055,121,264 | +191,491,912 |
| maxObservedSys | 1,989,634,552 | 1,972,115,960 | -17,518,592 |

Structural source/Coder evidence, not a runtime allocation inference: retained appender state is O(T) map entries and slice headers; cache and graph share the normalized slice; no second retained diagnostic-object copy exists.

## Same-work and equivalence

| Field | Accepted A001 | A002 candidate | Result |
|---|---:|---:|---|
| Files scanned / parsed / failed | 1,359 / 887 / 0 | 1,359 / 887 / 0 | equal |
| Parser bytes | 5,430,581 | 5,430,581 | equal |
| Graph nodes / relationships | 95,762 / 129,716 | 95,762 / 129,716 | equal |
| Dependency edges / projection unresolved | 13,360 / 735 | 13,360 / 735 | equal |
| DB rows nodes / relationships | 95,762 / 129,716 | 95,762 / 129,716 | equal |
| Top-level operations | 30 | 30 | equal denominators |
| Resolution children | 17 | 17 | equal denominators |
| Resolution semantic counters, diagnostics, outcomes | accepted A001 object | candidate object | exact JSON equality excluding timers |
| canonical graph.json | 840,614,023 bytes; 8763AA2D089328F3D782A5F01E06178D171922E3A5CE96A4EE6707510976C920 | same | exact hash match |
| public stdout.json | 8,824 bytes; F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4 | same | exact hash match |
| stderr.log | 213 bytes; 0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC | same | exact hash match |
| Ladybug lbug | 597,950,464 bytes; DC88B68B26BE75954A8BEE85F4CD426F0074171E03AAAC6D55B49C18EDDE1AB7 | 597,946,368 bytes; 2FEC9F35F7662B45789865DC9EB40A97555AAA1A68DB243BBC3A8D9B778C5AF4 | raw-byte mismatch; not the governing semantic gate |
| meta.json | 219 bytes; 5D5227EB91053443D51A51E9EC15D85532C1EF49E7F701B45B24F81076D9F1D9 | 219 bytes; 2FC9B672D97543DF42BED79F7CF19F113A70E1AA3A0561A871DA263557A9D43E | raw-byte mismatch; not the governing semantic gate |

The complete target HEAD/status array and the complete Anvien HEAD/status/staged/source-hash boundary are stored in preflight.json and validation.json. Target HEAD/status were exactly unchanged before and after measurement. Anvien HEAD remained cc420e7ad719d90dc4b2d9991be0249e8d648daa, staged path count remained 0, and every four-file A002 hash remained unchanged. The only Anvien write after measurement is this authorized new report plus raw packet files under ignored .tmp.

## Complete 30-operation inventory

Shares use each same-run process wall. Ranks are descending duration in each 30-row run.

| Row | Boundary | Operation | A001 s | A002 s | Delta | Rank | Process share | Denominator |
|---|---|---|---:|---:|---:|---:|---:|---|
| B1-P1A-OP001 | analyzer_internal | resolution | 184.481061700 | 19.040468000 | -165.440593700 (-89.678904%) | 1 -> 3 | 66.097148% -> 13.925609% | {"runs":1} |
| B1-P1A-OP002 | analyzer_internal | db_load | 40.950769400 | 39.621102900 | -1.329666500 (-3.246988%) | 2 -> 1 | 14.672124% -> 28.977649% | {"runs":1} |
| B1-P1A-OP003 | analyzer_internal | parse | 12.840209000 | 9.266697100 | -3.573511900 (-27.830637%) | 4 -> 6 | 4.600479% -> 6.777375% | {"runs":1} |
| B1-P1A-OP004 | analyzer_internal | graph_snapshot | 18.124932500 | 12.521999500 | -5.602933000 (-30.912849%) | 3 -> 4 | 6.493926% -> 9.158203% | {"nodes":95762,"relationships":129716} |
| B1-P1A-OP005 | analyzer_internal | semantic_enrichment | 5.503521100 | 4.700809600 | -0.802711500 (-14.585417%) | 6 -> 7 | 1.971840% -> 3.438027% | {"runs":1} |
| B1-P1A-OP006 | analyzer_internal | db_runner_resolve | 1.488161300 | 0.861169000 | -0.626992300 (-42.132012%) | 9 -> 9 | 0.533189% -> 0.629832% | {"runners":1} |
| B1-P1A-OP007 | analyzer_internal | cross_file_binding | 0.809683000 | 0.567778800 | -0.241904200 (-29.876408%) | 10 -> 12 | 0.290099% -> 0.415256% | {"runs":1} |
| B1-P1A-OP008 | cli_outer | ai_context | 2.066776600 | 35.002946800 | +32.936170200 (+1593.600886%) | 7 -> 2 | 0.740499% -> 25.600072% | {"baseSkills":49,"generatedFiles":4} |
| B1-P1A-OP009 | analyzer_internal | analyzer_orchestration | 0.783394400 | 0.785403800 | +0.002009400 (+0.256499%) | 11 -> 10 | 0.280680% -> 0.574420% | {"analyzeRuns":1} |
| B1-P1A-OP010 | analyzer_internal | processes | 0.626072300 | 0.709669800 | +0.083597500 (+13.352691%) | 12 -> 11 | 0.224314% -> 0.519031% | {"runs":1} |
| B1-P1A-OP011 | analyzer_internal | scan | 5.869840300 | 10.412393200 | +4.542552900 (+77.388015%) | 5 -> 5 | 2.103087% -> 7.615302% | {"runs":1} |
| B1-P1A-OP012 | cli_outer | file_projection | 0.473783400 | 0.301935900 | -0.171847500 (-36.271321%) | 14 -> 15 | 0.169750% -> 0.220827% | {"files":1359,"nodes":95762,"relationships":129716} |
| B1-P1A-OP013 | analyzer_internal | routes | 1.867265600 | 1.280667900 | -0.586597700 (-31.414797%) | 8 -> 8 | 0.669017% -> 0.936641% | {"runs":1} |
| B1-P1A-OP014 | analyzer_internal | documents | 0.128885400 | 0.096750000 | -0.032135400 (-24.933313%) | 19 -> 18 | 0.046178% -> 0.070760% | {"runs":1} |
| B1-P1A-OP015 | analyzer_internal | db_runner_close | 0.253016400 | 0.400254600 | +0.147238200 (+58.193145%) | 16 -> 14 | 0.090652% -> 0.292734% | {"runners":1} |
| B1-P1A-OP016 | analyzer_outer | analyze_setup | 0.216351500 | 0.062211300 | -0.154140200 (-71.245265%) | 17 -> 19 | 0.077516% -> 0.045499% | {"analyzeRuns":1} |
| B1-P1A-OP017 | analyzer_internal | tools | 0.535691500 | 0.405545300 | -0.130146200 (-24.294991%) | 13 -> 13 | 0.191931% -> 0.296603% | {"runs":1} |
| B1-P1A-OP018 | analyzer_internal | communities | 0.039965300 | 0.037951100 | -0.002014200 (-5.039872%) | 21 -> 21 | 0.014319% -> 0.027756% | {"runs":1} |
| B1-P1A-OP019 | cli_outer | registry_meta | 0.355614000 | 0.110769600 | -0.244844400 (-68.851170%) | 15 -> 17 | 0.127412% -> 0.081013% | {"repositories":1} |
| B1-P1A-OP020 | cli_outer | cli_preparation | 0.056014100 | 0.047864000 | -0.008150100 (-14.550086%) | 20 -> 20 | 0.020069% -> 0.035006% | {"commands":1} |
| B1-P1A-OP021 | cli_outer | output_publication | 0.036380300 | 0.002035800 | -0.034344500 (-94.404114%) | 22 -> 26 | 0.013035% -> 0.001489% | {"outputs":1} |
| B1-P1A-OP022 | analyzer_internal | orm | 0.138500100 | 0.112688900 | -0.025811200 (-18.636232%) | 18 -> 16 | 0.049623% -> 0.082417% | {"runs":1} |
| B1-P1A-OP023 | cli_outer | cli_startup | 0.022272100 | 0.013315600 | -0.008956500 (-40.213990%) | 23 -> 22 | 0.007980% -> 0.009739% | {"commands":1} |
| B1-P1A-OP024 | analyzer_internal | structure | 0.012747400 | 0.010573600 | -0.002173800 (-17.052889%) | 24 -> 23 | 0.004567% -> 0.007733% | {"runs":1} |
| B1-P1A-OP025 | analyzer_internal | graph_compact | 0.012060400 | 0.009142000 | -0.002918400 (-24.198202%) | 25 -> 24 | 0.004321% -> 0.006686% | {"inputNodes":95762,"inputRelationships":129716} |
| B1-P1A-OP026 | analyzer_internal | mro | 0.000019700 | 0.002183900 | +0.002164200 (+10985.786802%) | 27 -> 25 | 0.000007% -> 0.001597% | {"runs":1} |
| B1-P1A-OP027 | analyzer_internal | benchmark_write | 0.008824100 | 0.000000000 | -0.008824100 (-100.000000%) | 26 -> 27 | 0.003162% -> 0.000000% | {"artifacts":1} |
| B1-P1A-OP028 | analyzer_internal | cobol | 0.000000000 | 0.000000000 | 0.000000000 (n/a) | 28 -> 28 | 0.000000% -> 0.000000% | {"runs":1} |
| B1-P1A-OP029 | cli_outer | memory_profile | 0.000000000 | 0.000000000 | 0.000000000 (n/a) | 29 -> 29 | 0.000000% -> 0.000000% | {"profiles":0} |
| B1-P1A-OP030 | cli_outer | cpu_profile_completion | 0.000000000 | 0.000000000 | 0.000000000 (n/a) | 30 -> 30 | 0.000000% -> 0.000000% | {"profiles":0} |

## Complete 17-child inventory

The raw benchmark recorder emits the 17 unique children in source/instrumentation order. comparison.json normalizes those exact rows into the required control-row order below; no child is added, omitted, or inferred.

| Row | Child | A001 s | A002 s | Delta | Rank | Parent share | Intervals | Denominator | Owner / complete call path |
|---|---|---:|---:|---:|---:|---:|---:|---|---|
| B2-P2A-A001-D001 | resolve_calls | 25.045225300 | 3.090914200 | -21.954311100 (-87.658669%) | 3 -> 3 | 13.576041% -> 16.233394% | 887 | {"calls":27890,"files":887} | E:\Anvien\internal\resolution\resolve.go:resolveCall:92-94; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each call in ir.Calls -> resolution.resolveCall |
| B2-P2A-A001-D002 | resolve_accesses | 91.241633200 | 8.894951800 | -82.346681400 (-90.251214%) | 1 -> 1 | 49.458536% -> 46.716036% | 887 | {"accesses":26042,"files":887} | E:\Anvien\internal\resolution\resolve.go:resolveAccess:95-97; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each access in ir.Accesses -> resolution.resolveAccess |
| B2-P2A-A001-D003 | resolve_type_annotations | 61.585096500 | 2.084835900 | -59.500260600 (-96.614707%) | 2 -> 4 | 33.382883% -> 10.949499% | 887 | {"files":887,"typeAnnotations":35003} | E:\Anvien\internal\resolution\resolve.go:resolveTypeAnnotation:98-100; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each annotation in ir.TypeAnnotations -> resolution.resolveTypeAnnotation |
| B2-P2A-A001-D004 | project_resolution_outcomes | 5.375049800 | 4.252390200 | -1.122659600 (-20.886497%) | 4 -> 2 | 2.913605% -> 22.333433% | 1 | {"graphNodes":36575,"graphRelationships":65045,"outcomes":86742,"referencesBySourceScope":25062,"referencesByTargetDef":25062,"workUnits":238486} | E:\Anvien\internal\resolution\outcome.go:projectResolutionOutcomes:114-116; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.projectResolutionOutcomes |
| B2-P2A-A001-D005 | emit_definition_nodes | 0.452126400 | 0.347973200 | -0.104153200 (-23.036301%) | 5 -> 5 | 0.245080% -> 1.827545% | 1 | {"definitions":25811,"exports":3539,"files":887,"runs":1,"workUnits":30237} | E:\Anvien\internal\resolution\emit.go:emitDefinitionNodes:77-81; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitDefinitionNodes |
| B2-P2A-A001-D006 | finalize_resolution_outcomes | 0.224157600 | 0.149869200 | -0.074288400 (-33.141147%) | 7 -> 6 | 0.121507% -> 0.787109% | 1 | {"outcomeMapEntries":86742} | E:\Anvien\internal\resolution\outcome.go:(*resolutionOutcomeCollector).finalize:110-113; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> (*resolutionOutcomeCollector).finalize |
| B2-P2A-A001-D007 | build_binding_occurrence_index | 0.075494100 | 0.081794600 | +0.006300500 (+8.345685%) | 8 -> 7 | 0.040922% -> 0.429583% | 1 | {"bindingLeaves":6922,"definitionsVisited":25811,"filePasses":1774,"files":887,"ownedDefIDsVisited":25811,"ownerBindingsInspected":30749,"runs":1,"scopesVisited":6502,"workUnits":97569} | E:\Anvien\internal\resolution\resolve.go:buildBindingOccurrenceIndex:62-65; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.buildBindingOccurrenceIndex |
| B2-P2A-A001-D008 | finalize_typescript_authority_results | 0.040710200 | 0.027542600 | -0.013167600 (-32.344720%) | 9 -> 9 | 0.022067% -> 0.144653% | 1 | {"authorityResults":23110} | E:\Anvien\internal\resolution\resolve.go:finalizeTypeScriptAuthorityResults:103-106; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.finalizeTypeScriptAuthorityResults |
| B2-P2A-A001-D009 | emit_import_edges | 0.348401700 | 0.079947000 | -0.268454700 (-77.053212%) | 6 -> 8 | 0.188855% -> 0.419879% | 1 | {"imports":6141,"targetDefinitions":4717,"targetFiles":4985,"workUnits":15843} | E:\Anvien\internal\resolution\emit.go:emitImportEdges:83; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitImportEdges |
| B2-P2A-A001-D010 | emit_typescript_external_symbols | 0.028817500 | 0.012664900 | -0.016152600 (-56.051358%) | 10 -> 10 | 0.015621% -> 0.066516% | 1 | {"authorityResults":21018,"resolvedRecords":0,"uniqueResolvedSymbols":0,"workUnits":21018} | E:\Anvien\internal\resolution\external_symbol.go:emitTypeScriptExternalSymbols:107-109; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitTypeScriptExternalSymbols |
| B2-P2A-A001-D011 | emit_method_dispatch_edges | 0.004458600 | 0.001872600 | -0.002586000 (-58.000269%) | 11 -> 11 | 0.002417% -> 0.009835% | 1 | {"heritageFacts":0,"memberEntries":7485,"ownerMemberOwners":1548} | E:\Anvien\internal\resolution\emit.go:emitMethodDispatchEdges:102; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitMethodDispatchEdges |
| B2-P2A-A001-D012 | assemble_resolution_result | 0.000000000 | 0.000000000 | 0.000000000 (n/a) | 13 -> 12 | 0.000000% -> 0.000000% | 1 | {"authorityResults":21018,"graphPointers":1,"metricsValues":1,"outcomes":86742,"referenceIndexes":1,"resultAssemblies":1} | E:\Anvien\internal\resolution\resolve.go:ResolveBoundInto:121-127; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> construct resolution.Result return value |
| B2-P2A-A001-D013 | binding_accumulator_dispose | 0.000000000 | 0.000000000 | 0.000000000 (n/a) | 14 -> 13 | 0.000000% -> 0.000000% | 1 | {"accumulatedEntries":7330,"deferredExecutions":1,"fileEntryBuckets":508,"fileScopeBuckets":508,"workUnits":8346} | E:\Anvien\internal\resolution\binding_accumulator.go:(*bindingAccumulator).dispose:71-74; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto deferred closure -> (*bindingAccumulator).dispose |
| B2-P2A-A001-D014 | emit_heritage_compatibility_edges | 0.000000000 | 0.000000000 | 0.000000000 (n/a) | 15 -> 14 | 0.000000% -> 0.000000% | 1 | {"heritageFacts":0,"runs":1} | E:\Anvien\internal\resolution\emit.go:emitHeritageCompatibilityEdges:85-89; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each w.heritage -> resolution.emitHeritageCompatibilityEdges |
| B2-P2A-A001-D015 | emit_unresolved_heritage_diagnostics | 0.000999500 | 0.000000000 | -0.000999500 (-100.000000%) | 12 -> 15 | 0.000542% -> 0.000000% | 1 | {"runs":1,"unresolvedHeritageFacts":7} | E:\Anvien\internal\resolution\resolve.go:emitUnresolvedHeritageDiagnostics:82; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitUnresolvedHeritageDiagnostics |
| B2-P2A-A001-D016 | finalize_resolution_metadata | 0.000000000 | 0.000000000 | 0.000000000 (n/a) | 16 -> 16 | 0.000000% -> 0.000000% | 1 | {"graphNodes":36575,"graphRelationships":65045,"metadataUpdates":1,"workUnits":101621} | E:\Anvien\internal\resolution\resolve.go:ResolveBoundInto:118-120; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> graphhealth.SetResolutionMetadata |
| B2-P2A-A001-D017 | runtime_setup | 0.000000000 | 0.000000000 | 0.000000000 (n/a) | 17 -> 17 | 0.000000% -> 0.000000% | 1 | {"graphAllocations":0,"newEmitterInvocations":1,"resolveBoundIntoInvocations":1} | E:\Anvien\internal\resolution\resolve.go:ResolveBoundInto:66-75; internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> conditional graph.New -> resolution.newEmitter |

## Raw packet

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826\preflight.json | 14,246 | 107F25075EE71A5A1491C3FCFDE6B66FF8DFA006AB415C25830AE621016A4080 |
| E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826\validation.json | 18,807 | B5DA04AFA1710A469EBE031DF569A679ADD733566CAF19BFE86E18436C45CA22 |
| E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826\comparison.json | 40,638 | 19E5FB16517224ED04F46357851D6627A6B176CA8A65BDE38C6742775AD3EDF9 |
| E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826\output-manifest.json | 2,220 | 5CA00714346045D224E3220BD11CE27EBB40131190BFB3144E06FD9F67A5C424 |
| E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826\candidate\benchmark.json | 149,351 | F265E54F261D8B2163BE4ED15BE4CC18472F3E9D9059B7512C61CB516E711F83 |
| E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826\candidate\process.json | 2,239 | 2ECCB7E2AC9C5984EFFD34F8B46AE4265571E88F68A41FAF473F68E7ED1735AE |
| E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826\candidate\stdout.json | 8,824 | F52193AD16DD18C3DBD78CC161B503CF4091BDFAC1B363962FEBC2C50AD7FEF4 |
| E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826\candidate\stderr.log | 213 | 0FDCF15222C02C25771C98E17D9220C89FDCD0294137A6AE4D9B3AAFD4761ADC |
| E:\Anvien\.tmp\child06a_a002_cheapapp_frozen_17child_20260826\candidate\resolution.cpu.pprof | 61,264 | 13D6D2EDA4DDF1FF26EBC91BF58945F07FCFC2D5E92C80135D95ACEF02326DA7 |

comparison.json contains the full machine-readable 30-operation and 17-child before/candidate/delta/rank/share/denominator inventories. validation.json contains per-child required-field checks, interval-sum checks, exact conservation, denominator equality, semantic/workload/output equality, and complete pre/post source/status proof.

No build, binary copy, baseline rerun, retry, source/test/script/ledger edit, Restaurant Manager access, Supervisor, disposition, stage, commit, cleanup, install, network, or subagent action occurred.
