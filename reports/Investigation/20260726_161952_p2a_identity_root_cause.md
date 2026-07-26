# P2-A — TypeScript extraction and graph identity root-cause trace

Status: **bounded investigation complete; not acceptance; no remediation**

## Scope and authority

This slice restarted the TypeScript extraction/identity check independently against the real target repository. Earlier reports were not used as evidence.

| Item | Value |
|---|---|
| Target | `E:\cheapapp.org` |
| Target HEAD | `a869876ab6262dacde6cd5d432d099a91852a646` |
| Target graph | `E:\cheapapp.org\.anvien\graph.json` |
| Graph SHA-256 | `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090` |
| Indexed commit | `a869876ab6262dacde6cd5d432d099a91852a646` |
| Indexed at | `2026-07-26T07:49:43Z` |
| Graph inventory | 1,359 files; 84,807 nodes; 114,125 relationships |
| Anvien source HEAD | `107e0157ec072e54b44246719da3e7accf76e1cb` |

The three target-source hashes were:

| File | SHA-256 |
|---|---|
| `modules/email/server/operations/email-operations-observability.ts` | `A9247972D44D0A74C411FB7CF83D681FEBA371C7A85EEB83F8986AFD6446E40C` |
| `modules/release-distribution/server/release-distribution-publication-state.ts` | `B53F966F82AEDAC58EC15A892B9B71BA0A5716638703105BB35B126CB398D474` |
| `modules/commercial-config/server/admin-commercial-config/read-admin-commercial-config.ts` | `44467B1D6F1778AE4728B4941136321F7831D8FD67EBE2C93A430A3D132F90AC` |

## Result matrix

| ID | Source truth | Graph observation | First divergence | Classification |
|---|---|---|---|---|
| P2A-1 | The array binding pattern at email lines 503-509 declares six local bindings. | No graph symbol exists for any of the six names; all six later `.map` calls are actionable `ResolutionGap` nodes. | `internal/providers/tsjs/definitions.go:64-67` rejects any `variable_declarator` whose `name` is not an `identifier`. | **CONFIRMED WRONG** |
| P2A-2 | `time` is declared independently at lines 207 and 214; `now` is declared independently at lines 262 and 501. Anvien ScopeIR retains four distinct definition IDs and scope bindings. | The graph contains one `Variable ...:time` node at line 207 and one `Variable ...:now` node at line 262. | `internal/resolution/indexes.go:814-823` drops range and scope from the graph ID; `internal/graph/types.go:96-100` replaces an existing node on duplicate ID. | **CONFIRMED WRONG** |
| P2A-3 | The three fixed files contain 21 top-level TypeScript exports: 17 + 3 + 1. | All 21 declarations are represented, but none is marked exported; each `file-detail` summary reports `exportedSymbolCount=0`. | TSJS `addDefinition` builds `DefinitionFact` without setting its available `Visibility` field at `internal/providers/tsjs/definitions.go:100-111`. | **CONFIRMED WRONG** |
| P2A-4 | `readAdminCommercialConfig` is one unique declaration at lines 10-23. | One correctly ranged Function node with the expected file-qualified UID exists. | No identity divergence for this unique-name control. Its export bit is still lost by P2A-3. | **BOUNDED VALID control** |

## P2A-1 — Six array-binding declarations are lost during extraction

Source at `email-operations-observability.ts:503-509` declares:

- `messageRows` at line 504;
- `attemptRows` at line 505;
- `eventRows` at line 506;
- `providerEventRows` at line 507;
- `readinessRows` at line 508;
- `suppressionRows` at line 509.

The TypeScript compiler AST classifies the parent name as `ArrayBindingPattern` and exposes all six identifiers as binding elements. The independent AST-to-graph probe found one source declaration for each and zero graph Variable nodes for each.

The exact Anvien ScopeIR probe independently reproduced the loss before graph emission. Its `missingBindingNames` contains all six names. This is consistent with the TSJS provider implementation:

- `internal/providers/tsjs/definitions.go:64-65` reads `variable_declarator.name`;
- `:66-67` returns unless that node kind is exactly `identifier`;
- therefore an `array_pattern` never reaches `addDefinition`, never receives a `DefinitionFact`, and never receives a scope `BindingFact`.

The downstream consequence is directly visible in the fresh graph. Anvien Cypher returned no symbol rows for the six names, while six source-backed gaps exist:

| Gap target | Source range | Classification |
|---|---|---|
| `attemptRows.map` | 586-591 | `in_repo_unresolved` / `analyzer_gap` |
| `eventRows.map` | 592-596 | `in_repo_unresolved` / `analyzer_gap` |
| `messageRows.map` | 597-608 | `in_repo_unresolved` / `analyzer_gap` |
| `providerEventRows.map` | 609-614 | `in_repo_unresolved` / `analyzer_gap` |
| `readinessRows.map` | 615-621 | `in_repo_unresolved` / `analyzer_gap` |
| `suppressionRows.map` | 622-626 | `in_repo_unresolved` / `analyzer_gap` |

The first loss is extraction, not graph persistence or a query projection.

## P2A-2 — Same-file locals remain distinct in ScopeIR but collide in the graph

The independent TypeScript AST inventory found:

| Name | Source declaration 1 | Source declaration 2 |
|---|---|---|
| `time` | `parseTime`, line 207 | reducer callback inside `latestIso`, line 214 |
| `now` | `buildEmailOperationsReport`, line 262 | `readEmailOperationsReport`, line 501 |

The TSJS provider retains distinct range-bearing identities and scope bindings:

| Source binding | ScopeIR definition ID | Scope |
|---|---|---|
| `time:207` | `def:...#207:8:Variable:time` | function scope `202:0-210:1` |
| `time:214` | `def:...#214:10:Variable:time` | callback function scope `213:46-221:3` |
| `now:262` | `def:...#262:8:Variable:now` | function scope `254:7-417:1` |
| `now:501` | `def:...#501:8:Variable:now` | function scope `497:7-634:1` |

This proves extraction and scope ownership remain distinct for these four local declarations. The first divergence occurs later:

1. `internal/resolution/indexes.go:814-823` computes a graph ID from label, cleaned file path, qualified name/name, and optional arity. It does not include `DefinitionFact.ID`, range, or scope.
2. Both `time` definitions therefore map to `Variable:modules/email/server/operations/email-operations-observability.ts:time`.
3. Both `now` definitions map to `Variable:modules/email/server/operations/email-operations-observability.ts:now`.
4. `internal/graph/types.go:96-100` replaces the existing node when the ID already exists.

The bounded count closes exactly for this file: the fresh ScopeIR probe emits 169 definitions; the graph contains 167 symbol nodes; the two duplicate-ID pairs reduce four definition facts to two graph nodes. The graph retains only `time:207` and `now:262`.

This is an identity-injectivity failure at graph-ID construction/emission, not a Tree-sitter parse failure.

## P2A-3 — TypeScript export visibility is not propagated

The independent TypeScript compiler AST found the following exported declaration lines:

- email file: `35, 48, 55, 61, 68, 76, 82, 91, 99, 101, 109, 118, 150, 158, 254, 438, 497` — 17 exports;
- release publication-state file: `1, 7, 13` — 3 exports;
- admin commercial-config read file: `10` — 1 export.

All 21 source exports have corresponding graph declaration nodes at the correct file/start-line control. None has `exported=true`, `isExported=true`, or `visibility` equal to `public`/`exported`. `file-detail` independently reports:

| File | Graph symbols | Source exports represented | Graph rows flagged exported | `exportedSymbolCount` |
|---|---:|---:|---:|---:|
| email observability | 167 | 17/17 | 0 | 0 |
| release publication state | 3 | 3/3 | 0 | 0 |
| admin commercial-config read | 1 | 1/1 | 0 | 0 |

The IR contract can carry this semantic: `internal/scopeir/facts.go:3-28` defines `DefinitionFact`, including `Visibility` at line 17. The TSJS collector does not populate it:

- `internal/providers/tsjs/definitions.go:100-111` constructs every `DefinitionFact` without `Visibility`;
- the exact ScopeIR probe reports empty visibility for every definition in all three files;
- `internal/resolution/emit.go:158-173` can only emit a non-empty `visibility` property if the provider supplied one;
- `internal/filecontext/context.go:1334-1340` consequently classifies every affected symbol as unexported.

The first divergence is the TSJS AST-to-IR path. The correct declaration nodes and ranges show this is not a file-discovery or declaration-existence failure.

## Bounded control and adjacent observation

`readAdminCommercialConfig` is a valid unique-name control for declaration identity:

- source: one Function declaration at lines 10-23;
- ScopeIR: one `def:...#10:7:Function:readAdminCommercialConfig`;
- graph: one `Function:...:readAdminCommercialConfig` at lines 10-23.

The same file also contains a graph gap for ambient `Promise` at line 13. The independent TypeScript checker resolves it to TypeScript library declarations with no diagnostic, while Anvien labels it `in_repo_unresolved/analyzer_gap`. That behavior is **confirmed inconsistent at this source site but is not root-caused in P2-A**, because ambient/compiler-library resolution is outside this extraction/identity slice.

## Blast-radius warning

No production code was changed. The pre-edit evidence required by repository rules was nevertheless collected for the causal code paths:

- `internal/providers/tsjs/definitions.go`, `internal/resolution/indexes.go`, `internal/resolution/emit.go`, `internal/graph/types.go`, and `internal/filecontext/context.go` are file-level `high` risk in fresh Anvien `file-detail` output;
- `graphIDForDef`, `Graph.AddNode`, `emitDefinitionNodes`, and `exportedSymbol` each report **CRITICAL** upstream impact;
- CRITICAL is a scope warning, not a prohibition. Any later remediation must be a separately planned slice with focused impact review.

## Analyze side-effect observation

The target graph was refreshed earlier in this worker with the supported direct-target command and was not rerun after the P2-A handoff. The resulting graph hash is the fixed hash stated above.

The fresh analyze completed at approximately `2026-07-26 14:49:43 +07` and wrote the target-local `.anvien` graph as expected. It also gave `AGENTS.md` and `CLAUDE.md` last-write timestamps of `14:49:44 +07`:

| Path | Current SHA-256 | Current last-write time |
|---|---|---|
| `AGENTS.md` | `B777E8D2E38E1405EFB1305F99287044CE1D2E078D679EEDCE45AD9CEF9EA2A9` | `2026-07-26T14:49:44.1765411+07:00` |
| `CLAUDE.md` | `1F5A3B5974BDCDB42FC942B9229746F1B37C19B68188563079B734B7A5BD46C2` | `2026-07-26T14:49:44.1898525+07:00` |

Both files and `.anvien/` are ignored by target `.gitignore` and are not tracked. A pre-analyze content-hash/timestamp baseline for these generated files was not captured in this fresh run. Therefore this report **cannot prove whether their content changed or whether analyze only rewrote/touched deterministic content**. It makes neither claim and does not revert them. The three assigned target source files currently have zero tracked diff and retain the hashes recorded above.

## Commands executed

Fresh target graph creation, before the P2-A handoff:

```powershell
$env:GIT_CONFIG_GLOBAL='E:\Anvien\.tmp\p1b_gitconfig'
anvien analyze E:\cheapapp.org --force --json `
  --benchmark-json E:\Anvien\reports\Investigation\p1b_ts_identity_analyze_benchmark.json `
  --benchmark-label p1b_ts_identity `
  --name cheapapp-accuracy-direct
```

Graph inspection:

```powershell
anvien file-detail <fixed-file> --repo E:\cheapapp.org --json --relationships 20 --unresolved 20 --linked 5
anvien cypher "MATCH (n:Variable) WHERE n.filePath = 'modules/email/server/operations/email-operations-observability.ts' AND n.name IN ['time','now','messageRows','attemptRows','eventRows','providerEventRows','readinessRows','suppressionRows'] RETURN n.id, n.name, n.startLine, n.endLine ORDER BY n.name" --repo E:\cheapapp.org
anvien cypher "MATCH (n) WHERE n.filePath = 'modules/email/server/operations/email-operations-observability.ts' AND n.name IN ['messageRows','attemptRows','eventRows','providerEventRows','readinessRows','suppressionRows'] RETURN n.id, label(n), n.name, n.startLine, n.endLine ORDER BY n.name" --repo E:\cheapapp.org
```

Independent probes:

```powershell
node .\.tmp\cheapapp-graph-root-cause-restart\p2a-identity\p2a_identity_probe.mjs
go run .\.tmp\cheapapp-graph-root-cause-restart\p2a-identity\scopeir_probe
powershell -NoProfile -File .\.tmp\cheapapp-graph-root-cause-restart\p2a-identity\p2a_target_state_probe.ps1
```

Anvien self-graph evidence was refreshed with `anvien analyze --force --json` from `E:\Anvien` before its `file-detail` and impact commands.

## Raw evidence

- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2a-identity\p2a_identity_source_graph_diff.json` — TypeScript AST versus raw graph facts.
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2a-identity\p2a_scopeir_probe.json` — exact TSJS ScopeIR definitions, bindings, missing binding names, visibility, and graph-ID collision groups.
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2a-identity\p2a_cli_evidence.json` — compact `file-detail` and Cypher results from the fresh target graph.
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2a-identity\p2a_anvien_root_cause_trace.json` — hashed Anvien source excerpts for the first-divergence paths.
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2a-identity\p2a_anvien_impact_evidence.json` — fresh Anvien file-detail and impact summaries.
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2a-identity\p2a_target_state_after.json` — target HEAD, hashes, current status, ignore rules, and the side-effect evidence limitation.
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2a-identity\p2a_fresh_analyze_benchmark.json` — copied analyze benchmark output; total duration 50.030129 seconds.

## Limitations

- Findings are proven only for the exact three files and named sites in this slice; repository-wide prevalence was not measured.
- The TypeScript AST probe used TypeScript 5.9.3 from `E:\Anvien\node_modules`; the checked syntax and binding forms are stable for these cases, but this is not a target-wide TypeScript 6.0.3 semantic audit.
- No target source was edited or copied. No report or probe was written into `E:\cheapapp.org`.
- No production Anvien code, tests, or fixtures were changed. No build, remediation validation, or acceptance review was performed.
- The adjacent `Promise` gap, member/property accuracy, import/barrel semantics, database projection, command projection, and global prevalence remain separate slices.
