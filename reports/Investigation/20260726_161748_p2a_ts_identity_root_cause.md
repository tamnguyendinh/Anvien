# P2-A — TypeScript identity/extractor first-divergence report

## Scope and verdict

This is a read-only, bounded trace for the P1-B TypeScript extraction/identity mismatches in `E:\cheapapp.org`. It uses the real target files and the existing target graph at `E:\cheapapp.org\.anvien`; the target was not copied into `E:\Anvien`, and no target file, graph, test, or Anvien production source was changed.

The bounded verdict is **root cause confirmed at three separate source-to-graph boundaries**:

1. Array binding elements are dropped by the TS/JS definition collector before a `DefinitionFact` exists.
2. Distinct same-name local definitions are emitted into the IR, but the graph identity function removes their source range; `Graph.AddNode` then replaces the earlier node with the same ID.
3. Export visibility is never populated in the TS/JS `DefinitionFact` constructed by `addDefinition`; the graph emitter only copies a non-empty visibility field, so exported declarations remain present but unmarked.

These are bounded first-divergence findings, not a global TypeScript accuracy estimate and not a remediation authorization.

## Evidence boundary

| Item | Captured value |
|---|---|
| Target | `E:\cheapapp.org` |
| Target HEAD | `a869876ab6262dacde6cd5d432d099a91852a646` |
| Target graph | `E:\cheapapp.org\.anvien\graph.json` |
| Graph SHA-256 | `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090` |
| Graph inventory | 84,807 nodes; 114,125 relationships |
| Freshness | `anvien status`: up-to-date; `file-detail`: `stale=false`, `changedSinceAnalyze=false` |

The independent TypeScript oracle and graph comparison are retained in:

- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p1b-identity-oracle.mjs`
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p1b-identity-oracle-output.json`
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2a-identity\probe_identity.go`
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2a-identity\probe_identity_output.txt`

## Reproduction facts

The TypeScript 5.9.3 oracle, built from the target `tsconfig.json`, found:

- six legal array-pattern bindings (`messageRows`, `attemptRows`, `eventRows`, `providerEventRows`, `readinessRows`, `suppressionRows`) at lines 503–509 of `modules/email/server/operations/email-operations-observability.ts`;
- two distinct `time` declarations at lines 207 and 214, and two distinct `now` declarations at lines 262 and 501;
- 21 exported declaration names across the three selected files.

The fresh graph comparison found zero matching definition nodes for the six pattern names, one graph node at the first range for `time`, one at the first range for `now`, and zero `exported`/`visibility` properties on all 21 selected exported definitions. `file-detail` independently reports `exportedSymbolCount=0` for each selected file.

The direct real-source parser/extractor probe adds the crucial boundary evidence:

- both `time` definitions and both `now` definitions are present in the extracted `ScopeIR`, with distinct ranges;
- no definition named any of the six array elements is present in the extracted `ScopeIR`;
- `DefinitionFact.Visibility` is empty for every selected definition;
- the probe's graph hash before and after remained unchanged.

## First divergence A — array binding elements are rejected by the collector

The source AST represents the six names as the `name` child of a `variable_declarator` whose node kind is an array binding pattern. In `internal/providers/tsjs/definitions.go:64-68`, the collector does:

```go
nameNode := child(node, "name")
if nameNode == nil || nameNode.Kind() != "identifier" {
    return
}
```

The early return occurs before `addDefinition` is called. Therefore the first wrong state is an absent `DefinitionFact`, not a graph serialization or query projection loss. The graph cannot emit a definition that the extractor never produced.

## First divergence B — range is removed from graph identity

The direct extractor probe shows both `time` facts and both `now` facts, with source ranges:

- `time`: lines 207 and 214;
- `now`: lines 262 and 501.

`internal/providers/tsjs/definitions.go:95-111` creates a range-bearing `DefinitionFact` and a range-bearing internal `defID` (`defID` at lines 141–143). However, cross-file resolution converts the fact to a graph identity in `internal/resolution/indexes.go:814-824`:

```go
name := def.QualifiedName
...
return graph.GenerateID(string(def.Label), cleanPath(def.FilePath)+":"+name+arity)
```

The range is not part of this ID. Both same-file `Variable` facts therefore become `Variable:modules/email/server/operations/email-operations-observability.ts:time` (and likewise for `now`). `internal/graph/types.go:96-104` treats a duplicate ID as replacement:

```go
if index, ok := g.nodeIndex[node.ID]; ok {
    g.Nodes[index] = node
    return
}
```

Thus the first graph-level divergence is identity construction, followed by deterministic last-write replacement. The source and IR are not collapsed; the persisted graph is.

## First divergence C — export visibility is not propagated

`scopeir.DefinitionFact` declares a `Visibility` field (`internal/scopeir/facts.go:3-25`), and `internal/resolution/emit.go:171-174` only adds the graph `visibility` property when that fact is non-empty. The TS/JS collector's `addDefinition` literal at `internal/providers/tsjs/definitions.go:100-111` does not assign `Visibility`. The `emitDefinitionKind` cases at lines 15–75 also do not inspect an enclosing export statement or pass an export/visibility value into `addDefinition`.

Consequently the definition node survives, but the export fact is absent at the IR boundary. This is why all 21 selected source exports are present as definitions while `file-detail` reports zero exported symbols.

## Candidate-owner impact evidence

Fresh self-graph impact commands were run against the Anvien index (HIGH/CRITICAL values are workflow warnings, not edit prohibitions):

| Owner | Bounded impact result |
|---|---|
| `collector.addDefinition` | direct result `LOW`; file risk high; graph traversal does not expose the full provider dispatch path |
| `collector.emitDefinitionKind` | direct result `LOW`; same dispatch-visibility limitation |
| `collector.emitExportStatement` | direct result `LOW`; it records re-export imports but does not add local declaration visibility |
| `graphIDForDef` | `CRITICAL`; 3 affected modules and 18 affected processes in the self-graph projection |
| `Graph.AddNode` | `CRITICAL`; 16 directly affected symbols, 6 affected modules, and 82 affected processes in the self-graph projection |

The LOW collector results are not evidence that the collector is unused; the direct real-source extraction probe proves it runs and emits the two same-name facts. They are a limitation of the graph's interface/dispatch coverage.

## Alternatives ruled out

| Alternative | Current evidence |
|---|---|
| stale target graph | target status is up-to-date; hash and commit are bound; hash stayed unchanged during probes |
| target file omission | all three selected files are parsed and present in `file-detail` |
| parser failure | the real-source extractor returns the exact `time`/`now` facts, and the TS oracle returns all six pattern bindings |
| `file-detail` display cap | raw graph comparison has one `time`/`now` node and no six pattern-name nodes; selected exports lack metadata in raw definitions |
| target or synthetic fixture contamination | every probe path is `E:\cheapapp.org`; no fixture is used; artifacts are under `E:\Anvien` |

## Limits

- The source oracle covers three fixed files, not every TS/JS binding form or every export syntax.
- The report does not decide the desired identity policy for overloads, generated declarations, or namespace members.
- It does not assess downstream context/impact/process projections (that is P2-C), remediation design, or build/test behavior.
- No production code, test, target graph, or target worktree was edited.

## Slice decision

`P2-A` TypeScript identity/extractor root causes are **bounded and confirmed** for the fixed P1-B cases. The evidence remains subject to independent Supervisor review before the investigation is closed.
