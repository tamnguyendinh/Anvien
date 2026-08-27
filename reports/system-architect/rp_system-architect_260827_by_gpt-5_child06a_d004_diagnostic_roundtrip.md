# Child 06A D004 Diagnostic Property Round-Trip Architecture Decision

Verdict: `ARCHITECT_D004_READY_FOR_PLANNER`

## Scope và current authority

- Role/lane: fresh visible D004 attempt-local Architect.
- Slice duy nhất: Child 06A `P2-A` -> unchecked parent `B1-P1A-OP001 resolution` -> unchecked child `B2-P2A-A001-D004 project_resolution_outcomes`.
- Current hierarchy: top-level `25/30` terminal; OP001 children `16/17` terminal; D004 là child mở duy nhất của OP001; parent vẫn unchecked.
- Expected/current HEAD: `24df6bade53dabeafbbec278e25315c17f55ba75`, exact match; worktree và staged set đều sạch tại input gate.
- Attribution authority: `E:\Anvien\reports\Investigation\rp_child06a_d004_residual_attribution.md`, SHA-256 `47C55BC6EB5A73CF0C6877E5F66067A7CBB22012BC61C45259A405E56C156272`, verdict `D004_RESIDUAL_ATTRIBUTION_READY`.
- Current source identity: `internal/resolution/outcome.go` SHA-256 `02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E`; `internal/graphhealth/diagnostics.go` SHA-256 `6DE54D97A1B95E877B686DC3459238598F6060D28CE6329357A80BD7FC376D30`.
- Blast radius đã được fresh attribution packet cung cấp và được consume, không rerun: `internal/resolution/outcome.go` HIGH; `projectResolutionOutcomes` CRITICAL, `124` symbols / `50` files / `26` modules / `42` processes. Đây là scope warning, không phải edit ban.
- Output của lane chỉ là report này. Không có SPEC, ADR, execution plan, ledger, production/test/script/target edit, build, test, measurement, staging hay commit.

## Source/evidence basis

Quyết định này dùng đầy đủ authority sau:

- `AGENTS.md`, `working-rules`, `System-Architect` và binding `plan-rules.md`;
- bốn standard Child 06A ledgers qua EOF, gồm current D004, accepted A003 và rejected A005 sections;
- D004 attribution report qua EOF và exact hash ở trên;
- current `internal/resolution/outcome.go`;
- direct definitions cần cho diagnostic ownership/parity trong `internal/graphhealth/diagnostics.go`, `internal/graphhealth/policy.go`, `internal/graph/types.go`, `internal/resolution/emit.go`, `internal/resolution/resolve.go`, direct P6C3/appender tests và preserve-only graph diagnostic reader.

Source hiện tại chứng minh:

1. `NewDiagnosticAppender` là production owner ghi resolution diagnostics; sau first touch nó chuẩn hóa mọi supported legacy property representation và write-through chính `[]graphhealth.Diagnostic` vào `node.Properties[DiagnosticPropertyKey]` qua `Graph.AddNode`.
2. `ResolveBoundInto` chạy tuần tự: tạo một emitter/run-scoped appender, emit diagnostics, finalize outcomes, rồi mới gọi `projectResolutionOutcomes`; không có goroutine, flush, persistence hay external reader chen giữa emission và projection.
3. `resolutionOutcomeDiagnosticSites` hiện đọc lại chính graph property, nhưng luôn `json.Marshal(value)` rồi `json.Unmarshal` thành `[]graphhealth.Diagnostic` trước khi dùng duy nhất `SourceSiteID` và `Note` cho status/canonical-byte parity.
4. Diagnostic graph property và `Diagnostic` là public/persisted carrier shape cần giữ nguyên; Graph JSON, Ladybug/native và graph-accuracy readers tiêu thụ shape đó sau publication.

## Target-separated accepted basis

Không average hoặc combine hai target.

| Boundary | Cheapapp accepted A003 | Restaurant Manager accepted A003 |
|---|---:|---:|
| D004 `project_resolution_outcomes` | `4.652523600 s` | `5.995737900 s` |
| parent `resolution` | `20.472602300 s` | `20.850792800 s` |
| analyzer | `93.531974900 s` | `98.020546700 s` |
| process | `95.630648200 s` | `101.096911900 s` |
| D004 causal helper CPU | `2.59 s` | `3.81 s` |
| helper JSON unmarshal / marshal CPU | `1.99 / 0.53 s` | `2.77 / 0.86 s` |

Exact D004 denominators phải được giữ riêng:

- Cheapapp: `graphNodes=36575; graphRelationships=65045; outcomes=86742; referencesBySourceScope=25062; referencesByTargetDef=25062; workUnits=238486`.
- Restaurant Manager: `graphNodes=66421; graphRelationships=112049; outcomes=186251; referencesBySourceScope=49259; referencesByTargetDef=49259; workUnits=463239`.

CPU/profile values chỉ giải thích cause; chúng overlapping, non-additive, không thay D004 wall và không dự đoán numeric speedup.

## Exact causal invariant và decision

### Architecture rule

Graph diagnostic property thực tế phải tiếp tục là carrier/source-of-truth được projection kiểm tra. Khi property đã là canonical in-memory `[]graphhealth.Diagnostic` do current write-through appender sản xuất, projection phải đọc trực tiếp typed slice và không được encode/decode lại toàn property chỉ để lấy `SourceSiteID` và `Note`.

Direct typed consumption chỉ hợp lệ khi representation và các field projection tiêu thụ chứng minh observationally equivalent với current JSON round-trip. Nếu property không có exact canonical type, hoặc có JSON-unstable consumed string data, code phải đi nguyên compatibility fallback hiện tại: marshal property, unmarshal vào `[]graphhealth.Diagnostic`, giữ nguyên fail-closed error path, rồi mới chạy parity checks.

Các rule không được thay đổi:

- vẫn scan chính `g.Nodes` theo current order và đọc chính `node.Properties[DiagnosticPropertyKey]`;
- vẫn ignore diagnostic không có tracked final outcome;
- vẫn reject resolved outcome có unresolved diagnostic;
- vẫn so sánh exact `Diagnostic.Note` với canonical `encodedBySourceSite` bytes;
- vẫn record diagnostic site và reject resolved/diagnostic overlap;
- không sidecar, registry, shadow carrier, cached site set hoặc emitter-owned substitute được phép thay graph property validation.

### Implementation suggestion, không phải architecture law

Trong `resolutionOutcomeDiagnosticSites`, dùng một narrow type fast path cho exact `[]graphhealth.Diagnostic`. Trước direct iteration, bảo đảm consumed strings có cùng semantics sau JSON round-trip; một cách compliant là guard `SourceSiteID` và `Note` bằng UTF-8 validity và fallback toàn property nếu guard không đạt. Mọi type khác tiếp tục chạy exact marshal/unmarshal block hiện tại.

Có thể inline logic hoặc dùng tối đa một private non-exported read-only helper trong cùng `outcome.go`. Không đổi signature/caller. Planner/Coder có thể chọn coding form tương đương nếu nó chứng minh đầy đủ architecture rule và không mở rộng owner.

### Rationale

Canonical resolution path đã trả diagnostic property về exact typed representation trước projection. JSON round-trip hiện tại không cung cấp thêm semantic authority cho canonical slice; nó chỉ allocate, encode và parse lại data mà helper đọc read-only. Direct consumption loại đúng selected residual mechanism trong normal production path, trong khi compatibility fallback bảo toàn base graphs/legacy dynamic shapes và exact current failure semantics.

## Exact allowed owner boundary

### Production owner duy nhất

- `internal/resolution/outcome.go`:
  - body của `resolutionOutcomeDiagnosticSites`;
  - tối đa một private non-exported helper chỉ phục vụ canonical typed-read/fallback decision của helper này;
  - import tối thiểu nếu cần cho equivalence guard.

Không đổi signature của `resolutionOutcomeDiagnosticSites` hoặc `projectResolutionOutcomes`. Call site trong `projectResolutionOutcomes` là preserve-only ngoài mechanical no-op formatting nếu gofmt yêu cầu; không được thay algorithm/coordinator của nó.

### Test owner, chỉ sau production đúng

- Tạo duy nhất `internal/resolution/outcome_diagnostic_sites_test.go` cho D004-specific coverage.

Mọi existing test file là run-only. Nếu implementation cần sửa existing test hoặc một production file/symbol khác, STOP và trả Main cho fresh architecture; Planner/Coder không được tự broaden.

### Preserve-only surfaces

- `resolutionOutcomeCollector`, `newResolutionOutcomeCollector`, `record`, `finalize`, `marshalResolutionOutcome`, outcome validation/clone/construction và immediate diagnostic emitters;
- outcome map/encoding, relationship evidence, two reference-index projections và resolved-site logic trong `projectResolutionOutcomes`;
- `internal/resolution/resolve.go`, `emit.go`, `types.go`, export-binding machinery và all A001-A005 owners;
- `internal/graphhealth/diagnostics.go`, `policy.go`, `Diagnostic`, `DiagnosticPropertyKey` và diagnostic normalization/policy;
- `internal/graph/types.go`, graph/public schemas, Graph JSON, Ladybug/native persistence/readback, graph-accuracy/http/API/CLI readers;
- analyzer/CLI instrumentation, A00x script/overlay contract, target repositories, D001-D003, D005-D017, P3 và Child 07.

## Expected observable gain

Không dự đoán numeric speedup hoặc percentage.

Expected observable effect trên mỗi target riêng:

- D004 wall thấp hơn vì canonical diagnostic properties không còn trải qua graph-wide JSON marshal/unmarshal;
- parent `resolution` giữ được direct D004 reduction;
- analyzer và process giữ được benefit, không chuyển cost sang sibling operation;
- cumulative JSON CPU/allocation dưới `resolutionOutcomeDiagnosticSites` giảm; đây là supporting cause/resource observation, không thay wall acceptance.

Node scan, status lookup và exact payload parity vẫn còn; direction không tuyên bố loại toàn bộ D004 cost.

## Required invariants

### Resource và lifecycle

- Canonical fast path dùng `O(1)` extra retained state ngoài borrowed slice header/loop locals; không copy diagnostic objects hoặc retained Note bytes.
- Không cache, map phụ theo node/site, global/cross-run state, goroutine, lock, I/O, serialization sidecar, flush, finalizer hoặc persistence lifecycle mới.
- Không mutate/append/sort diagnostic slice hay graph property; ownership của slice vẫn thuộc graph.
- Compatibility fallback có cùng allocation/error behavior như current code cho noncanonical/JSON-unstable property.

### Failure và fail-closed

- Missing property vẫn skip.
- Noncanonical property vẫn dùng exact current JSON marshal/unmarshal conversion and errors; không silently drop, coerce hoặc accept unsupported shapes.
- Canonical direct path chỉ được dùng khi conversion cannot change the consumed values; nếu không chứng minh được thì fallback.
- Duplicate outcome, missing outcome, non-resolved reference carrier, resolved/unresolved overlap và payload drift errors giữ nguyên condition/message class và return timing relative to parity validation.
- Không giảm validation workload để tạo speedup.

### Determinism, order và mutation

- Giữ exact graph-node traversal, diagnostic slice traversal, SourceSiteID membership và error-first ordering.
- Không reorder diagnostics, relationships, references, evidence hoặc outcomes.
- Repeated execution trên cùng input tạo byte-identical graph/output; input graph/property không bị mutate bởi D004 read path.

### Public output, persistence và readers

- Exact nodes, relationships, IDs, labels, properties, diagnostic fields/count/order, outcome bytes, Evidence.Note bytes và reference-index carriers không đổi.
- Exact in-memory graph, canonical Graph JSON, stdout/stderr, Ladybug/native logical readback, metadata và downstream graph-health/graph-accuracy/API/CLI behavior không đổi.
- Không schema, exported signature, persisted field hoặc compatibility reader change.

## Production-first và tests-after-production boundary

Future Planner phải dịch theo exact order:

1. Future sole Coder chạy fresh required repo gate ngay trước edit: `anvien --help`, một `anvien analyze --force`, file-detail `internal/resolution/outcome.go`, exact upstream impact cho mọi edited symbol; ghi HIGH/CRITICAL đầy đủ như scope warning.
2. Implement production-only fast-path/fallback trong allowed owner. Inspect diff và STOP nếu có owner/signature/contract drift.
3. Chỉ sau production đúng mới tạo `internal/resolution/outcome_diagnostic_sites_test.go`.
4. Focused D004 tests phải chứng minh:
   - canonical `[]graphhealth.Diagnostic` trả exact same site membership/status/payload result và không mutate graph/slice;
   - untracked diagnostic bị ignore;
   - resolved-plus-diagnostic overlap fail closed;
   - mismatched `Note` fail closed;
   - multiple nodes/diagnostics giữ traversal/error-first behavior;
   - legacy `[]any`/`map[string]any` property đi fallback và có exact current result;
   - unsupported/unmarshalable property giữ current error;
   - JSON-unstable typed consumed strings không được đi unchecked fast path và phải match legacy round-trip behavior.
5. Chạy holder/lock/process clearance theo repo rule, rồi canonical full build phải exit `0` trước mọi test execution.
6. Sau build PASS, chạy focused new test và run-only regressions ít nhất:
   - `TestDiagnosticAppenderMatchesLegacySemantics`;
   - `TestP6C3FinalOutcomeStatusPrecedenceAndCarriageMatrix`;
   - `TestResolveAttachesSourceBackedUnresolvedDiagnostics`;
   - `TestP6DStructuredOutcomeProjectionPreservesGraphJSONAndHealthParity`;
   - `TestP6C3AnalyzeResultPreservesFinalOutcomesAndGraphCarriage`;
   - `TestP6C3AnalyzeCapabilityOutcomesRetainAcceptedAuthorityStatus`;
   - `TestP6C3ResolutionOutcomesPreserveGraphJSONAndLadybugParity`;
   - `TestP6C3NativeLadybugResolutionOutcomeReadback`;
   - full packages `internal/resolution`, `internal/graphhealth`, `internal/analyze`, `internal/lbugload`, `internal/lbugnative`.
7. Known preserve-only golden phải được ghi đúng sự thật, không gọi package PASS; bất kỳ new/changed failure nào block D004.
8. Build/validation xong mới freeze candidate và chạy target measurements; detect/stage/commit vẫn locked đến post-measurement Supervisor và Main disposition.

## Exact two-target measurement contract

Mỗi target dùng chính accepted A003 packet của nó làm `before`; không rerun accepted before, không average/combine, không dùng CPU/build/test time thay wall.

Candidate identity: accepted A003 checkpoint + đúng D004 production bytes trong `outcome.go` + unchanged accepted 17-child instrumentation/overlay/native/runtime/build contract. Test source không được ảnh hưởng executable behavior. Mỗi target launch đúng một candidate run với cùng workload/options/exclusion như accepted A003.

### Cheapapp

- Before: D004 `4.652523600 s`; parent `20.472602300 s`; analyzer `93.531974900 s`; process `95.630648200 s`.
- Denominator: `graphNodes=36575; graphRelationships=65045; outcomes=86742; referencesBySourceScope=25062; referencesByTargetDef=25062; workUnits=238486`.

### Restaurant Manager

- Before: D004 `5.995737900 s`; parent `20.850792800 s`; analyzer `98.020546700 s`; process `101.096911900 s`.
- Denominator: `graphNodes=66421; graphRelationships=112049; outcomes=186251; referencesBySourceScope=49259; referencesByTargetDef=49259; workUnits=463239`.

### Packet requirements per target

- child D004 -> parent `resolution` -> analyzer -> process wall, recorded separately;
- complete `30/30` operation rows, ordered `17/17` OP001 child rows, exact denominators, child-sum/parent residual conservation và overlap `0`;
- exact files/parser/workload, graph nodes/relationships, resolution metrics, diagnostics/outcomes, SourceSiteIDs, complete ordered Evidence, reference indexes;
- byte-identical canonical Graph JSON, stdout và stderr; logical Ladybug/native DB readback and affected readers exact;
- `startAllocBytes`, `endAllocBytes`, `maxObservedSys` và any exact D004 allocation/JSON-cause observation available without new broad instrumentation;
- target HEAD/status, executable/DLL/provenance, overlay/candidate-source hashes và one-launch/exit identity.

Binding campaign `KEEP` rule không đổi: D004, parent và process wall phải thấp hơn trên từng target và later exact Supervisor phải PASS. Analyzer là mandatory no-cost-transfer boundary: bất kỳ unexplained analyzer regression nào STOP/blocks promotion dù child/parent giảm. Secondary CPU/memory improvement không bù cho wall failure.

## A005 non-duplication

Rejected A005 direction là record-time canonical outcome-byte ownership: sửa collector `record/finalize`, tạo retained SourceSiteID-to-JSON sidecar/finalized bundle, đổi projection consumption và narrow `resolve.go` wiring. Candidate đó là `SUPERVISOR_PASS / NO_KEEP / ROLLBACK_COMPLETE` và tuyệt đối không được tái dùng.

D004 direction này:

- không đổi khi nào/how many times `marshalResolutionOutcome` chạy;
- không retain/reuse canonical outcome bytes ngoài current `encodedBySourceSite` projection map;
- không đổi collector, finalized outcome lifecycle, `resolve.go` hoặc immediate diagnostic construction;
- chỉ loại redundant conversion của graph diagnostic property đã canonical typed ngay trong `resolutionOutcomeDiagnosticSites`.

Controlled perturbation chứng minh non-duplication: A005 làm sibling `marshalResolutionOutcome` biến mất, nhưng selected helper vẫn material `2.52 s` Cheapapp / `3.34 s` Restaurant. Vì vậy D004 xử lý residual còn lại, không phục hồi hay biến thể A005.

## Alternatives considered

### Unconditional `[]Diagnostic` assertion

Đây là code ngắn hơn nhưng không safe. `NodeProperties` là `map[string]any`, base graph/decoded/legacy inputs có supported `[]any`/`map[string]any` representations, và current appender contract chủ ý normalize chúng on first touch. Bỏ fallback sẽ panic, silently skip hoặc đổi fail-closed behavior. JSON round-trip cũng canonicalize invalid UTF-8; direct use không có equivalence guard có thể đổi `SourceSiteID` lookup. Vì vậy không có strictly simpler safe variant hơn selected guarded fast-path + exact fallback.

### Emitter-owned diagnostic site sidecar

Rejected vì tạo duplicate carrier authority và có thể PASS sidecar trong khi actual graph property drift/mutation sai. Nó mở rộng `emitter/resolve.go` ownership, resource lifecycle và parity surface nhưng lại né chính invariant phải kiểm tra actual graph carrier.

### Export/reuse graphhealth normalization reader

Rejected cho attempt này vì mở rộng production owner sang graphhealth và current normalization có thể re-enter structured diagnostic policy/decoder work. Selected mechanism không cần thay policy/normalization contract; one-file read-only fast path nhỏ hơn và dễ rollback hơn.

## Rollback và mandatory STOP

Exact rollback chỉ gồm D004 hunk trong `resolutionOutcomeDiagnosticSites`, optional single private helper/import, new D004 test file và frozen candidate/overlay packet của attempt. Không rollback broad, không chạm accepted A003/A001-A005/protected bytes.

STOP và trả Main nếu xảy ra bất kỳ điều nào:

- cần production file/symbol khác hoặc sửa existing test file;
- cần đổi graphhealth `Diagnostic`, property key, appender/normalizer/policy, graph types, collector, outcome marshal lifecycle, `projectResolutionOutcomes` algorithm/signature hoặc `resolve.go`;
- bỏ/loosen compatibility fallback, silently coerce/drop property, hoặc không chứng minh JSON-equivalent direct path;
- tạo sidecar/cache/registry/cross-run state, mutate graph/slice, reorder output hoặc đổi error-first behavior;
- public/persisted shape, canonical bytes, Graph JSON, Ladybug/native/readers, determinism, failure hoặc lifecycle drift;
- reuse A005 canonical outcome-byte direction;
- canonical build hoặc required test có new/changed failure;
- target packet thiếu, mixed/incomparable, denominator/workload/overlay mismatch, hoặc không giữ target separation;
- D004/parent/process không giảm trên một target, analyzer có unexplained regression, hoặc resource behavior vượt bounded `O(1)` extra retained state;
- cần profile/source discovery/measurement mới trước Planner ngoài exact validation/measurement contract đã ghi.

## Evidence limits

- Không có exact diagnostic-bearing node count, JSON round-trip invocation count hoặc byte count; report không bịa chúng.
- Profile samples không cho exclusive helper wall hoặc numeric saving.
- Direction chứng minh loại conversion trên canonical current production representation; compatibility fallback có thể vẫn chạy cho rare noncanonical properties, và graph-node scan/parity work vẫn tồn tại.
- Chỉ candidate measurement mới chứng minh D004/parent/analyzer/process effect. Architecture report không claim speedup, correctness acceptance, `KEEP`, attempt result hoặc terminal child state.

## Handoff

- Exact verdict: `ARCHITECT_D004_READY_FOR_PLANNER`.
- One selected direction: guarded direct consumption of canonical typed diagnostic graph properties, with exact legacy/noncanonical JSON fallback and unchanged actual-carrier parity checks.
- Commit reference: none; Main owns checkpoint. Không stage hoặc commit trong lane này.
- Next owner: Main Orchestration 01a0421e-a38e-7192-8e98-8a09fa72f04d.
