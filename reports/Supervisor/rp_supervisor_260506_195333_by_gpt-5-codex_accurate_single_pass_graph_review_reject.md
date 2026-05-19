# Supervisor Report: Accurate Single-Pass Graph Plan Review

Verdict: REJECT

Scope reviewed:
- Plan: `docs/plans/2026-05-06-accurate-single-pass-graph-plan.md`
- Code window: `837853d..8ce9317`
- Authority used for this verdict: the specified plan plus current code/tests only. No SPEC, architecture doc, or other markdown authority was used for the conclusion.

## Blocking Findings

### [HIGH] Scope-resolved member edges can point at non-existent `def:*` ids

File: `avmatrix/src/core/ingestion/emit-references.ts:127`

Issue: `emitReferencesToGraph` maps a scope `SymbolDefinition` to a graph node through `createGraphNodeResolver`, but when semantic matching is ambiguous it falls back to `def.nodeId` (`emit-references.ts:398`). For TypeScript methods/properties, the AST capture emits only `@declaration.owner` (`typescript-javascript.ts:143`, `typescript-javascript.ts:156`, `typescript-javascript.ts:448`) and does not emit a qualified member name. `ScopeExtractor` therefore falls back to the simple name (`scope-extractor.ts:495`, `scope-extractor.ts:505`). If two classes in the same file both declare `save()`, the resolver sees two graph `Method` nodes with `name=save`, cannot choose uniquely, and emits `CALLS` to `def:src/app.ts#...:Method:save`.

Repro evidence run against current head:

```json
{
  "methodDefs": [
    {
      "nodeId": "def:src/app.ts#1:10:Method:save",
      "qualifiedName": "save",
      "ownerId": "def:src/app.ts#1:0:Class:A"
    },
    {
      "nodeId": "def:src/app.ts#2:10:Method:save",
      "qualifiedName": "save",
      "ownerId": "def:src/app.ts#2:0:Class:B"
    }
  ],
  "relationships": [
    {
      "type": "CALLS",
      "sourceId": "Function:src/app.ts:run",
      "targetId": "def:src/app.ts#1:10:Method:save"
    },
    {
      "type": "CALLS",
      "sourceId": "Function:src/app.ts:run",
      "targetId": "def:src/app.ts#2:10:Method:save"
    }
  ]
}
```

Why this blocks the plan: the plan requires the default graph to emit accurate scope-aware edges and to remain queryable after DB load. `def:*` is not a valid persisted node label in this graph path. The DB load path derives labels with `nodeId.split(':')[0]` (`lbug-adapter.ts:417`) and skips invalid relationship label pairs (`lbug-adapter.ts:104`, `lbug-adapter.ts:106`; fallback also skips them at `lbug-adapter.ts:669`, `lbug-adapter.ts:671`). These resolved edges can be silently dropped during persistence.

Fix direction:
- Emit owner-qualified declaration names for TypeScript/JavaScript members, e.g. `A.save` / `B.save`, in the AST scope-capture path.
- Resolve graph nodes by a stable owner-qualified key or by owner membership, not by `(label, filePath, simpleName)` alone.
- Never add a relationship whose mapped `sourceId` or `targetId` is absent from the graph; count it as skipped/missing instead.
- Add a regression covering two same-file classes with same method/property names and verify emitted relationships target real graph nodes and survive CSV/LadybugDB load.

### [HIGH] Duplicate edge guard discards scope audit metadata for legacy-overlap edges

File: `avmatrix/src/core/ingestion/emit-references.ts:132`

Issue: the duplicate guard builds a semantic edge key from existing graph relationships (`emit-references.ts:504`, `emit-references.ts:512`) and simply skips the scope-resolved edge when that key already exists (`emit-references.ts:132`). The existing legacy edge is left unchanged. That means the single persisted edge for the resolved relation can still have no `resolutionSource`, no `evidence`, and no `fileHash`, even though the scope resolver had that audit data.

Repro evidence run against current head:

```json
{
  "emitStats": {
    "edgesEmitted": 1,
    "skippedDuplicateEdge": 1
  },
  "calls": [
    {
      "id": "legacy-call",
      "sourceId": "Function:src/app.ts:run",
      "targetId": "Method:src/app.ts:A.save",
      "confidence": 0.5,
      "reason": "legacy call"
    }
  ]
}
```

Why this blocks the plan: storage/readback plumbing exists (`schema.ts:427`, `csv-generator.ts:620`, `graph-read-service.ts:38`, `graph-read-service.ts:101`, `local-backend.ts:161`, `local-backend.ts:179`), but it only helps when the emitted relationship carries audit fields. In the default pipeline, common `CALLS`/`USES`/`ACCESSES` overlaps can keep the legacy edge and drop the scope audit edge, so query/context/impact after DB load are not guaranteed to expose the scope-resolution evidence the plan requires.

Fix direction:
- Keep the non-duplicating behavior, but merge or upgrade the existing relationship with the scope-resolved `resolutionSource`, `evidence`, `fileHash`, and preferably the higher confidence/reason.
- Add a test where a legacy edge already exists, `emitReferencesToGraph` sees the scope-resolved equivalent, and the graph still has one relationship with scope audit metadata.
- Extend the test through CSV generation/load/readback for at least one overlapped `CALLS` edge.

### [HIGH] Full plan validation gate is not passing on current head

File: `avmatrix/test/integration/cli-e2e.test.ts:175`

Issue: the plan lists `cd avmatrix && npm test` as a validation command. On current head, targeted scope/benchmark tests pass, but full `npm test` fails. The failures include `cli-e2e` timing out on `analyze command runs pipeline on mini-repo`, and multiple `skills-e2e` `analyze --skills` runs exiting with code `3221226505`.

Observed result:

```text
Test Files  2 failed | 220 passed | 1 skipped (239)
Tests       11 failed | 5143 passed | 98 skipped (6690)
Errors      17 errors
```

Fix direction:
- Reproduce `test/integration/cli-e2e.test.ts` and `test/integration/skills-e2e.test.ts` in isolation.
- Determine whether the analyze timeout/process exit is caused by the new resolution path, worker lifecycle, or an existing environment/native parser issue.
- Re-run full `npm test` after the fix; targeted tests alone are not enough for approval because the touched code is in the default analyze path.

## Source-Level Clearance Notes

- Pipeline and parse worker path: clear for the current slice. `resolutionPhase` is in the default phase list after `crossFilePhase` (`pipeline.ts:64`, `pipeline.ts:75`); parse reads file contents once per chunk before worker parsing (`parse-impl.ts:221`); worker scope extraction reuses the existing tree root (`parse-worker.ts:1482`, `parse-worker.ts:1487`); parse phase preserves `parsedFiles` (`parse-impl.ts:367`, `parse-impl.ts:533`), finalizes once (`parse-impl.ts:480`), and attaches indexes (`parse-impl.ts:487`).
- Scope artifact and resolver path: clear with caveat at emission boundary. `ParsedFile` carries `fileHash`, local defs, imports, and reference sites (`parsed-file.ts:47`); resolver builds deterministic chunks serially for now (`scope-reference-resolver.ts:125`, `scope-reference-resolver.ts:129`), preserves per-reference file hash (`scope-reference-resolver.ts:196`), and builds a `ReferenceIndex` (`scope-reference-resolver.ts:253`). The blocker is graph-node mapping after resolution.
- TypeScript/JavaScript scope capture path: blocked by finding 1. Owner metadata is preserved for methods/properties (`typescript-javascript.ts:143`, `typescript-javascript.ts:156`), but no owner-qualified declaration name is emitted in `emitNamedDeclaration` (`typescript-javascript.ts:445`), so `ScopeExtractor` falls back to simple names (`scope-extractor.ts:505`).
- Graph emission path: blocked by findings 1 and 2. `emitReferencesToGraph` emits mapped relationships (`emit-references.ts:127`, `emit-references.ts:136`), but ambiguous graph node resolution falls back to `def.nodeId` (`emit-references.ts:398`) and duplicate edges are skipped instead of metadata-merged (`emit-references.ts:132`).
- Persistence/readback/MCP audit plumbing: clear as storage plumbing, not as full behavior coverage. Relationship schema has `resolutionSource`, `evidence`, and `fileHash` (`schema.ts:427`); CSV writes those columns (`csv-generator.ts:620`); fallback inserts them (`lbug-adapter.ts:681`); graph readback maps them (`graph-read-service.ts:101`); MCP context/impact helpers return them (`local-backend.ts:161`, `local-backend.ts:2471`). Findings above prevent all resolved edges from reliably reaching this plumbing with valid targets and metadata.
- Benchmark/metrics CLI: clear for the slice. `analyze --benchmark-json` writes a benchmark snapshot (`analyze.ts:331`, `analyze.ts:342`); key metrics include phase and scope-resolution counters (`analyze-benchmark-snapshot.ts:181`, `analyze-benchmark-snapshot.ts:210`); `benchmark-compare` is wired (`index.ts:55`, `benchmark.ts:36`).
- Workerization scaffold: clear as scaffold only. `workerData` support exists for future readonly indexes (`worker-pool.ts:73`, `worker-pool.ts:119`), while resolution remains serial chunk execution on current head (`scope-reference-resolver.ts:129`). This matches an intermediate implementation slice, not final plan completion.

## Validation Run

Passed:
- `git diff --check 837853d..HEAD`
- `cd avmatrix && npx tsc --noEmit`
- `cd avmatrix && npx vitest run test/unit/scope-resolution/typescript-single-pass-parity.test.ts test/unit/scope-resolution/scope-reference-resolver.test.ts test/unit/scope-resolution/emit-references.test.ts test/unit/scope-resolution/resolution-phase.test.ts test/unit/analyze-benchmark-snapshot.test.ts test/unit/benchmark-compare-command.test.ts` (`6` files, `32` tests passed)

Failed:
- `cd avmatrix && npm test` (`11` failed tests, `17` unhandled worker errors)

## Required Fix List For Resubmission

1. Fix scope-to-graph member node mapping so same-name methods/properties in the same file resolve to real, distinct graph nodes.
2. Add fail-closed emission checks so unresolved graph-node mapping never creates `def:*` relationship endpoints.
3. Merge or upgrade audit metadata on semantic duplicate edges instead of discarding scope-resolved metadata.
4. Add regression tests for ambiguous same-name member mapping and duplicate-edge audit metadata preservation, including persistence/readback for at least one overlapped edge.
5. Make the full `npm test` gate pass or produce a narrowed, reproducible root-cause report if an unrelated pre-existing environment failure is proven.

## Overall Coder Evaluation

The implementation style is additive and mostly follows the plan's staged approach: reuse shared `ParsedFile`/`ReferenceIndex` contracts, preserve worker-produced scope facts, expose counters, and add benchmark tooling. The remaining problems are not style issues; they are runtime correctness and auditability gaps at the graph emission boundary. Because the default analyze path can now emit invalid relationship endpoints and can drop scope audit metadata on overlapped edges, this batch is not approvable yet.
