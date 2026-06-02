# Anvien Web UI Full Graph Render Slowness

Date: 2026-06-01
Repository: `E:\Anvien`
Scope: problem report only; no code changes were made.
Status: open

## 1. Problem Summary

The Web UI is slow when opening the Anvien repository graph.

The main issue is not ordinary React rendering. The current Web UI loads the full semantic graph into the browser, parses it, builds an in-memory `KnowledgeGraph`, converts the whole graph into Graphology/Sigma data, applies layout, and only then filters what is visible.

For the Anvien repo, the initial Web graph load is too large for this path.

## 2. Current Measured Size

Fresh analyze state during the investigation:

```text
Indexed commit: fd8f05f
Status: up-to-date
Files scanned: 816
Parsed code: 598
Graph nodes: 96,212
Graph relationships: 131,685
.anvien/graph.json size: 327,985,627 bytes
```

Graph-health inventory:

```text
ResolutionGap nodes: 69,065
ResolutionGap count: 69,988
File layer files: 816
Unresolved files: 590
```

## 3. Benchmark Evidence

Backend/API measurements:

```text
GET /api/repos
  status: 200
  time: ~523 ms
  size: 433 bytes

GET /api/graph?repo=Anvien
  status: 200
  time: ~23.4 s
  size: 360,949,816 bytes

GET /api/graph?repo=Anvien&stream=true
  status: 200
  time: ~22.5 s
  size: 367,980,803 bytes
```

Browser benchmark against the running local Web UI:

```text
Browser DOMContentLoaded: ~1.18 s
Browser Ready time: ~78.1 s
/api/graph browser download: ~19.2 s
GraphCanvas conversion/layout: ~33.5 s
Longest main-thread long task: ~55.4 s
Total long-task time: ~58.0 s
Frame drops recorded: 91
JS heap used after load: ~662 MB
```

Runtime diagnostics after load:

```text
graphConversion.lastNodeCount: 96,212
graphConversion.lastRelationshipCount: 131,685
graphConversion.lastMs: ~33,489 ms
visibleViewportNodeCount after default filters: 8,569
layout.starts: 0
manualOptimizerInvocations: 0
```

The important ratio is:

```text
loaded graph nodes: 96,212
default visible overview nodes: 8,569
```

Most loaded nodes are not visible in the default Web overview.

## 4. Root Cause

### 4.1 Full graph payload is sent to the browser

`internal/httpapi/graph.go` streams all graph nodes and all graph relationships from `/api/graph`.

Relevant owners:

```text
internal/httpapi/graph.go: handleGraph
internal/httpapi/graph.go: streamGraphNDJSON
internal/httpapi/graph.go: graphPayload
```

The `stream=true` mode improves progress visibility, but it does not reduce payload size. The browser still receives the whole graph.

### 4.2 Frontend accumulates the whole stream before render

`anvien-web/src/services/backend-client.ts` parses the NDJSON stream into arrays:

```text
parseNdjsonGraphResponse()
  nodes.push(record.data)
  relationships.push(record.data)
  return { nodes, relationships, semanticStatus }
```

This is not progressive rendering. It is progressive download followed by full in-memory parse/accumulation.

### 4.3 Frontend then builds another full graph model

`anvien-web/src/hooks/useAppState.local-runtime.tsx` builds a `KnowledgeGraph` after the download:

```text
const newGraph = createKnowledgeGraph(result.semanticStatus)
for (const node of result.nodes) newGraph.addNode(node)
for (const rel of result.relationships) newGraph.addRelationship(rel)
setGraph(newGraph)
```

This duplicates work after the browser already accumulated arrays from the stream.

### 4.4 GraphCanvas converts and lays out the whole graph before filtering

`anvien-web/src/components/GraphCanvas.tsx` converts the full graph:

```text
const sigmaGraph = knowledgeGraphToGraphology(graph, communityMemberships)
recordGraphConversion(...)
setSigmaGraph(sigmaGraph)
```

`knowledgeGraphToGraphology()` in `anvien-web/src/lib/graph-conversion.ts`:

```text
adds all nodes
applies clustered layout
adds display relationships
```

Only after this does `GraphCanvas` apply label/depth/semantic/graph-health filters. This means hidden default nodes still pay the download, parse, memory, conversion, and layout cost.

## 5. Main Data Driver

`ResolutionGap` dominates the graph:

```text
ResolutionGap nodes: 69,065
Total graph nodes: 96,212
```

These nodes are important diagnostic data, but they do not need to be loaded into Sigma for the initial overview. They are a better fit for:

- graph-health summaries;
- file/detail diagnostics;
- on-demand drilldown;
- filtered/lazy graph layers.

## 6. Why The UI Feels Slow

The visible delay combines multiple costs:

1. Backend reads and serializes a very large graph snapshot.
2. Browser downloads roughly 361-368 MB.
3. Browser parses NDJSON into arrays.
4. Browser builds `KnowledgeGraph`.
5. Browser converts full graph into Graphology/Sigma.
6. Browser applies layout for all nodes before default visibility filters.
7. The main thread is blocked by long synchronous work; the longest observed long task was about 55 seconds.

The result is a long "Loading graph" / "Processing" wait before the UI reaches Ready.

## 7. What Is Not The Main Cause

The investigation did not find evidence that automatic layout optimizer is the cause.

Runtime diagnostics showed:

```text
layout.starts: 0
layout.stops: 0
manualOptimizerInvocations: 0
```

The slow path happens before any manual layout optimizer action.

## 8. Suggested Direction For A Future Plan

This report is not the implementation plan, but the likely direction is:

1. Add a small default graph overview endpoint for Web initial load.
2. Do not include `ResolutionGap` nodes in the default Sigma graph.
3. Return summary counts and diagnostics separately from rendered nodes.
4. Lazy-load diagnostic nodes and detail relationships when the user opens a file, graph-health lens, or explicit diagnostic filter.
5. Filter before Graphology conversion, not after.
6. Avoid sending or constructing relationships that are not visible in the current graph mode.
7. Consider worker-side graph parsing/conversion if full graph mode is still needed.

## 9. Current Owner Areas

Backend/API:

```text
internal/httpapi/graph.go
internal/filecontext
internal/graphhealth
```

Frontend data loading:

```text
anvien-web/src/services/backend-client.ts
anvien-web/src/hooks/useAppState.local-runtime.tsx
```

Frontend rendering/conversion:

```text
anvien-web/src/components/GraphCanvas.tsx
anvien-web/src/lib/graph-conversion.ts
anvien-web/src/lib/graph-layout.ts
anvien-web/src/hooks/useSigma.ts
```

## 10. Closure Criteria For A Future Fix

A future implementation should be considered successful only if measured numbers improve, for example:

```text
Initial Web graph payload: much smaller than current ~361 MB
Browser Ready time: materially below current ~78 s
GraphCanvas conversion/layout: materially below current ~33.5 s
Longest main-thread long task: materially below current ~55 s
Default overview still shows useful architecture context
ResolutionGap diagnostics remain accessible through drilldown/lens workflows
```

