# Báo cáo Supervisor: Đánh giá mức độ giải quyết vấn đề của Anvien Graph Identity/Resolution Multi-plan

Verdict: REJECT

## Metadata

- Thời điểm review: 2026-07-28 22:24:05 +07:00
- Reviewer: gpt-5-codex
- Repository: E:/Anvien
- Problem authority: reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien
- Plan được review: docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan
- Phạm vi: semantic adequacy và execution-readiness của roadmap + 7 child plan sets; không sửa plan, không sửa production code
- HEAD khi review: c637152fc211992105e4067daedf87a377da8723

## Claim được đánh giá

Claim cần kiểm là: bộ multi-plan hiện có thật sự đủ để giải quyết toàn bộ problem report, theo đúng thứ tự identity -> binding -> export -> module/re-export -> ambient/diagnostics, với các gate downstream, migration, persistence và target acceptance.

Kết luận này không đồng nhất “plan có nhiều coverage” với “vấn đề đã được sửa”. Việc implementation/evidence còn pending tự nó chỉ chứng minh chưa có remediation proof; verdict REJECT ở đây chủ yếu dựa trên các lỗ hổng contract, dependency và acceptance có thể cho phép campaign đóng khi yêu cầu gốc vẫn chưa được chứng minh.

## Cơ sở đã đọc và kiểm chứng

- Đọc toàn bộ 573 dòng problem report, toàn bộ roadmap 215 dòng, 7 child plan, 7 actual-status, 7 evidence ledger, 7 benchmark ledger (29 file Markdown trong multi-plan).
- Đọc toàn bộ reader matrix 213 dòng mà P2/P7 dùng làm denominator.
- Đối chiếu với AGENTS.md, CLAUDE.md, working rules đính kèm, planner skill, các template planner và supervisor skill.
- Ba audit độc lập được dùng như kiểm tra chéo:
  - audit_identity_binding: child 01-03;
  - audit_export_ambient: child 04-07;
  - audit_governance: lifecycle, ledger và authority.
  Các kết luận cuối cùng dưới đây đều đã được tự kiểm tra lại trong file và dòng nguồn; không coi verdict của subagent là bằng chứng tự thân.
- Kiểm tra cấu trúc trên đĩa: 7 child plan, 98 implementation slice (11+42+17+15+4+6+3), 7 P0; 98 implementation checkbox và 21 Pn closure checkbox vẫn chưa hoàn thành.
- Quét cross-child evidence: roadmap có yêu cầu E2-PNC-HANDOFF1/E2-PNC-NEXTSTATUS1, nhưng 28 child file không khai báo hai ID này. Các ID chỉ xuất hiện trong roadmap.
- Worktree trước khi tạo báo cáo chỉ có problem artifact chưa tracked của người dùng; không có production diff.

## Phần coverage tốt (chỉ là design intent)

| Cụm problem | Coverage hiện có | Giới hạn |
|---|---|---|
| Identity | Tách Declaration/Symbol, range/selectionRange, deterministic tuple, RelationshipID, collision và shadow-v2 được mô tả khá đầy đủ trong child 01. | Exact manifest/handshake metadata chưa khớp đủ với contract problem; chưa có implementation proof. |
| Persistence/migration | Child 02 có reader matrix S0-S11, old-reader handshake, staging/CAS, registry/group/cache/embedding, lease/GC và fault matrix. | Matrix hiện chỉ là planned seed; các guard/evidence runtime đều pending. |
| Binding | Child 03 có recursive array/object/rest/default/holes/computed paths, declaration-vs-assignment, variable/parameter/catch/loop contexts và target 6/6. | Dependency và acceptance không đóng hết các symptom .map(), type-independent emission và structured unsupported diagnostic. |
| Export | Child 04 có ExportFact, meaning lanes, direct/default/alias/type-only/star/namespace syntax và target 21/21. | ModuleRequestFact/ImportBindingFact và derived public-export ownership chưa hoàn tất; một slice vừa ngoài scope traversal vừa đòi derived values. |
| Module/re-export | Child 05 có ModuleRef, export table, alias/proof, cycle, ambiguity, dedupe, budgets, no-global-fallback và target barrel 2/2. | Hai acceptance bắt buộc của problem về zero-physical-declaration barrel và path-resolution count chưa được nêu thành gate. |
| Ambient/diagnostics | Child 06 có profile, embedded catalog, local package boundary, security/budget, lazy external materialization, immutable outcomes và graph-health projection. | Thiếu capability/completeness modes, một số nguồn declaration và downstream external isolation. |
| Final validation | Child 07 có determinism, version/fault matrix, target boundary, Docker/Playwright, S0-S11 parity và năm-run performance. | Parity có thể đồng nhất một kết quả sai; không thay thế semantic oracle cho context/impact/rename/groups. |

## Findings blocking

### B1 — P3-C có dependency gate sai, cho phép projection trước khi hoàn tất loop bindings

Problem report yêu cầu walker bao phủ for-of/for-in và mọi leaf phải trở thành binding thật (problem report:195-212). Child 03 đặt P3-B2A Integrate for-of/for-in ở trước P3-C (child 03 plan:296-303), và P3-B2A là slice sở hữu loop contexts (child 03 plan:515-558). Tuy nhiên Implementation Gate của P3-C chỉ yêu cầu P3-A/B/B1/B2, bỏ P3-B2A (child 03 plan:560-592).

Roadmap cũng đã ghi nhận nguyên văn đây là legacy omission và không sửa trong mechanical copy (roadmap:104-109). Vì vậy một agent có thể hợp lệ mở graph projection khi for-of/for-in facts chưa PASS. Đây là lỗ hổng thực thi trực tiếp, không phải vấn đề trình bày.

### B2 — Handoff/successor-freshness contract của roadmap không được bind vào child closure

Roadmap bắt buộc child không-terminal phải refresh actual-status của successor, append refresh-log, cập nhật next actions/work steps, và phát qualified E2-PNC-NEXTSTATUS1; closure phải phát qualified E2-PNC-HANDOFF1 (roadmap:33-36,86-102).

Pn-C của child 01-07 chỉ là close generic (ví dụ child 06 plan:634-645 và child 07 plan:516-527). Chúng không có bước, Implementation Gate, Acceptance hay Evidence Target cho:

- exact successor actual-status refresh;
- refresh-log row và affected next-action/work-step update;
- stale-successor block;
- E2-PNC-NEXTSTATUS1;
- E2-PNC-HANDOFF1;
- terminal child 07 roadmap/campaign refresh.

Closure Evidence của các ledger chỉ ghi Reserved for final Supervisor, commit, successor-status, and handoff evidence (child 06 evidence:167-169; child 07 evidence:193-195), không phải evidence declaration. Do đó roadmap yêu cầu child 07 chờ sáu qualified handoff (roadmap:92), nhưng child plans không có gate máy đọc được để tạo hoặc kiểm tra các handoff đó.

### B3 — Quy tắc closure tạo deadlock với P7

Roadmap yêu cầu child 01-06 close/commit/handoff tuần tự trước khi child 07 mở (roadmap:34-37,86-94), trong khi cả bảy benchmark ledger đều nói “A final plan cannot close while any required P7 performance row remains pending” (ví dụ child 01 benchmark:90, child 06:92, child 07:103).

Đọc theo literal rule, child 01-06 không thể đóng local plan vì P7 thuộc child 07; nhưng child 07 lại không thể mở nếu các child trước chưa đóng. Ít nhất cần phân biệt rõ local child closure với campaign/release closure. Hiện ambiguity này có thể làm execution dừng vĩnh viễn hoặc khiến agent tự ý bỏ qua một rule.

### B4 — Binding acceptance không đóng các acceptance bắt buộc của problem

Problem report yêu cầu đồng thời: emission không phụ thuộc type inference; pattern chưa hỗ trợ phải có structured extraction diagnostic/metric; sáu .map() không còn ResolutionGap; nested shadowing trỏ đúng Symbol; và không double-count import (problem report:210-223).

Child 03 có matrix/walker test và import gate (child 03 plan:315-365), nhưng:

- P3-A không yêu cầu test/metric cho unsupported-pattern diagnostic;
- P3-B non-goal “no fake per-leaf type inference” không chứng minh declaration vẫn được emit khi inference thất bại (child 03 plan:370-377);
- P3-C2 chỉ ghi 6/6 fields, same-name IDs và row parity, không có assertion “sáu map sites / 0 ResolutionGap”, cũng không có shadowing oracle rõ ràng (child 03 plan:1055-1095).

Vì vậy target có thể đạt 6/6 fact fields nhưng vẫn còn symptom downstream mà problem report yêu cầu loại bỏ.

### B5 — ModuleRequestFact/ImportBindingFact có dependency nhưng không có producer và derived export state không có owner rõ

Problem report đề xuất ba fact độc lập: ModuleRequestFact, ImportBindingFact và ExportBindingFact (problem report:241-266). Child 04 đặt module-request/provider ở inspect-only và không có slice tạo/bind hai fact đầu (child 04 plan:314-330; ownership table:197-198). Child 05-A lại yêu cầu accepted P4 module-request facts và gate chỉ mở khi facts đó available (child 05 plan:325-374). Audit toàn bộ multi-plan không tìm thấy ImportBindingFact.

Đồng thời ExportFact của child 04 chứa reachableThroughBarrel và publicApi (child 04 plan:238-242), nhưng P4-C tuyên bố re-export traversal out of scope rồi vẫn đòi test direct/export-entry/reachable/publicApi (child 04 plan:469-486). P5 mới sở hữu traversal (child 05 plan:380-445), nhưng không có slice/owner/metric cho ba count mà problem tách riêng: directExportedDefinitionCount, resolvedExportEntryCount, publicApiSymbolCount (problem report:286-305).

Hậu quả là có dependency producer bị hở và không có một source of truth xác định ai tính derived public API. P4 có thể được đóng với direct 21/21 nhưng không chứng minh được resolved export surface và package public API.

### B6 — Ambient declaration universe thiếu capability/completeness contract và thiếu một số nguồn declaration

Problem report yêu cầu nguồn repo declarations, project-owned .d.ts, installed packages, stdlib, intrinsics, ambient modules/global augmentations; outcome phải giữ confidence/completeness mode; metadata phải công bố exact/structural/degraded (problem report:401-475).

Child 06 mô tả outcome với target/stage/severity/actionability/retryability và status matrix (child 06 plan:147-165,589-613), nhưng không định nghĩa field/enum/gate cho exact|structural|degraded hoặc confidence/completeness. P6 tests nêu stdlib, package present/absent và triple-slash, nhưng không có slice/acceptance cụ thể cho project-owned .d.ts, ambient modules/global augmentations và merge/parsing semantics (child 06 plan:273-287,306,361,416,461).

Quét toàn bộ 28 child file cho thấy không có occurrence của capability enum, completeness mode hay confidence/completeness contract. Việc có origin ambient_module trong ExternalSymbolRef không thay thế loader và acceptance cho nguồn đó.

### B7 — ExternalSymbol downstream isolation bị bỏ sót

Problem report quy định external symbol mặc định không tham gia in-repo rename và có thể bị loại khỏi impact/process traversal trừ khi user chọn include_external (problem report:500; shared gate:517).

Không child nào định nghĩa include_external hoặc contract tương đương. P7-C chủ yếu kiểm canonical parity, field differences và orphan refs trên S0-S11 (child 07 plan:442-491); reader matrix các dòng context/impact/rename/groups chủ yếu là generation/protocol guards (reader matrix:50-60,118-143), không phải semantic oracle cho external exclusion hay same-name selection.

Do đó mọi surface có thể cùng trả một symbol sai mà vẫn đạt differing_records == 0.

### B8 — Exact manifest/handshake không chứa đủ metadata mà problem yêu cầu

Problem report yêu cầu graph metadata gồm graphSchemaVersion, identitySchemaVersion, scopeIRSchemaVersion, graphGeneration, analyzerVersion, columnEncoding và source/config fingerprint (problem report:104-116).

Plan có câu khẳng định positionEncoding/analyzer build được persist/check (child 02 plan:150), nhưng active manifest và wire contract được chốt tại child 02 plan:152-158 chỉ liệt kê protocol/build/schema/scope/generation/configHash/catalogHash; không có positionEncoding/columnEncoding, writer analyzerVersion hay source fingerprint. P2-A còn gọi handshake fields hiện tại là fixed (child 02 plan:318-352).

Reader không thể validate một invariant khi dữ liệu cần validate không nằm trong manifest/wire contract. Đây là contract mismatch trước khi cutover, không thể để implementation tự diễn giải.

### B9 — Hai acceptance module resolver trong problem chưa được chuyển thành gate đo được

Problem report yêu cầu barrel không có physical declaration vẫn có export surface và path-resolution count không bị export resolution thay đổi (problem report:389-397). P5 semantic vector có cycles/ambiguity/aliases/meanings (child 05 plan:298-318), P5-D có 2/2 calls/proof/no-false-gap (child 05 plan:490-539), nhưng không có assertion hoặc metric nào cho hai yêu cầu trên.

Đây là lỗ hổng khiến resolver có thể đúng ở hai fixture target nhưng vẫn sai ở barrel rỗng hoặc làm thay đổi module path accounting.

### B10 — Promise/Math có expected outcome mâu thuẫn giữa P6 và P7

P6-C3/P6-D yêu cầu Promise/Math phải là resolved_external và không bao giờ resolved_intrinsic (child 06 plan:523-583). Nhưng top-level P7 acceptance lại cho phép “resolved external/intrinsic outcomes or explicit external-capability outcomes” (child 07 plan:291-304). Problem report chỉ cho external hoặc external-capability status, không cho phép coi Promise/Math là intrinsic (problem report:459-500,521-527).

Gate cuối yếu hơn gate upstream, nên một implementation sai có thể được P7 chấp nhận.

### B11 — One-file/one-responsibility chưa đủ granularity để Supervisor enforce

Problem report yêu cầu mỗi job có bảng File / Trách nhiệm duy nhất / Được phép liên kết / Không được chứa, và Supervisor phải từ chối slice vi phạm (problem report:557-573).

Các child chỉ có một owner map chung ở cấp campaign (ví dụ child 06 plan:185-221). Slice không có bảng exact production/test/generated ownership cho chính job đó; nhiều nơi chỉ để dedicated owner, narrow adapter hoặc allowlist sẽ được xác định sau pre-flight. Structural authoring audit vì thế chưa chứng minh được rule ở granularity mà problem yêu cầu.

### B12 — Acceptance evidence của các adapter không khớp evidence ledger

Ví dụ child 04 P4-C1A acceptance chỉ khai báo E4-P4C1A-REVIEW1 (child 04 plan:601-610), trong khi evidence ledger yêu cầu impact/source/build/test/review/detect/commit (child 04 evidence:161-169). Pattern này lặp ở các adapter C1A-C1I.

Một slice có thể bị đánh dấu đủ chỉ bằng review record dù các bằng chứng build/source/detect/commit mà rule chung yêu cầu chưa được bind vào Acceptance. Đây là traceability gap có thể làm closure giả.

## Traceability verdict theo nhóm problem

| Nhóm bắt buộc trong problem | Kết quả đối chiếu | Lý do chưa đủ |
|---|---|---|
| Identity/Declaration-Symbol | Coverage mạnh nhưng chưa hoàn tất | B8 và toàn bộ implementation proof còn thiếu. |
| Binding-pattern | Chưa đủ | B1 và B4. |
| Export facts/visibility | Chưa đủ | B5 và B12. |
| Module/re-export resolver | Coverage thuật toán tốt nhưng chưa đủ | B5, B9 và B10. |
| Ambient/external/diagnostics | Chưa đủ | B6 và B7. |
| Persistence/migration/downstream/capacity | Chưa đủ | B3, B7, B8; các benchmark/evidence thực thi còn pending. |

## Phân biệt design coverage với proof thực tế

Roadmap vẫn ghi candidate, legacy plan còn active, authority cutover và production implementation chưa được phép (roadmap:4,14-23). Campaign completion vẫn yêu cầu child 01-06 implemented/validated/committed và child 07 hoàn tất final acceptance (roadmap:205-215). Evidence ledger của child 06 và child 07 đánh dấu mọi implementation/validation/commit evidence là pending (child 06 evidence:149-169; child 07 evidence:149-195). Đây là lý do không thể ghi nhận problem đã được giải quyết trong repository hiện tại.

Actual-status của các child 02-07 cũng vẫn ghi blocked pending predecessor handoff/owner authority và kết luận “only next permitted slice is P1-A” (ví dụ child 02 actual:5,278; child 06 actual:5,264; child 07 actual:5,437), trong khi next-phase table của chúng là P2/P6/P7. Nội dung stale/copy này cần được refresh trước khi agent khác dùng làm current execution state.

Bản Supervisor trước đó chỉ xác nhận bounded seven-child authoring/copy closure và nêu rõ không bao gồm authority cutover, production implementation hay target validation. Không được dùng bản đó làm bằng chứng rằng các defect trong problem đã được sửa.

## Không chạy trong review này

- Không chạy build, test, Docker, browser/Playwright, target analyze hoặc production runtime: implementation chưa được owner mở và review này chỉ đánh giá document plan.
- Không chạy Anvien graph mutation/query: không có code/topology claim cần thêm vào phán quyết; các graph counts trong ledgers được coi là historical/document evidence, không phải remediation proof.
- Không sửa roadmap, child plan, ledger, source hay target. Report này là artifact duy nhất được tạo trong lượt review.

## Disposition bắt buộc trước khi yêu cầu review lại

Không thực hiện các chỉnh sửa dưới đây trong lượt này; đây là điều kiện để owner/plan author xử lý ở lượt được ủy quyền:

1. Sửa dependency P3-C để bắt buộc P3-B2A; thêm các target/metric gate về .map() zero-gap, type-independent emission, unsupported-pattern diagnostic và shadowing.
2. Biến handoff/successor refresh thành evidence targets và gates thực sự trong từng Pn-C; khai báo qualified IDs; tách local child closure khỏi campaign/P7 release closure để loại deadlock.
3. Bổ sung producer/owner rõ cho ModuleRequestFact, ImportBindingFact, ba export counts và package public API; tách P4 syntax facts khỏi P5 traversal-derived state.
4. Bổ sung capability/completeness modes, mọi nguồn DeclarationUniverse được problem yêu cầu, parser/merge acceptance và external-symbol isolation cho context/impact/rename/process/groups.
5. Chốt lại manifest/handshake với analyzerVersion, position/column encoding và source fingerprint; thống nhất Promise/Math expected outcome ở mọi child.
6. Bổ sung hai resolver acceptance bị thiếu, per-job ownership tables và full evidence IDs cho adapter slices.
7. Sau khi plan được sửa và owner chấp thuận, refresh tất cả P0/actual-status từ current HEAD và chạy lại Supervisor review trước khi mở implementation.

## Integrity statement

Report này không thay đổi bất kỳ plan nào. Verdict duy nhất của review đối với claim “multi-plan hiện đủ để thật sự giải quyết problem” là REJECT.

