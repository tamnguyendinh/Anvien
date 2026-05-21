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

Vi du cho repo AVmatrix-GO:

```text
avmatrix-web/src/**
=> frontend

avmatrix-web/test/**
avmatrix-web/e2e/**
=> frontend + test

internal/**
cmd/**
=> backend

internal/httpapi/**
=> api

avmatrix-web/src/services/backend-client.ts
=> frontend + api_client

avmatrix-launcher/**
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

AVmatrix `query` hien co dau hieu tra ket qua nhieu. Query ve unresolved/resolution co the nhay sang launcher/web client, thay vi tim dung `resolve.go`, `emit.go`, `diagnostics.go`.

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
- avmatrix-web/src/lib/graph-health-filters.ts

Intent: graph clustering island layout
Expected:
- web graph layout code
- layout optimizer code
- graph health filters/lens code neu co lien quan

Intent: runtime reset hidden terminal window
Expected:
- avmatrix-launcher/src/main.go
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

Neu Web UI co lens moi nhung CLI/query/context/impact khong hieu cac lop moi, AVmatrix se bi lech giua cac be mat:

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

## 21. Cac cau hoi can chot truoc khi tao plan

Nhung cau hoi can chot:

1. App Layer ban dau co chi gom BE/FE hay gom luon shared/test/docs/config?
2. Test nen la app layer rieng hay la modifier tren BE/FE?
3. `internal/contracts` nen la backend hay shared_contract?
4. `avmatrix-web/e2e` nen la frontend hay test?
5. `internal/httpapi` nen la `api` App Layer rieng hay backend functional area? Huong thao luan hien tai nghieng ve `api` App Layer rieng.
6. API client trong frontend nen gan `frontend + api_client` hay gan vao API Ring?
7. API contract/generated contract nen nam trong `api`, `shared_contract`, hay `api_contract` modifier?
8. ResolutionGap nen la node graph that, virtual node trong API response, hay filter/lens tinh o UI?
9. Co can persist ResolutionGap vao `.avmatrix/graph.json` hay tinh luc load API?
10. Query benchmark nen dat trong docs, reports, hay internal testdata?
11. App Layer ring co can hien docs/test/config mac dinh hay chi hien khi bat filter?

## 22. Rui ro thiet ke

Mot so rui ro can tranh:

- Gan sai BE/FE lam graph dep gia tao nhung sai nghia.
- Tao qua nhieu ResolutionGap node lam graph no kich thuoc.
- Tron ResolutionGap vao topology lam quay lai loi cu.
- Coi unresolved la dead code khi chua co bang chung.
- De UI chi co metadata nhung khong co filter/lens de nguoi dung dung duoc.
- Cai tien query dua tren mot vai query mau ma khong co benchmark.
- Them rule phan loai qua dac thu cho AVmatrix-GO ma khong dung cho repo khac.

## 23. Huong ket luan hien tai

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

## 24. Dieu da dong thuan trong thao luan

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
