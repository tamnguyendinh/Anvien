# Analyze / graph-generation performance root-cause investigation

- Status: **COMPLETE — evidence handoff; no implementation or architecture verdict**
- Measurement date: 2026-08-23 (Asia/Bangkok)
- Repository and only measured target: `E:\Anvien`
- Evidence root: `E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d`
- Scope owner: Root Cause lane — measurement, causal diagnosis, and source attribution only

## Executive determination

The current controlled runtime spends **533,746.9551 ms (89.469% of benchmark wall time)** in resolution. The evidence does not stop at that phase label.

### VERIFIED CAUSE

The primary current-runtime bottleneck is a repeated algorithmic call path that scans the complete `workspace.imports` slice for per-fact import ownership checks and normalizes every visited import path again:

`ResolveBoundInto` → per-file calls/accesses/type annotations → `repositoryReceiverClaimed` and/or direct import checks → `explicitImportNameClaimed` (and, for calls, `explicitImportCallState`) → `cleanPath` → `strings.ReplaceAll` + `filepath.Clean` + `filepath.ToSlash`.

The CPU, allocation, call-graph, source, and workload evidence converge:

- `explicitImportNameClaimed`: **39.43 s flat / 319.95 s cumulative**, or **4.80% / 38.95%** of all CPU-profile samples.
- Its `cleanPath` child: **280.15 s cumulative**.
- `explicitImportCallState`, a separate full-import scan: **3.90 s flat / 47.51 s cumulative**, including **43.45 s** in `cleanPath`.
- `cleanPath` across the run: **0.88 s flat / 324.75 s cumulative**, **39.54%** of samples; **149,590.52 MB (83.59%)** of profile-estimated allocated bytes and **3,056,971,217 objects (91.54%)**.
- The stored import path is already normalized during workspace construction, but the hot loops normalize it again on every visit.

For misses, the source-level shape is a full scan per query; therefore the mechanism is `O(Q × I)` over import-claim queries `Q` and stored imports `I`, with early exit only on a match. The present artifacts do not expose exact `Q`, exact `I`, or loop-iteration counts, so no fabricated multiplicity is reported.

### VERIFIED SECONDARY CURRENT COST

Structured unresolved-diagnostic processing is a substantial secondary source/algorithm cost:

`emitUnresolvedReference` or `emitTypeScriptOutcomeDiagnostic` → `AppendDiagnosticToNode` → `appendDiagnosticsFromProperties` → re-normalize existing diagnostics → `normalizeDiagnosticMetadata` → `decodeStructuredResolutionOutcome` → two JSON unmarshals plus validation.

- `AppendDiagnosticToNode`: **0.01 s flat / 113.20 s cumulative (13.78%)**.
- `decodeStructuredResolutionOutcome`: **0.27 s flat / 112.93 s cumulative (13.75%)**.
- `encoding/json.Unmarshal` beneath this path: **109.19 s cumulative**.
- Decoder allocation: **20,314.18 MB (11.35%)**; append path allocation: **22,481.16 MB (12.56%)**.
- Workload denominator: **97,806 successfully attached unresolved-reference diagnostics**.

This is a verified contributor in the profiled runtime. It is not evidence by itself that any named slice or commit caused a historical regression.

### CONTRIBUTING FACTORS

- Allocation-driven garbage collection: the heap profile estimates **178,951.60 MB (~174.76 GiB)** allocated and **3,339,541,271** allocated objects over the run. `runtime.gcBgMarkWorker` has **222.39 s cumulative samples (27.07%)**; `runtime.gcDrain` has **221.79 s** and `runtime.scanObject` **134.39 s**. These concurrent inclusive samples overlap application paths and are not additive wall time. GC is a runtime consequence of the verified transient-allocation mechanisms, not an independent primary cause.
- Database persistence: exact named wall time **25,403.8452 ms (4.258%)**, with 120,853 node rows and 167,428 relationship rows, 18 node COPY tables, 93 relationship pairs, and zero fallback, failure, or skipped relationship.
- Parse: exact named wall time **14,739.7751 ms (2.471%)** for 763 successful code files and 6,974,573 parser input bytes; meaningful but not the root.
- Graph snapshot serialization: inside the in-run unphased bucket, `writeGraphSnapshot` has **10.76 s cumulative CPU samples**, allocates **2,024.95 MB**, and writes a **666,772,547-byte** JSON snapshot. It is material but far smaller than the hot resolver path.
- In-run unphased work is **14,911.2894 ms (2.500%)**; wrapper-minus-benchmark is **3,346.3641 ms**. Current artifacts cannot partition all of those wall times exactly.

### UNVERIFIED HYPOTHESES / CLAIMS NOT MADE

- No comparable baseline exists for this exact runtime, source state, workload, graph semantics, profiling mode, and output. A historical regression, regression ratio, or regression-causing commit/slice is **unverified**.
- Exact function invocation counts, `workspace.imports` length, import-loop iterations, hit/miss distribution, and diagnostics-per-node normalization multiplicity are **not instrumented**.
- Exact wall time separately attributable to filesystem I/O, native DB I/O versus native compute, lock wait, GC pauses/concurrent GC, graph-snapshot encode versus flush, DB close, heap-profile write, AI-context write, registry write, file projection, and final JSON serialization is **not recoverable from CPU + heap pprof and current benchmark boundaries**.
- No claim is made that TypeScript catalog lookup or external-symbol materialization is the primary cause; measured evidence is against that claim.

## Authority and boundaries observed

- Read-only diagnosis except for permitted artifacts under the exact evidence root and this report.
- No code, tests, plan, ledger, `AGENTS.md`, or skill edits by this lane.
- No Git mutation, commit, cleanup, network/install/package/build, alternate worktree, target repository access, protected/shared cache access, or graph-reader command after the measurement.
- No second `analyze` invocation.
- The worktree was already dirty with Owner/other-lane work; those changes were preserved and not modified by this lane.
- Pre-run inspection found no analyze lock. Fifteen processes named `anvien` were observed by name/PID only; no true analyze-lock holder was found, and no process was killed.

## Controlled runtime identity

| Identity item | Verified value |
|---|---|
| Executable | `E:\Anvien\anvien\bin\anvien.exe` |
| Version | `1.2.8` |
| Executable size | `73,751,552` bytes |
| Executable SHA-256 | `62F7CFBB26C73A7B4C22AC64904CE5521C546DD6D5D6DB79E014E99691180569` |
| Go build version | `go1.26.3` |
| VCS revision/state | `c28660118fc606ea3e19ad2f6cfb206a768f46f1+dirty` |
| Build tag | `ladybugdb` |
| Runtime manifest | `E:\Anvien\anvien\bin\anvien-runtime.json`; source `E:/Anvien`, Windows/amd64, `ladybugdb` |
| Native library | `E:\Anvien\anvien\bin\lbug_shared.dll`, 20,230,656 bytes |
| Native library SHA-256 | `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |

The invocation used the direct E-only executable, never the bare/global command. `anvien --help` and `anvien analyze --help` were checked first; the selected runtime advertised `--benchmark-json`, `--benchmark-label`, `--cpuprofile`, and `--memprofile`. It did not advertise runtime-trace, block-profile, or mutex-profile flags.

All process-local home/config/temp variables (`ANVIEN_HOME`, `HOME`, `USERPROFILE`, `APPDATA`, `LOCALAPPDATA`, `TEMP`, `TMP`, and XDG variables) were redirected to descendants of the exact evidence root. Git safe-directory handling was process-local. No C-drive executable, target repository, network, or protected root was used or inspected as profiling authority.

## Exact invocation and outcome

Working directory: `E:\Anvien`

```powershell
& 'E:\Anvien\anvien\bin\anvien.exe' analyze . --force --benchmark-json 'E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d\analyze-benchmark.json' --benchmark-label 'analyze-performance-root-cause-01a02c2d' --cpuprofile 'E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d\analyze-cpu.pprof' --memprofile 'E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d\analyze-heap.pprof' --progress --json
```

| Result | Value |
|---|---:|
| Start | `2026-08-23T10:06:51.0177628+07:00` |
| End | `2026-08-23T10:16:50.9416411+07:00` |
| Retained process elapsed | `599,915.2275 ms` |
| Exit | `0` |
| Files scanned / parsed / failed | `2,075 / 763 / 0` |
| Graph nodes / relationships | `120,853 / 167,428` |
| File-projection dependency edges / unresolved files | `20,310 / 478` |
| Graph snapshot bytes | `666,772,547` |
| Graph snapshot SHA-256 | `38B32C12CF558D09DE832234EE38BCBDA52FEA72F6C6DB6C68510A56E1624739` |

The independently retained successful full-build evidence pointer reported 2,074 scanned / 763 parsed / 0 failed / 120,843 nodes / 167,418 relationships / 20,310 dependency edges / 478 unresolved projections. The controlled run therefore processed **one more file and emitted ten more nodes and ten more relationships**, while retaining dependency and unresolved-projection counts. The measured workload was not reduced.

Pre-run snapshot seal: 666,758,770 bytes, SHA-256 `C616A2B68C511FDE9743E6A69E5AEC01F297BC6CD79637CBE19A89AAFC0DFE30`.

DB row counts equal the final in-memory graph counts (120,853 / 167,428), with zero DB fallback/failures/skips. This is count-level persistence parity evidence; no post-measurement graph-reader query was run.

## Wall-time accounting

Benchmark timer: **596,568.8634 ms**. Named-phase sum: **581,657.5740 ms**. In-run unphased: **14,911.2894 ms (2.500%)**.

| Named phase | Wall ms | % benchmark total |
|---|---:|---:|
| resolution | 533,746.9551 | 89.469% |
| db_load | 25,403.8452 | 4.258% |
| parse | 14,739.7751 | 2.471% |
| semantic_enrichment | 3,773.8566 | 0.633% |
| cross_file_binding | 1,737.8772 | 0.291% |
| scan | 588.5365 | 0.099% |
| documents | 551.8579 | 0.093% |
| processes | 447.3999 | 0.075% |
| routes | 336.0774 | 0.056% |
| tools | 181.6245 | 0.030% |
| communities | 119.2247 | 0.020% |
| structure | 15.2747 | 0.003% |
| orm | 13.1238 | 0.002% |
| mro | 2.1454 | <0.001% |
| cobol | 0 | 0% |

Requested category separation:

- Scan, parse, resolution, communities, processes, semantic, and DB load have exact named-phase wall times above.
- Serialization does not have its own benchmark phase. Source and pprof place graph snapshot serialization after DB load and inside the 14.911 s in-run residual; it has 10.76 s cumulative CPU samples, but not an exact wall duration.
- AI-context generation is after `analyze.Run`; pprof attributes 1.60 s cumulative samples, but there is no dedicated wall boundary.
- Other post-run paths have cumulative CPU samples of 0.54 s for file projection, 0.09 s for registry/commit recording, and 0.05 s for forced-GC heap-profile writing. These overlap the 3.346 s wrapper residual and are not exact wall timings.
- Lock acquisition and force-storage preparation happen before the benchmark timer. No lock holder was observed, but exact pre-timer wall split is unavailable.

## Profile interpretation rules

- CPU profile duration is **599.49 s** with **821.41 CPU sample-seconds (137.02% of wall)**. Multiple threads/goroutines, concurrent GC, and native activity allow sample-seconds to exceed elapsed wall time.
- “Flat” is exclusive sampled CPU time. “Cumulative” is inclusive sampled CPU time. Caller and child rows overlap; they must not be summed as wall time or independent savings.
- Heap `alloc_space` and `alloc_objects` values are pprof sampling estimates/extrapolations. `inuse_space` was captured after the CLI forced a GC for the heap profile.
- The profiler adds overhead. This run is authoritative for attribution of current hot paths, not an uninstrumented throughput baseline.

## Ranked granular bottlenecks

Rows are ranked by their demonstrated contribution to the current run. Inclusive rows intentionally expose call paths and are explicitly marked as overlapping.

| Rank | Granular owner / exact call path | Wall envelope | CPU flat / cumulative (% of 821.41 s samples) | Allocation / live heap | Workload denominator | Causal confidence and artifact |
|---:|---|---|---|---|---|---|
| 1 | `ResolveBoundInto` → call/access/type/receiver checks → `explicitImportNameClaimed` → `cleanPath` | Within resolution: 533,746.9551 ms; function wall is not separately timed | `explicitImportNameClaimed` **39.43 / 319.95 s (4.80% / 38.95%)**; child `cleanPath` 280.15 s. Total `cleanPath` **0.88 / 324.75 s (0.11% / 39.54%)** | Import-claim path **129,779.47 MB (72.52%)**, **2,648,342,788 objects (79.30%)**; `cleanPath` **149,590.52 MB (83.59%)**, **3,056,971,217 objects (91.54%)** | Per-fact checks over all `w.imports`; output counters: 19,297 resolved calls, 12,709 resolved accesses, 14,492 resolved type refs, 93,078 unresolved refs. Exact Q/I absent. | **High — verified primary cause.** `cpu-peek-resolution-hotpaths.txt`, `cpu-list-resolution-hot-source.txt`, `heap-alloc-space-peek-hotpaths.txt`, `heap-alloc-objects-peek-hotpaths.txt` |
| 2 | `resolveCall` → `explicitImportCallState` → full `w.imports` scan → `cleanPath` | Same resolution envelope; overlaps rank 1 | `explicitImportCallState` **3.90 / 47.51 s (0.47% / 5.78%)**; `cleanPath` child 43.45 s | **19,499.04 MB (10.90%)**, **400,176,533 objects (11.98%)**, overlapping total `cleanPath` | Call facts; 19,297 resolved calls plus unresolved call sites not separately counted | **High — verified same algorithmic family.** Same four hot-path artifacts |
| 3 | unresolved emit → `AppendDiagnosticToNode` → reload and normalize existing diagnostics → `decodeStructuredResolutionOutcome` → JSON decode/validation | Same resolution envelope; function wall absent | Append **0.01 / 113.20 s (0.001% / 13.78%)**; decoder **0.27 / 112.93 s (0.033% / 13.75%)**; JSON unmarshal 109.19 s | Append **22,481.16 MB (12.56%)**; decoder **20,314.18 MB (11.35%)**, **251,538,202 objects (7.53%)** | 97,806 attached diagnostics; per-node existing-list lengths absent | **High — verified secondary cause.** `cpu-peek-resolution-hotpaths.txt`, `cpu-list-diagnostic-hot-source.txt`, alloc artifacts |
| 4 | transient allocation → concurrent mark/drain/scan GC | Run-wide, overlaps ranks 1–3 | `gcBgMarkWorker` **222.39 s cumulative (27.07%)**; `gcDrain` 221.79 s; `scanObject` 134.39 s | Total **178,951.60 MB**, **3,339,541,271 objects**; post-GC live heap **601.04 MB**; max Sys **2,524,863,096 bytes** | Entire run | **High as contributor/consequence; not independent root.** `cpu-top.txt`, `cpu-cum.txt`, heap top artifacts |
| 5 | `loadGraph` → CSV export → Ladybug COPY/native query/commit | **25,403.8452 ms (4.258%) exact wall** | `loadGraph` 24.24 s cum; CSV export 5.19 s; load 19.05 s; native query 21.71 s (overlapping contexts) | CSV export 457.98 MB | 120,853 node rows; 167,428 relationship rows; 18/93 COPY owners; zero fallback/failure/skip | **High contribution measurement; not root.** `cpu-peek-persistence-serialization.txt`, benchmark JSON |
| 6 | `writeGraphSnapshot` → per-node/per-relationship `json.MarshalIndent` → buffered write/rename | Within 14,911.2894 ms unphased; exact wall absent | 0 flat / **10.76 s cumulative (1.31%)**; JSON owner 10.49 s | **2,024.95 MB (1.13%)** | 120,853 nodes + 167,428 relationships; output 666,772,547 bytes | **Medium-high contributor; wall split blind.** persistence/source peek and alloc artifact |
| 7 | parse files | **14,739.7751 ms (2.471%) exact wall** | `parseFiles` 14.01 s cumulative (1.71%) | Not a top allocation owner | 763/763 succeeded, 0 failed, 6,974,573 bytes | **High measured non-root.** benchmark JSON, `cpu-cum.txt` |
| 8 | outcome collector validate/clone/sort, then outcome projection/merge/diagnostic-carrier validation | Same resolution envelope; overlaps higher rows | collector finalize 0.04 / **0.29 s (0.035%)**; projection 0.04 / **3.84 s (0.47%)**, of which diagnostic sites 2.69 s | finalize 47.43 MB; projection 783.37 MB | Resolution-outcome count not emitted; 97,806 diagnostic inputs and graph carrier counts provide outer workload | **High evidence against finalization/projection as primary.** `cpu-peek-stage-owners.txt`, persistence/finalize alloc artifact |
| 9 | TypeScript authority catalog/load/actual lookup/finalize/external materialization | Same resolution/unphased envelopes; overlaps rank 1 wrapper | catalog construction/load **0.15 s**; actual `Authority.LookupGlobal` and `LookupMember` ~0.01 s each; wrapper member lookup 6.09 s is 6.06 s `repositoryReceiverClaimed`; finalize 0.01 s; external emit 0.01 s | external emit 1.01 MB; finalize 6.52 MB | 5,495 capability-unavailable; 0 external declarations resolved | **High evidence against catalog/materialization as primary.** `cpu-peek-typescript-authority-catalog.txt` |
| 10 | AI-context, file projection, registration, heap-profile write, benchmark write, final JSON/output | Wrapper residual **3,346.3641 ms**, not individually timed | AI context 1.60 s; projection 0.54 s; register 0.09 s; heap profile 0.05 s; benchmark 0.01 s cumulative | Not top alloc owners in retained tables | 120,853/167,428 graph supplied to post-run consumers | **Medium contribution evidence; exact wall split blind.** persistence peek; CLI source |

## Resolution owner decomposition

This section distinguishes actual call owners rather than inferring from the phase name.

### Workspace/index construction

- `BuildCrossFileBinding` → `buildWorkspace`: **1.59 s cumulative samples (0.19%)**, **485.36 MB alloc (0.27%)**.
- Cross-file named wall: **1,737.8772 ms (0.291%)**.
- Source builds `importsByReceiver` while resolving imports, but the hot import-name checks traverse `w.imports` directly.
- Conclusion: workspace construction is not the current bottleneck.

### Per-file/per-fact resolution

`ResolveBoundInto` iterates every file and then every call, access, and type annotation at [internal/resolution/resolve.go:91](../../internal/resolution/resolve.go#L91). Its cumulative sample attribution is **496.63 s (60.46%)**.

| Fact owner | Inclusive samples | Direct child attribution (overlapping inclusive contexts) |
|---|---:|---|
| `resolveCall` | 219.14 s (26.68%) | receiver claim 64.82 s; direct import-name scan 60.90 s; import-call-state scan 47.51 s; unresolved emit 27.34 s; TS outcome record 7.56 s; actual catalog/member lookup 4.21 s |
| `resolveAccess` | 191.01 s (23.25%) | receiver claim 90.26 s; direct import-name scan 51.64 s; unresolved emit 44.00 s; TS outcome record 2.79 s; actual catalog/member lookup 1.00 s |
| `resolveTypeAnnotation` | 81.78 s (9.96%) | annotation import-claim scan 46.59 s; TS outcome record 18.20 s; unresolved emit 15.30 s; actual TS annotation lookup 0.90 s |

The `explicitImportNameClaimed` caller split is:

- `repositoryReceiverClaimed`: **160.82 s (50.26% of the callee context)**.
- direct `resolveCall`: **60.90 s (19.03%)**.
- direct `resolveAccess`: **51.64 s (16.14%)**.
- `typeScriptAnnotationImportClaimed`: **46.59 s (14.56%)**.

### External authority/catalog lookup

- `tsstdlib.NewAuthority` / catalog load: 0.15 s samples.
- The actual catalog methods are at sample-floor scale (~0.01 s each).
- `lookupTypeScriptMember` appears as 6.09 s, but 6.06 s is its repository receiver/import-claim guard; this is rank 1, not catalog search.
- `recordTypeScriptLookup` is 28.55 s, of which 28.38 s is downstream diagnostic append/decode.
- External declarations resolved: 0; external capability unavailable: 5,495.

### Outcome finalization, merge/sort, and graph projection

- `resolutionOutcomeCollector.finalize`: 0.29 s cumulative, dominated by sort (0.17 s) and validation (0.06 s).
- `finalizeTypeScriptAuthorityResults`: 0.01 s.
- `emitTypeScriptExternalSymbols`: 0.01 s.
- `projectResolutionOutcomes`: 3.84 s, including `resolutionOutcomeDiagnosticSites` 2.69 s and marshal 0.54 s.
- These are real costs but cannot explain 533.747 s resolution wall.

## Source attribution

### Primary verified mechanism

1. Workspace import construction normalizes the fact path before storing it:
   - [`resolveImports`](../../internal/resolution/indexes.go#L287), especially [path normalization at line 294](../../internal/resolution/indexes.go#L294) and [append to `w.imports` at line 303](../../internal/resolution/indexes.go#L303).
2. The hot import-name owner ignores the keyed receiver index and scans the full slice:
   - [`explicitImportNameClaimed`](../../internal/resolution/resolve.go#L900), loop at line 906 and repeated `cleanPath(item.Fact.FilePath)` at line 907.
3. Calls perform another independent scan with repeated normalization:
   - [`explicitImportCallState`](../../internal/resolution/export_resolution.go#L306), loop at line 319 and normalization at line 321.
4. The repeated normalizer allocates through three string/path transformations:
   - [`cleanPath`](../../internal/resolution/indexes.go#L989), `filepath.ToSlash(filepath.Clean(strings.ReplaceAll(...)))` at line 993.
5. Call/access/type and receiver owners invoke the hot checks:
   - [`resolveCall`](../../internal/resolution/resolve.go#L385), [`resolveAccess`](../../internal/resolution/resolve.go#L630), [`resolveTypeAnnotation`](../../internal/resolution/resolve.go#L750), [`typeScriptAnnotationImportClaimed`](../../internal/resolution/resolve.go#L824), and [`repositoryReceiverClaimed`](../../internal/resolution/resolve.go#L872).

This establishes a source/algorithm/data/runtime chain: large per-fact workload → repeated whole-import traversal → repeated string/path allocation → GC pressure → dominant CPU and wall envelope.

### Secondary verified mechanism

1. [`emitUnresolvedReference`](../../internal/resolution/emit.go#L180) records a structured outcome and calls `AppendDiagnosticToNode` at line 218.
2. [`AppendDiagnosticToNode`](../../internal/graphhealth/diagnostics.go#L22) normalizes the incoming diagnostic at line 39, then calls `appendDiagnosticsFromProperties` at line 40.
3. [`appendDiagnosticsFromProperties`](../../internal/graphhealth/diagnostics.go#L92) loads and normalizes all existing diagnostics, normalizes the incoming diagnostic again, scans buckets, and sorts after an append.
4. [`normalizeDiagnosticSlice`](../../internal/graphhealth/diagnostics.go#L216) re-runs normalization for each existing value.
5. [`decodeStructuredResolutionOutcome`](../../internal/graphhealth/diagnostics.go#L327) unmarshals the same `Note` into an envelope at line 335 and into the full outcome at line 346 before nested validation.

This establishes a second source/algorithm/data chain: 97,806 diagnostic appends → repeated normalization of existing values plus duplicate JSON parsing → ~113 s cumulative samples and ~20–22 GB allocation on the path.

### Persistence and serialization owners

- [`loadGraph`](../../internal/analyze/analyze.go#L555) exports graph CSVs and loads them through the native runner. Exact phase wall is 25.404 s.
- [`writeGraphSnapshot`](../../internal/analyze/analyze.go#L642) serializes all nodes and relationships, flushes/closes, and atomically renames the temp file. Per-item `json.MarshalIndent` begins at [`writeIndentedJSONValue`](../../internal/analyze/analyze.go#L764).
- [`analyze.Run`](../../internal/analyze/analyze.go#L208) starts its benchmark timer only after path resolution, storage-lock acquisition, and force preparation at line 230; snapshot writing is at lines 456–461, and benchmark writing follows metric finalization at lines 462–468.
- CLI post-run owners are visible at [`newAnalyzeCommand`](../../internal/cli/command.go#L155), lines 227–262: CPU profile, `analyze.Run`, heap profile, registration, AI context, file projection, and output.

## Evidence for/against coupling with current named graph work

These are current-code call-path observations, not phase/slice framing and not regression attribution.

| Area | Current evidence | Causal statement |
|---|---|---|
| TypeScript standard-library authority/catalog | Catalog construction 0.15 s; actual authority lookup at sample-floor scale; 0 external declarations resolved | Evidence **against** catalog search as the primary current bottleneck. Its wrapper can still enter the primary repository/import scan. |
| External target/materialization | `emitTypeScriptExternalSymbols` 0.01 s, 1.01 MB; resolved external declarations 0 | Evidence **against** external materialization as the primary current bottleneck. |
| Structured resolution outcomes | Outcome finalization 0.29 s and projection 3.84 s are small; downstream diagnostic parsing is 112.93 s | Structured outcome finalization/projection is not primary, but structured diagnostic consumption is a verified secondary contributor. |
| Current graph-health projection change | The profiled binary contains and samples `decodeStructuredResolutionOutcome`; the working source diff contains that decoder/policy | Direct current-runtime coupling is verified. Chronology, before/after delta, and regression attribution are not. |

No statement assigns the slowdown to P6-B, P6-C2, P6-C3, P6-D, or any commit solely from timeline.

## Historical evidence and comparability

| Evidence | Retained timing/workload | Comparability conclusion |
|---|---|---|
| Successful `scripts\full-build.ps1` task evidence | Exact command elapsed **678,182 ms (11m18.182s)**; output 2,074/763/0 and 120,843/167,418. The “~8 minute analyze” statement was inferred from polling, not an exact timer. | Workload pointer only. The successful script resolved a global wrapper outside E, so it is not runtime/profile identity authority. The 1,042,071 ms turn included a lost-handle first attempt plus rerun and is not a single build timing. |
| `reports\benchmark\2026-05-16-120908-anvien-self-analyze.json` | 13,588.1 ms; 676/520/0; 19,984/47,622; parse 5,586.7 ms; resolution 1,057.8 ms; DB 4,902.5 ms | Materially smaller, older graph/workload and different semantics. Historical context only; not a baseline for a regression ratio. |
| `reports\Investigation\p1b_ts_identity_analyze_benchmark.json` | 50,030.129 ms; 1,359/887/0; 84,807/114,125; scan 8,224.339 ms; parse 9,006.084 ms; resolution 2,654.981 ms; semantic 1,123.075 ms; DB 18,411.840 ms | Materially different corpus/output and no matching runtime/profile state. Historical context only. |

Historical artifact seals:

- May 16 composite benchmark SHA-256: `798AAED09DD8E51726CEF549D967C9CB0D26D0E7CFB88089FA0583F169D8CF29`.
- P1-B benchmark SHA-256: `98FA317B2E9268A2A236278B05CA6348B395AF8419F4197B61DEA9B6D50B737A`.

## Evidence artifact seals

Primary raw artifacts:

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| `analyze-benchmark.json` | 9,711 | `66CBC05732B4AA2E55F6C3C84718EC802D82CB9CD608F6F529873704DD21187E` |
| `analyze-cpu.pprof` | 392,687 | `35C6D4DD35284B67D467B4ADBBBAAFE71F5B134FB41F968F88921563E4E4D9FB` |
| `analyze-heap.pprof` | 107,005 | `7874869AFE9F5B79F171F43F0FF9899CDC8CC95C57C2605AF80ABBBE97BD870A` |

Derived pprof evidence:

| Artifact | Bytes | SHA-256 |
|---|---:|---|
| `cpu-top.txt` | 8,331 | `A00F29CB5BD4267874D42D819FB1D5CAA32FE4A3D65314551358DDF39266C62B` |
| `cpu-cum.txt` | 10,435 | `2FE2FC630EC969CBB639275C9FD2BA399A98DC52850D32C8914C71D921CC0B68` |
| `cpu-peek-resolution-hotpaths.txt` | 16,612 | `664AD55F183066EA1937B9A876EED405C7C70D079D8C3EE00EEB9A80A916443B` |
| `cpu-list-resolution-hot-source.txt` | 3,637 | `531052F14D1DF5A018E0E664A5E23AC8D6103E14FD16CB5B34F2B97654C0CAB0` |
| `cpu-list-resolution-driver-source.txt` | 33,035 | `91F8301DE6AAA94A2CB8B15D88DBEE3D96D0BC03A48A98D7CB9C5CE5635163E2` |
| `cpu-list-diagnostic-hot-source.txt` | 4,844 | `7F9BB23643071BC722BF6DC3792AF45FC285A0B8F91B48B61DAD7021013A30A1` |
| `cpu-peek-stage-owners.txt` | 12,729 | `1B949306A6BFA54EC0EAB3D422501EDDDB5E314D872C47D05575DA8D4DA87A69` |
| `cpu-peek-typescript-authority-catalog.txt` | 7,477 | `1B709146714C355A2EB757226672A5B1526139D0FE7B784C1D121FBC55EC3609` |
| `cpu-peek-persistence-serialization.txt` | 12,232 | `C9103424406E05F899F2C1C1FD5CAC9A0A2E0C2FEB752B0CBD2A8009F8F95155` |
| `heap-alloc-space-top.txt` | 7,525 | `D1CF01F2566600BF0D5E7D78F62B2467FA289A362B5F4CD3E3E63366967E1122` |
| `heap-alloc-objects-top.txt` | 6,024 | `F9135432BC862D0CF7F9B4B60BDD84BEDA855247EA1072F721B2A9280AB0829E` |
| `heap-inuse-space-top.txt` | 8,361 | `5428DB527B5EA5AF07A8922C3AD2E2AA86D6392B3866B517A6322A885D54F46B` |
| `heap-alloc-space-peek-hotpaths.txt` | 15,456 | `9E7DE6E51E8F5BBE34A4605C535EE1B3C58ABD293FA197967D6259C34F0D05EF` |
| `heap-alloc-objects-peek-hotpaths.txt` | 11,595 | `84BAC6E7365C7B193B02DCB9697F90D9C5CFB587AE1AB9433865F9631B0E9D05` |
| `heap-alloc-space-peek-persistence-finalize.txt` | 7,264 | `B2938B33ADF51172E48DC7EB111591B29BB3F22B60B40715A583216DA7154602` |

All artifacts remain under `E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d`; no cleanup was performed.

## Measurement gaps and minimal additional-measurement request

The existing run is sufficient to identify the primary and secondary current causes. A second run is **not needed for that conclusion**. It would only be needed to obtain exact invocation multiplicity and to partition residual/blocking/I/O wall time.

### Request to Main — do not execute without explicit new authority

Authorize exactly one future E-only run only after a separately authorized profiling-only runtime exposes instrumentation that the current CLI does not have. The diagnostic runtime must retain the same source identity and semantics and must write only beneath:

`E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d\followup`

Required instrumentation, with no graph-output behavior change:

1. Runtime trace, block profile, and mutex profile.
2. Monotonic nanosecond boundaries for storage lock acquire/wait, force preparation, TypeScript authority construction, graph compact, DB-runner creation, DB load substeps, DB close, snapshot encode/flush/close/rename, benchmark write, forced-GC heap-profile generation, registration, AI-context generation, file projection, and final output serialization.
3. Counters for:
   - `explicitImportNameClaimed` calls, items visited, hits, and misses, tagged by direct caller;
   - `explicitImportCallState` calls/items visited/hits/misses;
   - `cleanPath` calls and input/output byte totals, tagged by caller;
   - call/access/type input fact counts, not only resolved outputs;
   - diagnostic append calls, existing-list length histogram, new-versus-existing normalizations, structured decoder calls, and JSON bytes decoded;
   - process I/O operation and byte deltas at the named persistence/serialization boundaries.

Expected future command after those flags exist (illustrative exact invocation contract; it is **not valid on the current binary** and must not be run now):

```powershell
& 'E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d\followup\anvien-instrumented.exe' analyze . --force --benchmark-json 'E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d\followup\analyze-benchmark.json' --benchmark-label 'analyze-performance-root-cause-followup' --cpuprofile 'E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d\followup\analyze-cpu.pprof' --memprofile 'E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d\followup\analyze-heap.pprof' --trace 'E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d\followup\analyze-runtime.trace' --blockprofile 'E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d\followup\analyze-block.pprof' --mutexprofile 'E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d\followup\analyze-mutex.pprof' --detail-metrics-json 'E:\Anvien\.tmp\analyze-performance-root-cause-01a02c2d\followup\analyze-detail.json' --progress --json
```

Expected proof from that one run:

- exact `Q`, `I`, average/worst import items visited, and caller-specific loop multiplicity;
- exact new-versus-existing diagnostic normalization/decode multiplicity and per-node list distribution;
- wall partition among CPU compute, allocation/GC pauses and concurrent GC, filesystem/native DB I/O, lock/mutex/block wait, snapshot/benchmark/post-run serialization, and remaining unexplained time;
- confirmation that the residual buckets reconcile to process elapsed without double-counting inclusive CPU samples.

The present lane did not create that runtime, did not add those flags, and did not run a second analyze.

## Handoff statement

The evidence establishes the current causal mechanism with high confidence while keeping historical attribution and unmeasured wall splits explicitly unverified. This report supplies measurements and source paths only. Architecture, planning, implementation, review, phase naming, and acceptance decisions belong to subsequent owners.
