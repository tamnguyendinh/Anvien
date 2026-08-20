# E3-P3C-REVIEW1 — independent Supervisor review

- Review ID: `E3-P3C-REVIEW1`
- Slice: Child 03 P3-C, “Project and resolve binding occurrences in the current graph”
- Repository: `E:\Anvien`
- Branch / reviewed HEAD: `master` / `33e73ec6319fd764d415cac8024b36f0ac75d70e`
- Assignment HEAD: `69089d17b6a248b2a7c03f45bfa0cfec3c78d7e4` (verified ancestor, exit 0)
- Review role: independent Supervisor; review and verdict only
- Verdict: **REJECT**
- Rejected invariants: **2 — Exact ownership**; **8 — Gap integrity**, as the direct downstream consequence of the same accepted owner drift
- Next responsible person: Orchestration main, which must route only the exact owner-drift repair boundary to the Coder lane

## Executive finding

The well-formed P3-C path is conserved end to end: the final exact bytes build, all focused and affected-package tests pass, seven binding occurrences become seven distinct `Variable` nodes and seven `DEFINES` endpoints, five reads plus one write resolve lexically, Graph JSON and Ladybug agree row for row, imports do not drift, and the well-formed boundary has no gaps, duplicates, or orphan endpoints.

The candidate nevertheless accepts malformed coordinated lexical-owner drift. An independent probe kept a binding leaf and its `DefinitionFact` in `src/a.ts`, moved both its `OwnedDefID` and matching local `Binding` together to a module scope in `src/b.ts`, and passed the two IRs to `resolution.ResolveInto` with a one-node sentinel graph. The candidate returned no error, mutated the graph from 1 to 4 nodes with 1 relationship, and recorded 1 unresolved reference:

```json
{"acceptedOwnerDrift":true,"nodesBefore":1,"nodesAfter":4,"relationshipsAfter":1,"unresolvedReferences":1}
```

That violates the required rule that file/range/selection/name/defID/scope ownership drift fail before graph mutation. It also creates the prohibited binding-caused unresolved/gap outcome. The defect is confined to the new binding-occurrence validation in `internal/resolution/resolve.go`.

## Authority and review boundary

The following authority was read completely before candidate review:

- `AGENTS.md`
- `.agents/skills/working-rules/SKILL.md`
- `.agents/skills/supervisor/SKILL.md`
- `.agents/skills/backend-development/SKILL.md`
- `.agents/skills/Data-Integrity/SKILL.md`
- the active Child 03 plan, evidence ledger, benchmark ledger, and actual-status ledger
- the graph-accuracy roadmap and `docs/contracts/graph-accuracy-contract.md`
- the complete Coder report and all three exact candidate files

The implementation-aware skills were applied only for inspect-only backend, conservation, and persistence review. No repair was made. The four Orchestration-owned dirty governance files were not edited and were not used as a semantic audit gate. No P3-C2 or later slice was opened. `E:\cheapapp.org` was not accessed. 1DevTool was not used. No hidden agent or separate lane was created.

Outside the Owner-mandated canonical build, all Git and Anvien review commands explicitly excluded `internal/aicontext/skills/**` and `.claude/skills/**`. The canonical script's built-in unexcluded final analyze was required by the Owner intervention; its graph was not used as review evidence and was immediately superseded by a separate forced refresh with both exclusions.

## Exact candidate freeze

All hashes matched at assignment capture, after the canonical build, and at the final pre-report freeze:

| Artifact | Assigned and final SHA-256 | Final match |
| --- | --- | --- |
| `internal/resolution/resolve.go` | `93f7b75b27c4f7ebd9061248df876195d497bf9e6a6b4852516bc45db8f1e512` | yes |
| `internal/resolution/p3c_binding_occurrence_test.go` | `af148e3ceb48a8dc5e0c5a1115685d63286c2c1e0b6e230104e67216ede7cbae` | yes |
| `internal/lbugload/p3c_binding_occurrence_persistence_test.go` | `c704b45fd350f2a1b064d79e78b4dc99f6378d44358b4e86149a09fa38d4a850` | yes |
| Coder report | `d279e1ec154752c3dc0ab31b81f4f31164e27198abf2cecb89817e77c0f096ca` | yes |

Filesystem timestamps independently support the Coder's code-first sequence:

- production: `2026-08-15T06:43:22.2307716Z`
- resolution test: `2026-08-15T06:53:11.1374211Z`
- persistence test: `2026-08-15T06:57:50.4509865Z`
- Coder report: `2026-08-15T07:11:42.2887949Z`

`git diff --check -- internal/resolution/resolve.go` was clean.

## Git and explicit-exclusion boundary

Commands included:

```text
git rev-parse HEAD
git branch --show-current
git merge-base --is-ancestor 69089d17b6a248b2a7c03f45bfa0cfec3c78d7e4 HEAD
git log --oneline 69089d17b6a248b2a7c03f45bfa0cfec3c78d7e4..HEAD
git diff --name-status 69089d17b6a248b2a7c03f45bfa0cfec3c78d7e4..HEAD -- . ":(exclude)internal/aicontext/skills/**" ":(exclude).claude/skills/**"
git status --short -- . ":(exclude)internal/aicontext/skills/**" ":(exclude).claude/skills/**"
git diff --cached --name-status -- . ":(exclude)internal/aicontext/skills/**" ":(exclude).claude/skills/**"
```

Results:

- HEAD remained `33e73ec6319fd764d415cac8024b36f0ac75d70e` on `master`.
- Assignment HEAD ancestry passed.
- The four intervening commits were Owner Orchestration documentation commits. The assignment-to-HEAD path diff was empty when both forbidden trees were excluded.
- Staged state was empty.
- The pre-report worktree outside the forbidden trees contained exactly the four Orchestration governance modifications, the three candidate code/test paths, the Orchestration handoff report, and the Coder report. The canonical build added no tracked or untracked worktree path.
- This Supervisor did not run detect-changes, `git add`, commit, push, merge, reset, or checkout.

## Graph freshness, file detail, and blast radius

`anvien --help` was run first. Before graph-based review, and again after the canonical build overwrote the index, the graph was refreshed with:

```text
anvien analyze . --force --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**"
```

Final excluded refresh: exit 0 in 40.4 seconds; scanned 1,104 files, parsed 624 code files, failed 0; 80,083 nodes and 119,011 relationships at `E:\Anvien\.anvien\graph.json`.

File inspection command:

```text
anvien file-detail internal/resolution/resolve.go --repo E:\Anvien --json --format compact --linked 10 --relationships 10 --unresolved 10
```

It reported 111 symbols, 87 inbound and 152 outbound relationships, 21 linked flows, 31 linked tests, `risk=high`, `stale=false`, and `changedSinceAnalyze=false`.

Exact upstream impact was run by canonical UID for all changed/new production symbols:

```text
anvien impact symbol --uid "Function:file30:internal/resolution/resolve.goname16:ResolveBoundIntoarity1:3occurrence65:def:internal/resolution/resolve.go#52:0:Function:ResolveBoundIntoscope56:scope:internal/resolution/resolve.go#52:0-103:1:Functionowner0:" --repo E:\Anvien --direction upstream --json
anvien impact symbol --uid "Function:file30:internal/resolution/resolve.goname27:buildBindingOccurrenceIndexarity1:1occurrence77:def:internal/resolution/resolve.go#135:0:Function:buildBindingOccurrenceIndexscope57:scope:internal/resolution/resolve.go#135:0-209:1:Functionowner0:" --repo E:\Anvien --direction upstream --json
anvien impact symbol --uid "Function:file30:internal/resolution/resolve.goname11:resolveCallarity1:4occurrence61:def:internal/resolution/resolve.go#311:0:Function:resolveCallscope57:scope:internal/resolution/resolve.go#311:0-440:1:Functionowner0:" --repo E:\Anvien --direction upstream --json
anvien impact symbol --uid "Function:file30:internal/resolution/resolve.goname13:resolveAccessarity1:4occurrence63:def:internal/resolution/resolve.go#465:0:Function:resolveAccessscope57:scope:internal/resolution/resolve.go#465:0-524:1:Functionowner0:" --repo E:\Anvien --direction upstream --json
```

All four were `CRITICAL`, a scope warning rather than an edit prohibition. `ResolveBoundInto` reached 7 impacted symbols across 5 files and 33 processes; the index builder, `resolveCall`, and `resolveAccess` each reached 6 symbols across 4 files and the same 33 processes. Health overlays reported no gaps or degraded resolution for these impact queries.

## Source and contract inspection

The production diff was 219 insertions and 11 deletions, confined to `internal/resolution/resolve.go`. Directly relevant resolution workspace/indexing, emission, source-site identity, graph JSON types, Ladybug CSV/load/schema, ScopeIR facts, and lexical-scope owners were inspected read-only.

Verified behavior:

- `buildBindingOccurrenceIndex` executes before graph selection/mutation.
- Leaf-to-definition matching requires normalized file, name, full range, and selection range, exactly one definition, one owner ID occurrence, one keyed local binding, and no duplicate accepted defID.
- `resolveScopedName` follows the actual parent chain; P3-C receiver/read/write emission accepts only a defID in the validated occurrence set. No global, same-name, or name-rescue fallback was introduced.
- Inner and outer `value` definitions retain distinct graph IDs and lexical endpoints.
- Existing `ACCESSES` semantics and source-site fields are reused.
- Graph JSON and Ladybug schemas carry all relevant node range/selection and relationship fields.
- Ladybug's default loader rejects skipped relationships and does not silently accept fallback/retry paths.
- No scanner, extraction, export/module/ambient, graph serialization, or Ladybug production file changed.

Rejected behavior:

- Owner collection records only `defID -> scope.ID`; validation never verifies that the owner scope shares the leaf/definition file or lexically contains the declaration range/selection.
- Local-binding validation keys only `(scopeID, name, defID)`. Therefore moving the owner and its matching local binding together to a wrong-file or wrong lexical scope satisfies the counts and passes validation.

## Canonical build and process boundary

The repository canonical entrypoint had SHA-256 `8db0a18864e73b32419691f5e46bdcae909e4a8d5b4442040833a05a1640b44f` and was invoked unchanged from `E:\Anvien`:

```text
$env:TEMP='E:\Anvien\.tmp\e3-p3c-review1\canonical-temp'
$env:TMP='E:\Anvien\.tmp\e3-p3c-review1\canonical-temp'
powershell -ExecutionPolicy Bypass -File .\scripts\full-build.ps1
```

Before that run, the global npm Anvien package was proven to be a junction to `E:\Anvien\anvien`. Fourteen exact `anvien mcp` executable/shell holder PIDs were identified once and terminated once. After 750 ms, remaining holders were 0, the analyze lock did not exist, and active Go/npm/esbuild/Vite build processes were 0. No repeated kill loop was used.

Canonical result: exit 0 in 112.6 seconds. npm package/runtime builds succeeded, `npm install -g .` succeeded, CLI version was `1.2.8`, the launcher/native/web build succeeded, and the canonical final analyze reported `failed=0`. All four assigned hashes then remained exact.

Earlier results from a modified temp copy of the build entrypoint were invalidated by Owner intervention and are not build evidence. The copy was removed. Only the unchanged canonical result above is evidence for `E3-P3C-BUILD1`.

## Independent tests and static checks

Focused gate:

```text
go test -count=1 -run '^TestP3C' -v ./internal/resolution ./internal/lbugload
```

Exit 0. Passed:

- `TestP3CBindingOccurrencesProjectAndResolveLexically`
- all cases in `TestP3CBindingOccurrenceProjectionFailsClosedOnOrphanOrDrift` (`missing_definition`, `missing_local_binding`, `duplicate_accepted_leaf`)
- `TestP3CBindingOccurrencesPreserveGraphJSONAndLadybugLoadParity`

Affected-package regression gate:

```text
go test -count=1 ./internal/resolution ./internal/lbugload ./internal/graph ./internal/lbugschema ./internal/providers/tsjs ./internal/scopeir
```

All six packages passed uncached.

Static gate:

```text
go vet ./internal/resolution ./internal/lbugload ./internal/graph ./internal/lbugschema ./internal/providers/tsjs ./internal/scopeir
```

Exit 0 with no diagnostics.

Independent hostile ownership gate:

```text
go run .\.tmp\e3-p3c-review1\owner-drift-probe
```

The probe intentionally exited nonzero when drift was accepted. Its result was the JSON quoted in the executive finding: the candidate accepted the cross-file owner/local-binding move, mutated the graph, and produced one unresolved reference. The temporary probe was not candidate code and was removed before verdict.

## Real Graph JSON and Ladybug boundary

A fresh repo-local TypeScript fixture reproduced the candidate's accepted P3-C source, including the import, outer and inner shadowed bindings, assignment-form destructuring, catch, loop, and destructured parameter. It was analyzed with:

```text
anvien analyze .\.tmp\e3-p3c-review1\boundary-repo --force --skip-git --name e3-p3c-review1-boundary --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**" --json
```

Exit 0 in 9.6 seconds: 3 files scanned, 2 TypeScript files parsed, failed 0, unresolved 0; graph 15 nodes / 21 relationships.

Graph JSON exact inventory:

| Definition | Full range | Selection range |
| --- | --- | --- |
| outer `rows` | `2:15-2:24` | `2:15-2:19` |
| outer `value` | `3:7-3:12` | `3:7-3:12` |
| inner `rows` | `6:17-6:26` | `6:17-6:21` |
| inner `value` | `7:11-7:16` | `7:11-7:16` |
| `caught` | `10:17-10:23` | `10:17-10:23` |
| `loop` | `11:14-11:18` | `11:14-11:18` |
| `arg` | `13:20-13:23` | `13:20-13:23` |

All seven had distinct occurrence/defID/scope-bearing node IDs and one `File:src/p3c.ts -> Variable` `DEFINES` endpoint.

| Reference | Range | Step | Exact endpoint |
| --- | --- | --- | --- |
| outer `value` read | `4:2-4:20` | 1 | outer line-3 `value` |
| assignment `value` write | `5:8-5:13` | 2 | outer line-3 `value` |
| inner `value` read | `8:4-8:22` | 1 | distinct inner line-7 `value` |
| `caught` read | `10:28-10:47` | 1 | line-10 `caught` |
| `loop` read | `11:31-11:48` | 1 | line-11 `loop` |
| `arg` read | `13:33-13:49` | 1 | line-13 `arg` |

Every relationship had a unique exact-range `SourceSite` ID, status `resolved`, `proofKind=scope-binding`, `targetRole=binding`, correct target text, `resolutionSource=scope-resolution`, confidence 1, the expected read/write reason, and the common source file hash. Counts were 7 Variables, 7 Variable `DEFINES`, 6 binding `ACCESSES` (5 reads / 1 write), 1 `IMPORTS`, 1 `USES`, 0 `ResolutionGap`, 0 `HAS_RESOLUTION_GAP`, 0 duplicate relationship IDs, 0 duplicate binding source-site IDs, and 0 orphan relationships.

Ladybug was queried independently through `anvien cypher` for all Variable IDs/ranges and all 17 binding-edge fields. It returned the same seven Variable rows and the same six relationships in the same source order, with identical IDs, endpoints, steps, source-site IDs/status, proof/role/text, ranges, resolution source, confidence, reasons, and file hash. Additional persisted counts were `DEFINES=7`, `IMPORTS=1`, `USES=1`, `ResolutionGap=0`.

```text
anvien resolution-inventory --graph .\.tmp\e3-p3c-review1\boundary-repo\.anvien\graph.json --json
```

reported 6 resolved references, 0 unresolved references, 0 source-backed unresolved, 0 unattributed unresolved, 0 ResolutionGap nodes/relationships, and zero in every gap bucket. The focused persistence test additionally passed its exact CSV relationship-column comparison and zero skipped/fallback/fallback-failure/warning assertions.

The fixture index was cleaned with `anvien clean --force`; the complete `E:\Anvien\.tmp\e3-p3c-review1` tree was then removed after exact-path and zero-reparse verification. The custom registry alias no longer appeared in `anvien list`.

## Mandatory invariant matrix

| # | Invariant | Result | Independent evidence |
| --- | --- | --- | --- |
| 1 | Occurrence conservation | PASS | Seven distinct Variable IDs and seven one-to-one `DEFINES` rows in both stores. |
| 2 | Exact ownership / fail before mutation | **REJECT** | Coordinated wrong-file owner plus local-binding move was accepted and mutated the sentinel graph. |
| 3 | Lexical resolution / no fallback | PASS | Source path uses `resolveScopedName` plus validated defID membership only; real endpoints follow scope chain. |
| 4 | Shadowing | PASS | Outer and inner `value` have distinct IDs/scopes; reads reach the nearest correct definition. |
| 5 | Reference semantics | PASS | Five reads and one write have exact ACCESS source sites, metadata, targets, and steps in JSON and Ladybug. |
| 6 | Assignment/import conservation | PASS | Assignment creates no eighth definition; it emits one write. Import remains one `IMPORTS` plus one `USES`. |
| 7 | Persistence parity | PASS | Graph JSON and Ladybug rows/counts/ranges/metadata match; focused load test reports zero skipped/fallback/warnings. |
| 8 | Gap integrity | **REJECT** | Well-formed boundary is gap-free, but the accepted malformed owner drift is not rejected and produces one binding unresolved/gap relationship. This is the direct consequence of invariant 2. |
| 9 | Scope preservation | PASS | Only the proven resolution production owner changed; extraction/scanner/persistence/export/module/ambient production remained unchanged. |
| 10 | Process boundary | PASS | Code-first timestamps; one-time holder cleanup; unchanged canonical full build passed on exact hashes; no detect/stage/commit/push. |

## External-state classification and cleanup

- Expected canonical-build side effects: global npm package refresh, launcher/protocol registration, generated ignored build outputs, and the repo-local `.anvien` index.
- The global npm package remains a junction to `E:\Anvien\anvien` as designed.
- Editor-owned `anvien mcp` processes respawned after the successful canonical build. They did not overlap the verified holder-clean build start or invalidate its completed result.
- The noncanonical temp build copy, hostile probe, boundary fixture/database, and review caches were removed. `E:\Anvien\.tmp\e3-p3c-review1` does not exist.
- No candidate, governance, forbidden-tree, or real-target byte was changed by this review. This report is the sole durable Supervisor artifact created.

## Exact repair boundary and locked bytes

Route only this repair to the Coder:

1. In `internal/resolution/resolve.go`, make binding-occurrence validation prove that the unique lexical owner agrees with the leaf/definition file and lexically contains the declaration range/selection before any graph mutation. The matching local binding must belong to that exact validated owner scope.
2. In `internal/resolution/p3c_binding_occurrence_test.go`, add a fail-before-mutation negative control that moves `OwnedDefID` and its matching local `Binding` together to a different-file or otherwise non-owning scope. It must return an error and leave a sentinel graph byte-for-byte/count-for-count unchanged.

Do not expand the repair into scanner, extraction, import, graph serialization, Ladybug production, or governance files. The current persistence test is unaffected and locked at SHA-256 `c704b45fd350f2a1b064d79e78b4dc99f6378d44358b4e86149a09fa38d4a850`. The current Coder report is also locked; a repaired candidate requires a new durable handoff and fresh exact hashes.

## Durable verdict

**Verdict: REJECT**

Orchestration main receives this verdict and must return only the exact invariant-2 owner-scope validation defect (with its invariant-8 gap consequence) to the Coder lane. P3-C must not be detected, staged, committed, or used to open P3-C2 in its current form.
