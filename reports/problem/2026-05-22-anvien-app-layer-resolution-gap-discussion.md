# Tong hop thao luan: App Layer, Resolution Gap, UI lens va query health

Date: 2026-05-22
Scope: discussion only, not an implementation plan
Status: expanded discussion record

## 1. Muc dich cua tai lieu

Tai lieu nay ghi lai phan thao luan ve van de graph hien tai chua du ro nghia cho repo lon.

Day khong phai la plan trien khai. Muc dich la gom lai cac nhan dinh da ban, cac gia thuyet can kiem tra, va huong thiet ke co kha nang giai quyet van de truoc khi tao bo plan rieng.

Trong thao luan, trong tam khong phai la dead code. Van de chinh la `unresolved_reference` sau khi da tach khoi `unknown_connectivity`, va cach lam sao de no tro thanh du lieu co the nhin, loc, gom cum, va hieu duoc tren Web UI.

## 2. Boi canh da dan toi thao luan nay

Truoc do Graph Health da tung gom `unresolved_reference` vao `unknown_connectivity`. Cach do sai ve y nghia:

- `unknown_connectivity` la van de topology.
- `unresolved_reference` la van de resolution/analyzer.

Sau khi tach, topology sach hon:

- node co edge ket noi van giu topology that, vi du `connected`;
- `unknown_connectivity` chi con dung cho topology that su khong xac dinh;
- diagnostic unresolved van duoc giu rieng.

Tuy nhien sau khi tach thi lo ra van de that su: `unresolved_reference` van con lon va van chua co cach hien thi/phan loai du ro de nguoi dung hieu ngay no la gi.

Trong graph quan sat luc thao luan:

- total nodes: khoang `22010`;
- counted semantic relationships: khoang `26906`;
- `unknown_connectivity`: `0` sau khi tach;
- diagnostic unresolved occurrences: khoang `51232`;
- diagnostic buckets: khoang `8880`;
- buckets theo fact family co nhieu nhat: `call`, `type-reference`, `access`, `heritage`.

Con so nay cho thay viec tach topology va diagnostic la can thiet, nhung chua du. Unresolved van la mot lop van de rieng can duoc nang thanh inventory/lens co y nghia.

## 3. Van de goc voi node type hien tai

Graph hien tai phan loai node chu yeu theo hinh thai symbol:

- Function
- Method
- Struct
- Interface
- Variable
- Property
- File
- Folder
- Package
- Process
- Community
- Section

Cach phan loai nay cho biet node la loai symbol gi, nhung chua cho biet node thuoc phan nao cua san pham.

Vi du `Function` co the la:

- backend resolver;
- frontend React helper;
- CLI launcher function;
- test helper;
- contract generator;
- provider extractor;
- runtime/session function.

Neu tat ca deu chi la `Function`, thi doc vao graph chua the thay ro van de dang nam o backend, frontend, shared contract, test, docs hay config.

Do do khi gap `unresolved_reference`, chinh du lieu graph/API moi chi cho biet:

```text
node nay co reference chua resolve
```

Nhung doc vao graph chua tra loi duoc:

- no thuoc Backend hay Frontend?
- no nam trong nhom chuc nang nao?
- target chua resolve co ve la callable, member, type, external symbol, builtin hay test helper?
- no co lam topology mat do tin cay hay chi la diagnostic khong quan trong?
- no co phai analyzer gap trong repo hay chi la reference ngoai repo?
- no co phai mot nhom van de lon trong mot app layer cu the hay khong?

Web UI chi la lop render va thao tac tren ket qua da co. Neu graph/API chua co cac lop phan loai nay, UI khong the hien thi cau tra loi dung mot cach on dinh.

## 4. `unknown_connectivity` da tach dung, nhung khong phai ket thuc van de

Nhan dinh da chot:

`unknown_connectivity` da duoc tach dung khoi `unresolved_reference`.

Sau khi tach:

- `unknown_connectivity` chi con la topology that su khong xac dinh duoc;
- node co unresolved diagnostic nhung co edge ket noi van giu topology that;
- `unresolved_reference` chuyen ve dung vai tro diagnostic/resolution gap;
- viec "unknown" khong bi mat, ma duoc dat lai dung cho cua no.

Nhung `unresolved_reference` van la van de lon vi no bao hieu graph co the dang thieu edge.

Neu analyzer khong resolve duoc reference:

```text
Function A goi Function B
nhung analyzer khong tim duoc B
=> edge A -> B khong duoc tao
=> B co the bi nhin nham thanh no_incoming hoac true_isolated
```

Vi vay cac topology status nhu:

- `no_incoming`
- `no_outgoing`
- `true_isolated`
- `detached_component`

co the bi giam do tin cay neu gan chung co nhieu unresolved in-repo references.

Ket luan cua phan nay: khong dua `unresolved_reference` quay lai `unknown_connectivity`, nhung phai co mot lop `Resolution Health` de noi ro graph dang thieu thong tin o dau.

## 5. Unresolved reference hien dang duoc sinh nhu the nao

Trong code hien tai, `unresolved_reference` duoc sinh khi resolver thay mot reference trong ScopeIR nhung khong resolve duoc target.

Cac luong chinh:

- `resolveCall` gap call target khong resolve duoc;
- `resolveAccess` gap access target khong resolve duoc;
- `resolveTypeAnnotation` gap type target khong resolve duoc;
- `emitUnresolvedHeritageDiagnostics` gap heritage target/owner khong resolve duoc;
- `emitUnresolvedReference` tao diagnostic;
- `AppendDiagnosticToNode` attach diagnostic vao source node.

Dieu quan trong: hien no chi attach vao source node duoi dang diagnostic evidence. No khong tao mot graph entity rieng cho target chua resolve.

Dang hien tai:

```text
Source node A
  diagnostics:
    kind: unresolved_reference
    factFamily: call
    targetText: collector
```

Dang nay giu duoc evidence, nhung chua du de UI coi no nhu mot vat the co the gom cum/loc/to mau/noi canh.

## 6. Vi sao chi classification diagnostic la chua du

Truoc do da co huong them:

- builtin
- standard library
- test framework
- external library
- in-repo unresolved
- unclassified

Huong nay dung, nhung moi la classification cap diagnostic.

No van chua tra loi day du:

- source nam o app layer nao?
- source nam trong functional area nao?
- unresolved target nen hien nhu node nao tren UI?
- target co the suy luan vai tro gi?
- gap nay co nen hien trong graph nhu mot island rieng?
- gap nay co nen anh huong canh bao topology confidence khong?

Neu chi tach classification nhung cuoi cung UI van khong biet chung la gi, thi ve trai nghiem nguoi dung van gan nhu chua giai quyet duoc van de.

## 7. App Layer / Runtime Surface

De xuat them mot lop phan loai rieng: App Layer / Runtime Surface.

Lop nay khong thay the node type hien co. No la truc phan loai thu hai, dung de noi node thuoc phan nao cua san pham.

App layer de xuat:

- `backend`
- `api`
- `frontend`
- `cli_launcher`
- `shared_contract`
- `docs`
- `test`
- `config`
- `unknown`

Day la lop BE/FE thuc dung nhat va nen lam truoc vi:

- de suy ra bang path/rule hien co;
- de hien thi tren UI;
- de gom cum graph truc quan hon;
- de phan tich `unresolved_reference` theo boi canh san pham;
- de lam nen cho query/impact/context/detect-changes sau nay;
- de tach repo lon thanh cac vung de nhin hon.

Trong thao luan bo sung, `api` nen la mot App Layer chinh, khong chi la functional area con. Ly do:

- API la ranh gioi giua Frontend va Backend.
- Web UI phu thuoc truc tiep vao API response shape va contract.
- Nhieu loi co the nam o serialization/handler/response contract, khong phai o backend logic hay frontend render.
- Neu API bi gop chung vao backend, graph kho nhin ra truong hop backend logic dung nhung API tra thieu field, hoac frontend dung nhung API contract sai.

Vi vay App Layer nen co toi thieu:

```text
backend
api
frontend
```

voi API nam o giua ve mat layout va nghia san pham.

Vi du:

```text
Node type: Function
App layer: backend
Functional area: resolution
```

hoac:

```text
Node type: Function
App layer: frontend
Functional area: web_graph_ui
```

hoac:

```text
Node type: File
App layer: shared_contract
Functional area: contracts
```

Vi du API:

```text
Node type: Function
App layer: api
Functional area: graph_health_api
```

hoac:

```text
Node type: Function
App layer: frontend
Functional area: api_client
```

## 8. Rule suy luan App Layer ban dau

Ban dau khong can AI phuc tap. Co the dung rule-based theo path, node label, va context.

Vi du cho repo Anvien:

```text
anvien-web/src/**
=> frontend

anvien-web/test/**
anvien-web/e2e/**
=> frontend + test

internal/**
cmd/**
=> backend

internal/httpapi/**
=> api

anvien-web/src/services/backend-client.ts
=> frontend + api_client

anvien-launcher/**
=> cli_launcher

contracts/**
internal/contracts/**
=> shared_contract / api_contract

cmd/generate-web-contracts/**
=> shared_contract / api_contract

docs/**
reports/**
*.md
=> docs

*_test.go
test/fixtures/**
=> test

config files, package files, build scripts
=> config or cli_launcher depending path
```

Can tranh ep tat ca vao BE/FE. Node khong ro nen vao `unknown` hoac `mixed`, khong duoc nhan sai de graph dep gia tao.

## 9. Functional Area

Sau App Layer, can them lop nho hon de phan loai theo nhom chuc nang.

Vi du:

- `resolution`
- `graph_health`
- `query`
- `mcp`
- `web_graph_ui`
- `layout`
- `contracts`
- `providers`
- `runtime`
- `analyzer`
- `session`
- `launcher`
- `cli`
- `reporting`

Functional Area khong nen thay the App Layer. Thu tu dung:

```text
App Layer first, Functional Area second, Node Type third.
```

Ly do: neu node type truoc thi `Function` cua BE va `Function` cua FE van bi tron vao nhau. Neu App Layer truoc, nguoi dung thay ngay van de o BE hay FE.

## 10. Resolution Gap Entity

Khong nen de `unresolved_reference` chi la diagnostic text treo tren source node.

Nen nang thanh entity/filter ro rang:

```text
ResolutionGap / UnresolvedSymbol
```

Metadata can co:

- `sourceNode`
- `sourceAppLayer`
- `sourceFunctionalArea`
- `factFamily`: `call`, `access`, `type-reference`, `heritage`
- `targetText`
- `inferredTargetRole`: `callable`, `member`, `type`, `external`, `builtin`, `test`, `unknown`
- `actionability`: `analyzer_gap`, `review`, `non_actionable`
- `resolutionSource`
- `filePath`
- `startLine`
- `count`
- `note`

Quan trong: unresolved target khong nen bi gia vo thanh `Function` that neu chua chac. Neu chua resolve duoc thi nen la `UnresolvedSymbol` hoac `ResolutionGap`, roi gan `inferredTargetRole`.

Vi du:

```text
Source: resolveCall
Source app layer: backend
Functional area: resolution
Unresolved target: collector
Fact family: call
Inferred target role: callable
Actionability: analyzer_gap
```

Thong tin nay co y nghia hon nhieu so voi chi hien:

```text
unresolved_reference: collector
```

## 11. Suy luan target role

Target role co the suy ra tu `factFamily` va `targetText`.

Vi du:

```text
factFamily=call
=> inferredTargetRole=callable

factFamily=access
=> inferredTargetRole=member

factFamily=type-reference
=> inferredTargetRole=type

factFamily=heritage
=> inferredTargetRole=type
```

Sau do co the tinh tiep classification:

```text
targetText=len
=> builtin

targetText=testing.T
=> test_framework

targetText=fmt.Errorf
=> standard_library

targetText=uuid.New
=> external_library hoac external_symbol

targetText=collector
=> likely in-repo unresolved neu khong match builtin/std/external
```

Target role la "du doan co kiem soat", khong phai ket luan target da ton tai trong graph.

## 12. Resolution Health can duoc hien thi nhu mot lens rieng

Khong nen tron Resolution Health vao Topology Health.

Topology Health:

- `connected`
- `no_incoming`
- `no_outgoing`
- `true_isolated`
- `detached_component`
- `unknown_connectivity`

Resolution Health:

- resolved references
- unresolved non-actionable
- external unresolved
- in-repo unresolved / analyzer gap
- unresolved call target
- unresolved access target
- unresolved type target
- unresolved heritage target

UI can noi ro:

```text
Topology: connected
Resolution: unresolved call target
Actionability: analyzer_gap
```

hoac:

```text
Topology: no_incoming
Resolution confidence: degraded
Reason: nearby/source references could not be resolved
```

## 13. Layout UI moi: multi-ring theo App Layer

Hien tai graph co mot vong tron lon chua cac quan dao node theo mau/type. Huong moi nen chuyen sang nhieu macro-ring theo App Layer.

Toi thieu:

```text
Backend Ring
API Ring
Frontend Ring
```

Co the them:

```text
Shared / Contract Ring
Docs / Test / Config Ring
Unknown / Mixed Ring
```

Nguyen tac:

```text
App Layer first, Node Type second.
```

Tuc la:

1. Chia node vao cac vong lon theo App Layer.
2. Ben trong moi vong, chia tiep thanh cac dao nho theo node type/filter mau hien co.

Vi du:

```text
Backend Ring
- Function island
- Method island
- Struct island
- Package island
- Process island
- ResolutionGap island

API Ring
- Handler island
- Graph API island
- Graph Health API island
- Session/API bridge island
- Contract serialization island
- ResolutionGap island

Frontend Ring
- Function island
- Component/UI island
- Hook island
- Service island
- State island
- ResolutionGap island
```

Cach nay giup nguoi dung nhin graph la thay van de tap trung o Backend, Frontend, Shared hay Docs/Test, thay vi tat ca bi tron trong mot vong lon.

API Ring nen nam giua Backend Ring va Frontend Ring vi no la cau noi giua hai lop nay:

```text
Backend Ring <-> API Ring <-> Frontend Ring
```

Neu sau nay them Shared/Contract Ring, no co the nam gan API Ring vi contract la hinh dang du lieu cua API.

## 14. Quan he giua mau node va ring

Mau hien tai theo node type/filter van nen giu.

Khac biet la vi tri khong con chi theo node type toan cuc. Vi tri phai theo hai cap:

```text
macro position = App Layer ring
micro position = node type island inside ring
```

Vi du:

- Function backend nam trong Backend Ring, Function island.
- Function frontend nam trong Frontend Ring, Function island.
- ResolutionGap backend nam trong Backend Ring, ResolutionGap island.
- ResolutionGap frontend nam trong Frontend Ring, ResolutionGap island.

Dieu nay tranh viec tat ca Function nam chung mot cum lon, khien BE/FE bi tron.

## 15. UI filter/lens can co

Web UI phai the hien duoc cac lop moi, khong chi de AI doc metadata.

Lens/filter nen co:

- Backend unresolved calls
- API unresolved handlers/contracts
- Frontend unresolved type refs
- Shared contract analyzer gaps
- External unresolved symbols
- Builtin/Test/Stdlib non-actionable
- In-repo analyzer gaps
- Resolution gaps by functional area
- Top app layers by analyzer gap count
- Top functional areas by unresolved count
- Top unresolved target text

Filter mong muon:

```text
App Layer: Backend
Functional Area: Resolution
Resolution Gap: Unresolved Call Target
Actionability: Analyzer Gap
```

Hoac:

```text
App Layer: Frontend
Resolution Gap: Unresolved Type Target
Actionability: Review
```

Hoac:

```text
App Layer: API
Functional Area: Graph Health API
Resolution Gap: Unresolved Type Target
Actionability: Analyzer Gap
```

## 16. Khong nen bien thao luan nay thanh dead-code problem

Da co luc thao luan cham vao dead code, node co don, no_incoming, detached_component. Nhung trong tam cua tai lieu nay khong phai dead code.

Dead code la mot ket luan triage khac, co the dung topology va resolution confidence lam input.

Van de hien tai la:

```text
unresolved_reference chua duoc bieu dien thanh du lieu co nghia tren graph/UI.
```

Dead-code lens co the la buoc sau. Khong nen de no lam lech huong thiet ke ResolutionGap.

## 17. Query health / query accuracy audit

Anvien `query` hien co dau hieu tra ket qua nhieu. Query ve unresolved/resolution co the nhay sang launcher/web client, thay vi tim dung `resolve.go`, `emit.go`, `diagnostics.go`.

Can audit de biet:

- index dung nhung ranking yeu?
- process extraction chua bao phu luong resolution?
- query mechanism cu khong con phu hop voi codebase hien tai?
- query dang dua qua nhieu vao process labels cu?
- query co thieu app layer / functional area lam ranking signal khong?

Can benchmark nho:

```text
Query Intent Benchmark
- intent
- expected files/symbols
- actual top results
- hit@5 / hit@10
- noise reason
```

Vi du intent `unresolved reference diagnostic generation` phai tim dung:

- `internal/resolution/resolve.go`
- `internal/resolution/emit.go`
- `internal/graphhealth/diagnostics.go`

Neu khong tim dung, query chua dang tin cho repo lon.

## 18. Query benchmark examples

Mot bo benchmark toi thieu nen co cac intent:

```text
Intent: unresolved reference diagnostic generation
Expected:
- internal/resolution/resolve.go
- internal/resolution/emit.go
- internal/graphhealth/diagnostics.go

Intent: graph health unknown connectivity separation
Expected:
- internal/graphhealth/compute.go
- internal/graphhealth/policy.go
- anvien-web/src/lib/graph-health-filters.ts

Intent: graph clustering island layout
Expected:
- web graph layout code
- layout optimizer code
- graph health filters/lens code neu co lien quan

Intent: runtime reset hidden terminal window
Expected:
- anvien-launcher/src/main.go
- runtime start/reset/stop functions
```

Benchmark nay khong thay the test. No do kha nang `query` giup nguoi dung va AI tim dung code trong repo lon.

## 19. CLI / command can cai tien

`analyze` giu vai tro goc va la source of truth. Khong nen lam lech y nghia cua `analyze`.

Nhung cac lenh con can hieu cac lop moi:

- `query`
- `impact`
- `context`
- `detect-changes`

Huong nang cap:

- `query`: tra ket qua kem App Layer, Functional Area, ranking reason, noise warning.
- `impact`: bao blast radius theo BE/FE va functional area, khong chi bao HIGH/CRITICAL chung chung.
- `context`: hien node type, app layer, functional area, topology, resolution gaps, process/community membership.
- `detect-changes`: bao changed app layers, changed functional areas, resolution health impact.

Lenh moi co the can:

- `resolution-gaps`: thong ke unresolved theo layer, fact family, target role, actionability.
- `inventory` hoac `lens`: inventory theo app layer / functional area / topology / resolution health.
- `query-benchmark`: do query co con tim dung codebase hien tai khong.

## 20. Vi sao can nang cap command, khong chi UI

Neu Web UI co lens moi nhung CLI/query/context/impact khong hieu cac lop moi, Anvien se bi lech giua cac be mat:

- UI nhin theo App Layer, nhung query khong tim theo App Layer.
- context khong noi symbol thuoc BE/FE/functional area nao.
- impact chi bao risk chung, khong noi anh huong layer nao.
- detect-changes khong noi thay doi co lam Resolution Health xau hon khong.

Vi vay cac lenh con can cung dung chung semantic layer:

```text
App Layer
Functional Area
ResolutionGap
Topology Health
Resolution Health
```

## 21. Muoi van de da thao luan tiep va huong chot

Phan nay ghi lai 10 van de tiep theo da duoc thao luan sau khi bo sung App Layer, API Ring va ResolutionGap.

### 21.1. App Layer single-label hay multi-label

Van de: mot node co the vua la frontend vua la test, hoac vua la api vua la shared contract.

Huong chot:

- Khong dung multi-label chong cheo kieu `frontend` + `test` tren cung mot truc hien thi.
- Cai nao khong phai hon hop thi la mot loai rieng.
- Cai nao hon hop thi tao loai rieng cho hon hop do.
- Khong de cung mot mo ta bi lap lai o nhieu node/layer khac nhau.
- Neu mot node la ket hop cua 3 thu thi tao mot loai ket hop rieng cho 3 thu do.

Vi du:

```text
frontend
backend
api
shared_contract
frontend_test
api_contract
api_shared_contract
frontend_api_client
docs
config
```

Ly do: nhin graph se ro hon. Vung nao la hon hop thi thay la hon hop, khong can doc nhieu tag moi hieu.

### 21.2. ResolutionGap persist hay virtual

Van de: `ResolutionGap` co nen ghi vao `.anvien/graph.json` nhu node/edge that, hay chi sinh virtual trong API/UI lens?

Huong chot:

- `ResolutionGap` nen persist vao graph.
- Khong chi la virtual trong UI/API.

Ly do:

- query/context/impact/detect-changes can cung nhin thay no;
- CLI va UI phai dung chung mot source of truth;
- graph nen phan anh dung nhung gi analyzer biet va khong biet;
- neu chi virtual o UI thi AI/CLI khong khai thac duoc day du.

### 21.3. API Ring, contract va frontend API client nam o dau

Van de: API co nhieu phan:

- API server handler;
- API contract/schema/generated types;
- frontend API client.

Huong chot:

- Khong ep vao mot ring duy nhat neu chung co vai tro khac nhau.
- Co bao nhieu loai ro rang thi tach thanh bay nhieu loai/ring phu hop.
- Khong bi gioi han so vong tron.
- Neu them ring giup tach ro van de thi cu them ring.

Vi du:

```text
API Ring
API Contract Ring
Frontend API Client group/ring
Shared Contract Ring
```

Nguyen tac: cai gi cang ro rang, khong xam lan, cang de kiem soat.

### 21.4. Test/Docs/Config la ring rieng hay modifier

Van de: test/docs/config co nen la ring rieng hay chi la modifier?

Huong chot:

- Khong so nhieu ring.
- Ring co the lon hoac nho.
- Neu tach ring giup ro van de thi nen tach.
- Khong can co dinh vao it ring.

Vi du:

```text
Frontend Test Ring
Backend Test Ring
Docs Ring
Config Ring
Generated Contract Ring
```

Neu mot loai la hon hop thi tao loai hon hop rieng, khong dung tag chong cheo lam mo nghia.

### 21.5. Functional Area suy ra bang rule nao

Van de: App Layer co the suy ra bang path kha de, nhung Functional Area kho hon.

Huong chot:

- Chon nguon nao cho ket qua chinh xac nhat.
- Khong chon huong co do chinh xac thap chi vi de lam.
- Neu can ket hop nhieu nguon de dat do chinh xac cao hon thi can thiet ke theo huong do.

Nguon co the can danh gia:

- path prefix;
- package/module name;
- process membership;
- community detection;
- import/call neighborhood;
- explicit config file;
- AI-assisted labeling sau analyze neu co bang chung va co the verify.

Tieu chi chon: chinh xac truoc, tien loi sau.

### 21.6. ResolutionGap edge nen tach nhu the nao

Van de: nen dung edge chung `HAS_RESOLUTION_GAP` hay tach nhieu loai edge nhu `UNRESOLVED_CALLS`, `UNRESOLVED_ACCESSES`, `UNRESOLVED_USES_TYPE`.

Huong chot:

- Tach cang nho cang de kiem soat.
- Khong nhat thiet chi co 3 loai.
- Neu fact family/target role/actionability can edge rieng de nhin ro thi nen tach.

Vi du co the co:

```text
UNRESOLVED_CALLS
UNRESOLVED_ACCESSES
UNRESOLVED_TYPE_REFERENCE
UNRESOLVED_HERITAGE
UNRESOLVED_EXTERNAL_SYMBOL
UNRESOLVED_BUILTIN_REFERENCE
UNRESOLVED_TEST_REFERENCE
```

Can tranh gop qua rong lam mat nghia.

### 21.7. Query benchmark nen la test hay report

Van de: query benchmark neu chi la report thu cong thi de quen.

Huong chot:

- Nen co lenh CLI moi co chuc nang nay.
- Khi can danh gia thi chi can chay lenh do de ra ket qua.

Vi du ten lenh can thao luan:

```text
anvien query-benchmark
anvien query-health
anvien benchmark-query
```

Lenh nay nen output duoc:

- intent;
- expected files/symbols;
- actual top results;
- hit@5 / hit@10;
- noise reason;
- pass/fail.

### 21.8. Schema/version compatibility va graph cu

Van de da neu: neu graph cu chua co App Layer/ResolutionGap thi UI/API/CLI xu ly sao?

Dieu chinh sau thao luan:

- Day khong phai van de can fallback doan mo.
- Tool da co rule ro: lam viec graph-based thi phai `anvien analyze --force` truoc.
- `analyze` la source of truth.

Huong chot:

```text
Run analyze first. Analyze is the source of truth. No stale graph fallback.
```

Quy tac:

- Khi them App Layer / ResolutionGap / API Ring, `analyze` phai sinh schema moi day du.
- UI/API/CLI lam viec tren graph da analyze theo schema hien tai.
- Neu graph thieu metadata moi thi coi la stale/incomplete graph.
- Khong fallback classify o load time.
- Khong show tam `unknown` de che van de.
- Khong co gang support graph cu bang heuristic.

### 21.9. Performance va graph size

Van de: neu 51k unresolved occurrences thanh node/edge that thi graph se lon hon.

Huong chot:

- Trong tam khong phai graph lon hay nho.
- Tool viet bang Go va bai toan chiu tai lon da la muc tieu cua tool.
- Cai can uu tien la su chinh xac.
- Khong hy sinh do chinh xac chi de lam graph nho hon.

Dieu can thiet ke:

- dung data model chinh xac;
- neu can aggregate/dedupe thi chi lam khi khong lam mat nghia;
- khong cap/cat bot evidence neu lam sai ban chat van de.

### 21.10. User-facing naming

Van de: ten hien thi nhu `Resolution Gap`, `Unresolved Symbol`, `Analyzer Gap`, `External Reference` can chot.

Huong chot:

- Chon ten phu hop voi van de that.
- Khong bi khoa vao danh sach ten ban dau.
- Neu van de mo rong thi co the them ten moi.
- Ten phai giup nguoi dung nhin graph hieu ngay y nghia.

Ten hien tai co the can tiep tuc thao luan:

```text
Resolution Gap
Unresolved Symbol
Analyzer Gap
External Reference
Non-actionable Reference
App Layer
API Layer
API Contract
Frontend API Client
```

## 22. Cac cau hoi con mo truoc khi tao plan

Sau 10 van de tren, cac cau hoi con can lam ro khi tao plan:

1. Danh sach App Layer/ring ban dau se gom chinh xac nhung loai nao?
2. Cach dat ten cho cac loai hon hop nhu `frontend_test`, `api_contract`, `frontend_api_client`.
3. ResolutionGap persist thanh node rieng, edge rieng, hay ca node va edge rieng theo tung fact family.
4. Functional Area dung rule nao de dat do chinh xac cao nhat.
5. Query benchmark command nen doc suite tu dau va output schema nao.
6. Web UI se hien ring nho/to nhu the nao de khong lam roi mat khi co nhieu ring.
7. Loai nao mac dinh hien, loai nao mac dinh an.
8. Tieu chi nao de khang dinh App Layer/Functional Area classification la dung.

## 23. Rui ro thiet ke

Mot so rui ro can tranh:

- Gan sai BE/FE lam graph dep gia tao nhung sai nghia.
- Tao qua nhieu ResolutionGap node lam graph no kich thuoc.
- Tron ResolutionGap vao topology lam quay lai loi cu.
- Coi unresolved la dead code khi chua co bang chung.
- De UI chi co metadata nhung khong co filter/lens de nguoi dung dung duoc.
- Cai tien query dua tren mot vai query mau ma khong co benchmark.
- Them rule phan loai qua dac thu cho Anvien ma khong dung cho repo khac.

## 24. Huong ket luan hien tai

Huong dung khong phai chi tach diagnostic.

Can nang graph thanh nhieu lop ngu nghia:

```text
Node Type
+ App Layer BE/API/FE/Shared/Test/Docs
+ Functional Area
+ ResolutionGap / UnresolvedSymbol
+ Multi-ring layout theo App Layer
+ UI filters/lens
+ CLI commands hieu cac lop nay
+ Query accuracy audit
```

Buoc dau hop ly nhat:

1. Them App Layer cho node.
2. Dung App Layer de chia layout thanh Backend Ring, API Ring va Frontend Ring.
3. Dua `ResolutionGap` / `UnresolvedSymbol` thanh entity/filter rieng.
4. Bo sung UI lens/filter cho Resolution Health.
5. Audit va nang cap `query` de phu hop voi graph/codebase hien tai.

## 25. Dieu da dong thuan trong thao luan

Cac diem da dong thuan:

- `unknown_connectivity` va `unresolved_reference` khong duoc tron lai.
- `unresolved_reference` la van de rieng, lon hon mot diagnostic text thong thuong.
- UI phai the hien duoc van de, khong chi AI doc duoc metadata.
- BE/API/FE la truc phan loai can co, va nen co truoc functional area sau.
- API nen la App Layer/ring rieng vi no la ranh gioi quan trong giua FE va BE.
- Layout nen chuyen sang nhieu vong lon theo App Layer, toi thieu Backend Ring, API Ring va Frontend Ring.
- Trong moi ring van giu cach chia dao theo node type/filter mau hien co.
- `analyze` la lenh goc, cac lenh con nhu `query`, `impact`, `context`, `detect-changes` moi la noi can nang cap.
- `query` can audit bang benchmark vi co dau hieu tra nhiu voi codebase hien tai.
- App Layer khong nen la multi-label chong cheo; loai hon hop nen thanh loai rieng.
- ResolutionGap nen persist vao graph.
- Khong so nhieu ring neu nhieu ring lam graph ro nghia hon.
- Su chinh xac quan trong hon viec graph lon hay nho.
- Query benchmark nen thanh mot lenh CLI rieng.
- Khong fallback graph cu bang doan mo; graph-based workflow phai analyze truoc.
