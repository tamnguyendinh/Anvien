# Child 02 P2-C Playwright evidence

- Generated: 2026-08-13T12:29:04.846Z
- Runtime: built backend http://127.0.0.1:4848 + built frontend preview http://127.0.0.1:5228
- Browser: Chromium headed/visible
- Controlled boundary: External HTTP graph/session/file responses only; mounted production AppStateProvider, ChatRuntimeContext, resolveNodeIds, handleNodeGroundingReference, and CodeReferencesPanel are exercised unchanged.
- Verdict: PASS

## Reader matrix

| Row | Verdict | Evidence |
| --- | --- | --- |
| C09 | PASS | {"verdict":"PASS","input":{"exact":"opaque-exact-alpha","suffix":"exact-alpha","nearMatchImpact":"suffix-alpha"},"expected":"Only the complete opaque graph ID is cyan-highlighted; suffix/near-match IDs select nothing and no node is red.","observed":"Mounted production tool_result path completed; graph screenshot shows one active cyan node, all non-exact nodes dim, and no red blast-radius node.","screenshot":"Reports/qa/playwright/child02-p2c/screenshots/03-c09-exact-id-only.png"} |
| C10 | PASS | {"verdict":"PASS","ambiguous":{"label":"Function","name":"sharedName","matchingOpaqueIds":["opaque-duplicate-a","opaque-duplicate-b"],"citationCount":0,"codePanelCount":0},"unique":{"opaqueId":"opaque-exact-alpha","label":"Function","name":"uniqueTarget","filePath":"src/unique.ts","range":{"startLine":10,"startCol":0,"endLine":12,"endCol":0},"visibleCitation":"L10–11"}} |
| C11 | PASS | {"verdict":"PASS","graphRange":{"startLine":10,"startCol":0,"endLine":12,"endCol":0},"fileAPIRequest":{"path":"src/unique.ts","repo":"qa-p2c","startLine":0,"endLine":60},"visibleCitation":"L10–11","lineStyles":{"line10":{"backgroundColor":"rgba(6, 182, 212, 0.14)","borderLeftColor":"rgb(154, 126, 99)","rect":{"x":248,"y":339.5,"width":547.2000122070312,"height":19.5,"top":339.5,"right":795.2000122070312,"bottom":359,"left":248}},"line11":{"backgroundColor":"rgba(6, 182, 212, 0.14)","borderLeftColor":"rgb(154, 126, 99)","rect":{"x":248,"y":359,"width":547.2000122070312,"height":19.5,"top":359,"right":795.2000122070312,"bottom":378.5,"left":248}},"line12":{"backgroundColor":"rgba(0, 0, 0, 0)","borderLeftColor":"rgba(0, 0, 0, 0)","rect":{"x":248,"y":378.5,"width":547.2000122070312,"height":19.5,"top":378.5,"right":795.2000122070312,"bottom":398,"left":248}}},"scrollState":{"scrollTop":0,"clientHeight":297,"scrollHeight":609},"scrollIntoViewCalls":[{"lineNumber":null,"text":"FunctionuniqueTargetsrc/unique.ts • L10–11Code not available in memory for src/unique.ts"},{"lineNumber":"10","text":"10"}],"expected":"One-based graph coordinates are converted once at the zero-based inclusive file API boundary; line 10 is targeted; lines 10-11 are highlighted; line 12 is excluded."} |

## Console and page errors

- Console warnings/errors: 0
- Page errors: 0

## Screenshots

- `Reports/qa/playwright/child02-p2c/screenshots/01-mounted-production-fixture.png`
- `Reports/qa/playwright/child02-p2c/screenshots/02-before-c09-tool-result.png`
- `Reports/qa/playwright/child02-p2c/screenshots/03-c09-exact-id-only.png`
- `Reports/qa/playwright/child02-p2c/screenshots/04-c10-ambiguous-fails-closed.png`
- `Reports/qa/playwright/child02-p2c/screenshots/05-c10-unique-persisted-reference.png`
- `Reports/qa/playwright/child02-p2c/screenshots/06-c11-lines-10-11-highlighted-line-12-excluded.png`

## Network evidence

```json
{
  "backendRequests": [
    {
      "method": "GET",
      "url": "http://127.0.0.1:4848/api/repo?repo=qa-p2c&awaitAnalysis=true"
    },
    {
      "method": "GET",
      "url": "http://127.0.0.1:4848/api/graph?repo=E%3A%2FAnvien&stream=true"
    },
    {
      "method": "GET",
      "url": "http://127.0.0.1:4848/api/repos"
    },
    {
      "method": "GET",
      "url": "http://127.0.0.1:4848/api/heartbeat"
    },
    {
      "method": "GET",
      "url": "http://127.0.0.1:4848/api/session/status?repoName=E%3A%2FAnvien"
    },
    {
      "method": "POST",
      "url": "http://127.0.0.1:4848/api/session/chat"
    },
    {
      "method": "POST",
      "url": "http://127.0.0.1:4848/api/session/chat"
    },
    {
      "method": "POST",
      "url": "http://127.0.0.1:4848/api/session/chat"
    },
    {
      "method": "GET",
      "url": "http://127.0.0.1:4848/api/file?path=src%2Funique.ts&repo=qa-p2c&startLine=0&endLine=60"
    }
  ],
  "fileRequest": {
    "path": "src/unique.ts",
    "repo": "qa-p2c",
    "startLine": 0,
    "endLine": 60
  },
  "failedResponses": [],
  "requestFailures": []
}
```
