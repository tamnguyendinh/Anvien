# Tong hop thao luan: App Layer, Resolution Gap, UI lens va query health

Date: 2026-05-22
Scope: discussion only, not an implementation plan

## 1. Van de goc

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

Cach phan loai nay cho biet "node la loai symbol gi", nhung chua cho biet node thuoc phan nao cua san pham.

Vi vay khi gap `unresolved_reference`, UI chi biet co reference chua resolve, nhung chua tra loi duoc:

- no thuoc Backend hay Frontend?
- no nam trong nhom chuc nang nao?
- target chua resolve co ve la callable, member, type, external symbol, builtin hay test helper?
- no co lam topology mat do tin cay hay chi la diagnostic khong quan trong?

## 2. Nhan dinh da chot

`unknown_connectivity` da duoc tach dung khoi `unresolved_reference`.

Sau khi tach:

- `unknown_connectivity` chi con la topology that su khong xac dinh duoc.
- node co unresolved diagnostic nhung co edge ket noi van giu topology that, vi du `connected`.
- `unresolved_reference` chuyen ve dung vai tro diagnostic/resolution gap.

Tuy nhien, `unresolved_reference` van la van de lon vi no bao hieu graph co the dang thieu edge. Neu thieu edge thi cac trang thai topology nhu `no_incoming`, `true_isolated`, `detached_component` co the bi sai hoac thieu tin cay.

## 3. App Layer / Runtime Surface

Can them mot lop phan loai rieng, khong thay the node type hien co.

De xuat app layer:

- `backend`
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
- de lam nen cho query/impact/context/detect-changes sau nay.

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

## 4. Functional Area

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

Functional Area khong nen thay the App Layer. Thu tu dung la:

```text
App Layer first, Functional Area second, Node Type third.
```

## 5. Resolution Gap Entity

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

## 6. Layout UI moi

Hien tai graph co mot vong tron lon chua cac quan dao node theo mau/type. Huong moi nen chuyen sang nhieu macro-ring theo App Layer.

Toi thieu:

```text
Backend Ring
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

Frontend Ring
- Function island
- Component/UI island
- Hook island
- Service island
- State island
- ResolutionGap island
```

Cach nay giup nguoi dung nhin graph la thay van de tap trung o Backend, Frontend, Shared hay Docs/Test, thay vi tat ca bi tron trong mot vong lon.

## 7. UI filter/lens can co

Web UI phai the hien duoc cac lop moi, khong chi de AI doc metadata.

Lens/filter nen co:

- Backend unresolved calls
- Frontend unresolved type refs
- Shared contract analyzer gaps
- External unresolved symbols
- Builtin/Test/Stdlib non-actionable
- In-repo analyzer gaps
- Resolution gaps by functional area

Filter mong muon:

```text
App Layer: Backend
Functional Area: Resolution
Resolution Gap: Unresolved Call Target
Actionability: Analyzer Gap
```

## 8. Query health / query accuracy audit

AVmatrix `query` hien co dau hieu tra ket qua nhieu. Query ve unresolved/resolution co the nhay sang launcher/web client, thay vi tim dung `resolve.go`, `emit.go`, `diagnostics.go`.

Can audit de biet:

- index dung nhung ranking yeu?
- process extraction chua bao phu luong resolution?
- query mechanism cu khong con phu hop voi codebase hien tai?

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

## 9. CLI / command can cai tien

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

## 10. Ket luan

Huong dung khong phai chi tach diagnostic.

Can nang graph thanh nhieu lop ngu nghia:

```text
Node Type
+ App Layer BE/FE/Shared/Test/Docs
+ Functional Area
+ ResolutionGap / UnresolvedSymbol
+ Multi-ring layout theo App Layer
+ UI filters/lens
+ CLI commands hieu cac lop nay
+ Query accuracy audit
```

Buoc dau hop ly nhat:

1. Them App Layer cho node.
2. Dung App Layer de chia layout thanh Backend Ring va Frontend Ring.
3. Dua `ResolutionGap` / `UnresolvedSymbol` thanh entity/filter rieng.
4. Bo sung UI lens/filter cho Resolution Health.
5. Audit va nang cap `query` de phu hop voi graph/codebase hien tai.
