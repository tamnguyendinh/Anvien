# Child 06A D004 Residual Attribution

Verdict: `D004_RESIDUAL_ATTRIBUTION_READY`

## Phạm vi và authority

- Slice: `P2-A`.
- Active parent/child: unchecked `B1-P1A-OP001 resolution` / unchecked `B2-P2A-A001-D004 project_resolution_outcomes`.
- Mục tiêu duy nhất của lane: xác định một causal mechanism hiện tại, attributable owner, và complete synchronous path của D004 từ source hiện tại, accepted A003 packets/profiles tách riêng theo target, và đúng phần perturbation D004 của A005 bị reject.
- Boundary: read-only attribution. Report này là durable write duy nhất. Không có architecture/solution, expected gain, production/test/script/ledger/plan edit, build/test, target analyze, measurement mới, candidate/attempt, Supervisor verdict, disposition, detect, stage hoặc commit.
- A005 chỉ là controlled perturbation/non-duplication evidence. A005 không phải accepted production authority và hướng canonical outcome-byte của nó không được tái dùng làm D004 solution.

## Identity và input gate

- Expected/current HEAD: `8a28683e010cf0b00ac4a3099dd2c0d8e4b69d69` / exact match.
- Subject: `docs(plan): apply Child 06A optimization cost floor`.
- Worktree sạch; staged set rỗng trước graph work.
- Current `internal/resolution/outcome.go` SHA-256: `02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E`.
- A003 executable SHA-256: `2DBBD7AC70C04FB62C3D8AB4A90F50E7F90C51251E82B67883A58B61D42B426B`.
- A005 executable SHA-256: `0178DAC396E0DD4D0BECCE7B803378A3400013F0B508EC95F9221A0020CF96F3`.
- Cheapapp A003 profile/comparison SHA-256: `806FAD79B12C567618778A8232E8521991882CE6C745AC2EBBE7B09ECC267107` / `8D64F905A4413E375DF8CA75E6465EE09BBBEA777DC755DF304F09BB67691C2F`.
- Restaurant Manager A003 profile/comparison SHA-256: `4272FFAB0ABE4C14C470A139A31A601F31323B1373E0FA2986A09935456FA3E2` / `ED19F20BEF5490C361D2A8F0C7634A8D2A7F7EC43371F5F96CF09117447B557C`.
- Cheapapp/Restaurant A005 profile SHA-256: `0DA6425130E12F9169BCC990C349710D4BA4277E9C30612B9509CE61D9EF9389` / `5FC396A5DD96859B9E863F91CD029ACCB0E55C2B8C1744CD8E9388D5426F98B3`.
- Mọi path bắt buộc tồn tại; mọi hash do handoff chỉ định đều khớp tuyệt đối.

## Exact result

### Causal mechanism hiện tại

Current D004 residual materially common trên cả hai target là **graph-wide diagnostic-property JSON round-trip** do `resolutionOutcomeDiagnosticSites` sở hữu:

1. Helper quét toàn bộ `g.Nodes`.
2. Với mỗi node có `graphhealth.DiagnosticPropertyKey`, helper luôn `json.Marshal(value)` từ graph property động.
3. Ngay sau đó helper `json.Unmarshal(raw, &[]graphhealth.Diagnostic)` để tạo lại typed diagnostic slice.
4. Chỉ sau round-trip này helper mới duyệt diagnostics, lookup final outcome theo `SourceSiteID`, kiểm tra resolved/unresolved status, so sánh exact `Diagnostic.Note` với canonical outcome bytes, và ghi nhận diagnostic sites.

Source anchors hiện tại:

- attributable owner: `internal/resolution/outcome.go::resolutionOutcomeDiagnosticSites`, lines `478-507`;
- full-node loop/property gate: lines `480-484`;
- JSON encode: line `485`;
- JSON decode: line `490`;
- status/byte checks: lines `493-504`;
- direct caller: `projectResolutionOutcomes` line `445`;
- measured D004 envelope owner: `projectResolutionOutcomes`, lines `399-455`.

`encoding/json` là causal descendant nhưng không phải repository owner. `projectResolutionOutcomes` sở hữu coordinator/envelope; `resolutionOutcomeDiagnosticSites` sở hữu chính loop và JSON round-trip tạo cost đã được profile chứng minh.

### Vì sao đây là exact owner thay vì broad envelope

- Current `projectResolutionOutcomes` flat CPU chỉ `0.03 s` trên Cheapapp và `0.01 s` trên Restaurant A003; phần lớn cumulative CPU nằm trong descendants.
- `resolutionOutcomeDiagnosticSites` là direct child lớn nhất và cùng xuất hiện vật chất trên cả hai accepted A003 profiles.
- Bên trong chính helper, JSON encode/decode chiếm gần toàn bộ subtree; graph-map/status work còn lại nhỏ hơn rõ rệt.
- Source chứng minh một operation cụ thể, không chỉ một denominator hay label: scan nodes -> property lookup -> JSON encode -> JSON decode -> exact diagnostic/outcome validation.

## Accepted A003 facts — targets giữ riêng

CPU samples dưới đây là cumulative/overlapping causal samples. Chúng không phải wall time, không cộng dồn thành wall, không dự đoán speedup, và không được average giữa targets.

| Fact | Cheapapp accepted A003 | Restaurant Manager accepted A003 |
|---|---:|---:|
| D004 wall | `4.652523600 s` | `5.995737900 s` |
| D004 denominator | `graphNodes=36575; graphRelationships=65045; outcomes=86742; referencesBySourceScope=25062; referencesByTargetDef=25062; workUnits=238486` | `graphNodes=66421; graphRelationships=112049; outcomes=186251; referencesBySourceScope=49259; referencesByTargetDef=49259; workUnits=463239` |
| Unresolved-reference diagnostic metric | `57683` | `129009` |
| Profile duration / total CPU samples | `20.46 / 25.80 s` | `20.84 / 23.96 s` |
| `projectResolutionOutcomes` cumulative | `4.28 s` | `5.44 s` |
| `resolutionOutcomeDiagnosticSites` cumulative | `2.59 s` | `3.81 s` |
| Helper -> `encoding/json.Unmarshal` | `1.99 s` | `2.77 s` |
| Helper -> `encoding/json.Marshal` | `0.53 s` | `0.86 s` |
| Sibling `projectResolutionOutcomes -> marshalResolutionOutcome` | `0.56 s` | `0.87 s` |
| Sibling `mergeExportBindingEvidence` cumulative | `0.95 s` | `0.49 s` |
| Sibling `projectReferenceIndexOutcomes` cumulative | `0.50 s` | `0.25 s` |

Diagnostic metrics là workload context, không phải exact count của diagnostic-bearing graph nodes hoặc số lần helper gọi JSON. Exact invocation/property-node count không được suy diễn.

## A005 controlled perturbation và non-duplication

A005 thay đổi canonical outcome-byte lifecycle: projection nhận bytes đã được record-time giữ lại và không còn gọi `marshalResolutionOutcome`. Candidate này correctness-PASS nhưng Main `NO_KEEP`, analyzer/process regressed, và production bytes đã được rollback. Chỉ D004-specific intervention được dùng tại đây.

| Fact | Cheapapp A003 -> rejected A005 | Restaurant A003 -> rejected A005 |
|---|---:|---:|
| D004 wall | `4.652523600 -> 4.023686900 s` | `5.995737900 -> 4.724638900 s` |
| D004 wall delta | `-0.628836700 s` | `-1.271099000 s` |
| D004-focused CPU cumulative | `4.28 -> 3.63 s` | `5.44 -> 4.29 s` |
| D004 `marshalResolutionOutcome` | `0.56 s -> no focused match` | `0.87 s -> no focused match` |
| `resolutionOutcomeDiagnosticSites` | `2.59 -> 2.52 s` | `3.81 -> 3.34 s` |
| Helper JSON split in A005 | `Unmarshal 2.01 s; Marshal 0.41 s` | `Unmarshal 2.39 s; Marshal 0.74 s` |

Interpretation bị giới hạn chặt:

- Intervention xác nhận projection-time outcome marshal là A005-owned sibling: khi exact call đó bị loại, symbol biến mất khỏi D004-focused A005 profiles và D004 wall/cumulative samples giảm trên từng target.
- Intervention đồng thời phân biệt residual owner: `resolutionOutcomeDiagnosticSites` và JSON round-trip của nó không bị A005 thay đổi, vẫn là direct child lớn nhất/material trên cả hai target sau perturbation.
- Vì vậy selected D004 attribution không trùng canonical outcome-byte direction đã bị reject. Không có A005 acceptance, speedup promotion, candidate reuse hay D004 solution inference.
- Chênh lệch CPU-sample của helper giữa A003/A005 chỉ là observed cumulative sample variation; không được đọc như một before/after speed result riêng của helper.

## Complete synchronous causal call path

### Upstream entry và D004 envelope

```text
internal/cli.newAnalyzeCommand.func1
  command.go:243 -> analyze.Run(...)
-> internal/analyze.Run
  analyze.go:365-370 -> runPhase(..., PhaseResolution, ...)
-> internal/analyze.runPhase
  analyze.go:1134-1154 -> synchronous run() and phase timing
-> internal/resolution.ResolveBoundInto
  resolve.go:57-128
  -> build binding occurrence index / create emitter
  -> emit definitions/imports/heritage
  -> synchronous calls, accesses, type-annotation loops
  -> method-dispatch and TypeScript finalization/external-symbol emission
  -> e.outcomes.finalize() at resolve.go:110
  -> projectResolutionOutcomes(...) at resolve.go:114
```

### D004 coordinator trước selected owner

```text
projectResolutionOutcomes (outcome.go:399-455)
-> build bySourceSite + encodedBySourceSite maps
   -> current per-outcome marshalResolutionOutcome at line 406
   -> cloneResolutionOutcome at line 410
-> for each graph relationship
   -> mergeSourceSiteIDs
   -> status validation
   -> mergeExportBindingEvidence
      -> exact-tuple dedupe
      -> sortExportBindingEvidence
      -> exportBindingEvidenceOrderFor
         -> JSON decode of export-binding Notes where applicable
-> projectReferenceIndexOutcomes(BySourceScope)
   -> mergeExportBindingEvidence
-> projectReferenceIndexOutcomes(ByTargetDef)
   -> mergeExportBindingEvidence
-> resolutionOutcomeDiagnosticSites at line 445
```

### Selected causal subtree và return path

```text
resolutionOutcomeDiagnosticSites (outcome.go:478-507)
-> for each g.Nodes entry
-> node.Properties[graphhealth.DiagnosticPropertyKey]
-> if present: encoding/json.Marshal(value) at line 485
-> encoding/json.Unmarshal(raw, &diagnostics) at line 490
-> for each diagnostic
   -> outcomes[diagnostic.SourceSiteID]
   -> reject resolved outcome with unresolved diagnostic
   -> compare diagnostic.Note == encoded[outcome.SourceSiteID]
   -> record diagnostic SourceSiteID
-> return diagnosticSites
-> projectResolutionOutcomes checks resolved/diagnostic overlap at lines 449-451
-> return to ResolveBoundInto
-> set resolution metrics/metadata and construct resolution.Result
-> analyze.Run assigns Graph/ResolutionOutcomes/Metrics
-> later synchronous graph consumers and publication boundary:
   MRO -> communities -> processes -> semantic enrichment -> Ladybug DB load
   -> Graph JSON snapshot -> CLI registry/meta, AI context, file projection, output
```

## Source boundaries và excluded siblings

Current `projectResolutionOutcomes` có bốn responsibility groups đã được source/profile phân biệt:

1. outcome map creation + current projection-time outcome serialization/cloning;
2. relationship evidence projection, bao gồm export-binding merge/sort/order decoding;
3. two reference-index projections;
4. graph diagnostic/outcome parity validation qua `resolutionOutcomeDiagnosticSites`.

Chỉ group 4 là selected exact residual mechanism/owner của report này. Group 1 chứa exact A005-controlled sibling và bị loại khỏi new/non-duplicative attribution. Groups 2-3 xuất hiện trong profiles nhưng nhỏ hơn và không được chọn. Report không tuyên bố chúng miễn cost hoặc đề xuất thay đổi chúng.

## Anvien graph/impact warning — full scope

Đúng một `anvien analyze --force` chạy từ `E:\Anvien` trước graph commands: exit `0`; `2282` scanned, `766` parsed code, `0` failed; graph `124599` nodes / `171465` relationships; indexed/current commit đều là `8a28683e...`; stale `false`. Không có graph refresh thứ hai.

### Containing file HIGH

`internal/resolution/outcome.go` là current/non-stale HIGH-risk file: `119` symbols, `87` inbound refs, `122` outbound refs, `117` local relationships, `148` unresolved source sites, `11` linked flows, `23` linked tests, `changedSinceAnalyze=false`.

### Exact symbol CRITICAL

`projectResolutionOutcomes` upstream impact depth `5`, tests included:

- risk `CRITICAL` (warning, không phải edit authority hoặc edit ban);
- `124` impacted symbols, `2` direct;
- `50` affected files;
- `26` affected modules;
- `42` affected processes;
- app-layer counts: `api_test=1; backend=7; backend_test=114; cli_launcher=2`;
- functional-area counts: `analyzer=25; cli=5; contracts=1; graph_health=1; providers=20; resolution=69; storage=3`.

Full affected-file inventory (`path: impactedSymbols`):

```text
cmd/access-candidate-audit/main.go:1; cmd/anvien/main.go:1;
internal/analyze/analyze.go:1; internal/analyze/analyze_test.go:10;
internal/analyze/legacy_resolver_conversion_test.go:5; internal/analyze/p6b_tsstdlib_test.go:5;
internal/analyze/p6c3_structured_outcome_test.go:2; internal/analyze/pipeline_parity_test.go:2;
internal/cli/command.go:2; internal/cli/command_test.go:1;
internal/contracts/legacy_contract_snapshot_conversion_test.go:1;
internal/graphaccuracy/access_candidate.go:1;
internal/lbugload/p3c_binding_occurrence_persistence_test.go:1;
internal/lbugload/p5d_export_proof_persistence_test.go:1;
internal/lbugload/p6c3_resolution_outcome_persistence_test.go:1;
internal/providers/astro/extract_test.go:1; internal/providers/c/extract_test.go:1;
internal/providers/cpp/extract_test.go:1; internal/providers/csharp/extract_test.go:1;
internal/providers/dart/extract_test.go:1; internal/providers/golang/extract_test.go:1;
internal/providers/java/extract_test.go:1; internal/providers/kotlin/extract_test.go:1;
internal/providers/php/extract_test.go:3; internal/providers/provider_parity_test.go:2;
internal/providers/python/extract_test.go:1; internal/providers/ruby/extract_test.go:2;
internal/providers/rust/extract_test.go:1; internal/providers/svelte/extract_test.go:1;
internal/providers/swift/extract_test.go:1; internal/providers/vue/extract_test.go:1;
internal/resolution/definition_collision_test.go:4; internal/resolution/export_resolution_test.go:4;
internal/resolution/graph_parity_test.go:1; internal/resolution/legacy_heritage_map_conversion_test.go:1;
internal/resolution/legacy_import_language_conversion_test.go:1; internal/resolution/legacy_p7_conversion_test.go:1;
internal/resolution/legacy_parity_test.go:4;
internal/resolution/legacy_scope_symbol_semantics_conversion_test.go:2;
internal/resolution/p2b_persistence_test.go:1; internal/resolution/p3c_binding_occurrence_test.go:2;
internal/resolution/p4c_export_projection_test.go:2; internal/resolution/p6b_tsstdlib_test.go:8;
internal/resolution/p6c2_external_symbol_test.go:3; internal/resolution/p6c3_structured_outcome_test.go:6;
internal/resolution/parser_integration_test.go:9; internal/resolution/proof_accuracy_golden_test.go:1;
internal/resolution/resolution_test.go:15; internal/resolution/resolve.go:3;
internal/resolution/type_alias_test.go:1
```

Full affected-module inventory (`module:hits/impact`):

```text
Resolution:68/direct; Analyze:19/indirect; Graphhealth:5/indirect; Cli:4/indirect;
Lbugload:3/indirect; Php:3/indirect; Providers:2/indirect; Ruby:2/indirect;
Access-candidate-audit:1/indirect; Astro:1/indirect; C:1/indirect; Contracts:1/indirect;
Cpp:1/indirect; Csharp:1/indirect; Dart:1/indirect; Golang:1/indirect;
Graphaccuracy:1/indirect; Httpapi:1/indirect; Java:1/indirect; Kotlin:1/indirect;
Lbugruntime:1/indirect; Python:1/indirect; Rust:1/indirect; Svelte:1/indirect;
Swift:1/indirect; Vue:1/indirect
```

Full affected-process names (graph reports `42`; duplicate `Main -> Result` rows are preserved as two process rows):

```text
Main -> IsGitRepo; Main -> NativeDBRunnerFactory; Main -> Result; Main -> Result;
Main -> StorageLockOptions; Main -> StoragePaths; Main -> AccessCandidateAuditInputs;
Main -> AddAccessReason; Main -> CleanPath; Main -> ContextFromCommand; Main -> GlobalDir;
Main -> NewVersionCommand; Main -> PropertyLabels; Main -> ReadMessage; Main -> Run;
Main -> Store; Main -> WriteAccessCandidateAuditResult; Main -> WriteRawMessageFramed;
Main -> SortedAccessLanguageKeys; Main -> SortedAccessReasonKeys;
NewAnalyzeCommand -> GitRoot; NewAnalyzeCommand -> IsGitRepo; NewRootCommand -> Store;
ResolveBoundInto -> AddNonEmpty; ResolveBoundInto -> ApplyFrameworkFact;
ResolveBoundInto -> CleanPath; ResolveBoundInto -> DefinitionNodesEqual;
ResolveBoundInto -> DiagnosticAppender; ResolveBoundInto -> Emitter;
ResolveBoundInto -> GenerateID; ResolveBoundInto -> Graph; ResolveBoundInto -> ReferenceIndex;
ResolveBoundInto -> RelationshipCallName; ResolveBoundInto -> ResolutionOutcomeCollector;
Run -> Adjacency; Run -> AppendUnique; Run -> GatherAncestors; Run -> IsOwnerLabel;
Run -> IsSupportedLanguage; Run -> NormalizePath; Run -> Query; Run -> StringProperty
```

Scope trên là blast-radius warning đầy đủ từ graph hiện tại. Lane này không dùng nó làm quyền sửa code.

## Evidence limits

- CPU profiles là sampled causal evidence, cumulative và overlapping; chúng không thay thế D004 wall, không cộng thành speedup, và không cho exclusive wall time.
- Không có exact count của nodes mang diagnostic property, JSON round-trip invocations, hay bytes được round-trip. Graph-node/outcome/reference denominators và diagnostic metrics không thay thế fact đó.
- Không có algorithm, architecture direction, expected gain hoặc rollback proposal trong report này.
- READY dựa trên conjunction: source owner/loop rõ ràng + same helper/stdlib subtree vật chất trong hai accepted A003 profiles + A005 controlled perturbation loại exact sibling marshal nhưng để selected subtree tồn tại trên cả hai target.

## Terminal handoff

- Exact verdict: `D004_RESIDUAL_ATTRIBUTION_READY`.
- Exact current mechanism: `resolutionOutcomeDiagnosticSites` quét graph nodes và JSON round-trip mọi present diagnostic property trước exact outcome/status/byte validation.
- Attributable owner: `internal/resolution/outcome.go::resolutionOutcomeDiagnosticSites` (`478-507`), invoked by `projectResolutionOutcomes` (`445`).
- Complete synchronous causal path và all current source boundaries được ghi ở trên.
- A005 chỉ phân biệt sibling/residual ownership; rejected direction không được tái dùng.
- Next owner: visible Main Orchestration task `01a0421e-a38e-7192-8e98-8a09fa72f04d`.
