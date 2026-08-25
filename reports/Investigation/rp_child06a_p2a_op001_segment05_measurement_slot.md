# Lane

Segment 05 direct-Owner-authorized OP001 child measurement, with Main correction from thread `01a036a0-38ab-73a3-9cca-02849eccf4bf`. Exact scope: `P2-A / B1-P1A-OP001 resolution / PARENT_DRILLDOWN_PENDING`. Boundary remained overlay/debug-artifact-only; this lane did not edit production source, tests, or Child 06A ledgers.

# Work

The accepted carried capture remains the source for the previously known process/analyzer/phase/parent values; it contained zero child timing rows. The overlay harness measured the required 17 exclusive children using the accepted Segment 03 source map, canonical LadybugDB runtime build contract, exact Go executable `C:\Program Files\Go\bin\go.exe` (`go version go1.26.3 windows/amd64`), and artifact-local process data.

The 09:20 run was terminated by Main after Owner authorization and is preserved as non-benchmark evidence under `E:\Anvien\.tmp\child06a_p2a_op001_resolution_slot_segment05\main-interrupted-20260825-0920-non-benchmark`. No partial number from it was reused. After confirming the process tree was gone, the replacement used:

`pwsh -NoLogo -NoProfile -NonInteractive -ExecutionPolicy Bypass -File "E:\Anvien\.tmp\child06a_p2a_op001_resolution_slot_segment05\run.ps1" -GoExe "C:\Program Files\Go\bin\go.exe" -Execute`

The replacement analyze exited `0`. `postprocess.ps1` then validated and materialized the child table without another build or analyze.

# Result

Status: `MEASUREMENT_COMPLETE_WITH_SEMANTIC_EQUIVALENCE_WARNING`.

Same-run process wall `653.475797100 s`; process CPU `830.109375000 s`; analyzer `630.160832200 s`; `analyzer_internal/resolution` parent `545.434182000 s` with denominator `{"runs":1}`. Child sum `545.364656400 s`; residual `0.069525600 s`. The measurement contains `2309` exclusive intervals and `0` overlaps.

| rank | child_id | duration_s | share_parent_pct | denominator | source file:function:start-end | full call path |
|---:|---|---:|---:|---|---|---|
| 1 | resolve_calls | 243.794553800 | 44.697337 | calls=72976; files=765 | E:\Anvien\internal\resolution\resolve.go:resolveCall:92-94 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each call in ir.Calls -> resolution.resolveCall |
| 2 | resolve_accesses | 204.093359700 | 37.418513 | accesses=43826; files=765 | E:\Anvien\internal\resolution\resolve.go:resolveAccess:95-97 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each access in ir.Accesses -> resolution.resolveAccess |
| 3 | resolve_type_annotations | 92.316121800 | 16.925254 | files=765; typeAnnotations=37389 | E:\Anvien\internal\resolution\resolve.go:resolveTypeAnnotation:98-100 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each annotation in ir.TypeAnnotations -> resolution.resolveTypeAnnotation |
| 4 | project_resolution_outcomes | 4.067254700 | 0.745691 | graphNodes=67882; graphRelationships=106991; outcomes=150404; referencesBySourceScope=46600; referencesByTargetDef=46600; workUnits=418477 | E:\Anvien\internal\resolution\outcome.go:projectResolutionOutcomes:114-116 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.projectResolutionOutcomes |
| 5 | emit_definition_nodes | 0.387684000 | 0.071078 | definitions=46608; exports=417; files=765; runs=1; workUnits=47790 | E:\Anvien\internal\resolution\emit.go:emitDefinitionNodes:77-81 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitDefinitionNodes |
| 6 | finalize_resolution_outcomes | 0.328439300 | 0.060216 | outcomeMapEntries=150404 | E:\Anvien\internal\resolution\outcome.go:(*resolutionOutcomeCollector).finalize:110-113 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> (*resolutionOutcomeCollector).finalize |
| 7 | build_binding_occurrence_index | 0.172202300 | 0.031572 | bindingLeaves=2674; definitionsVisited=46608; filePasses=1530; files=765; ownedDefIDsVisited=46608; ownerBindingsInspected=30782; runs=1; scopesVisited=11368; workUnits=139570 | E:\Anvien\internal\resolution\resolve.go:buildBindingOccurrenceIndex:62-65 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.buildBindingOccurrenceIndex |
| 8 | finalize_typescript_authority_results | 0.096993100 | 0.017783 | authorityResults=5704 | E:\Anvien\internal\resolution\resolve.go:finalizeTypeScriptAuthorityResults:103-106 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.finalizeTypeScriptAuthorityResults |
| 9 | emit_import_edges | 0.088903700 | 0.016300 | imports=4887; targetDefinitions=1381; targetFiles=5562; workUnits=11830 | E:\Anvien\internal\resolution\emit.go:emitImportEdges:83-83 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitImportEdges |
| 10 | emit_typescript_external_symbols | 0.011376900 | 0.002086 | authorityResults=5495; resolvedRecords=0; uniqueResolvedSymbols=0; workUnits=5495 | E:\Anvien\internal\resolution\external_symbol.go:emitTypeScriptExternalSymbols:107-109 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitTypeScriptExternalSymbols |
| 11 | emit_method_dispatch_edges | 0.007767100 | 0.001424 | heritageFacts=16; memberEntries=6510; ownerMemberOwners=1075 | E:\Anvien\internal\resolution\emit.go:emitMethodDispatchEdges:102-102 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitMethodDispatchEdges |
| 12 | assemble_resolution_result | 0.000000000 | 0.000000 | authorityResults=5495; graphPointers=1; metricsValues=1; outcomes=150404; referenceIndexes=1; resultAssemblies=1 | E:\Anvien\internal\resolution\resolve.go:ResolveBoundInto:121-127 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> construct resolution.Result return value |
| 13 | binding_accumulator_dispose | 0.000000000 | 0.000000 | accumulatedEntries=904; deferredExecutions=1; fileEntryBuckets=139; fileScopeBuckets=139; workUnits=1182 | E:\Anvien\internal\resolution\binding_accumulator.go:(*bindingAccumulator).dispose:71-74 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto deferred closure -> (*bindingAccumulator).dispose |
| 14 | emit_heritage_compatibility_edges | 0.000000000 | 0.000000 | heritageFacts=16; runs=1 | E:\Anvien\internal\resolution\emit.go:emitHeritageCompatibilityEdges:85-89 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each w.heritage -> resolution.emitHeritageCompatibilityEdges |
| 15 | emit_unresolved_heritage_diagnostics | 0.000000000 | 0.000000 | runs=1; unresolvedHeritageFacts=45 | E:\Anvien\internal\resolution\resolve.go:emitUnresolvedHeritageDiagnostics:82-82 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> resolution.emitUnresolvedHeritageDiagnostics |
| 16 | finalize_resolution_metadata | 0.000000000 | 0.000000 | graphNodes=67882; graphRelationships=106991; metadataUpdates=1; workUnits=174874 | E:\Anvien\internal\resolution\resolve.go:ResolveBoundInto:118-120 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> graphhealth.SetResolutionMetadata |
| 17 | runtime_setup | 0.000000000 | 0.000000 | graphAllocations=0; newEmitterInvocations=1; resolveBoundIntoInvocations=1 | E:\Anvien\internal\resolution\resolve.go:ResolveBoundInto:66-75 | internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> conditional graph.New -> resolution.newEmitter |

Largest measured child: `resolve_calls`, owner `E:\Anvien\internal\resolution\resolve.go:resolveCall:92-94`, processing `72976` calls across `765` files. Its source-owned work adds call/receiver-read references and graph relationships, records repository/intrinsic/TypeScript authority outcomes, appends authority-site/diagnostic state, and updates resolution metrics. Full call path: `internal/analyze.Run -> internal/analyze.runPhase(PhaseResolution) -> resolution.ResolveBoundInto -> for each ir in w.files -> for each call in ir.Calls -> resolution.resolveCall`.

Semantic equivalence against the carried capture failed only for `GraphNodesEmitted` (`67882` versus `67783`, `+99`) and `GraphRelationshipsEmitted` (`106991` versus `106892`, `+99`). Therefore this capture is not before/after speed evidence; the same-run child durations, ranking, interval proof, and arithmetic are directly measured and retained.

Raw artifacts: `benchmark.json`, `process.json`, `resolution.cpu.pprof`, `child_metrics.json`, `measurement_validation.json`, `bottleneck_ranking.json`, and `bottleneck_ranking.md` under `E:\Anvien\.tmp\child06a_p2a_op001_resolution_slot_segment05`.

# Checkpoint

Replacement analyze exit code `0`; 17 measured child IDs are unique and in the authorized order; ranking is descending; duration sums and residual arithmetic match; CPU profile is non-empty. Production/test files were not edited by this lane. Scoped git status did change concurrently during the run: `actual-status.md` left the pre-run dirty list while the existing internal-source and three remaining ledger entries stayed dirty. The 09:20 interrupted run remains explicitly non-benchmark. Current P2-A state remains owner/Main-controlled pending ledger transition.

# Next Owner

Main correction thread `01a036a0-38ab-73a3-9cca-02849eccf4bf` and P2-A Main `01a0366e-1f78-77d2-bf10-eed74cae9673` own independent raw-evidence verification and immediate benchmark/plan/evidence/actual-status updates. The proved implementation candidate is the measured #1 `resolve_calls`; this lane makes no optimization recommendation and performs no further run.
