# Architect Review: Target-Scoped Compact Impact

**Recipient:** User / Owner  
**Review time:** 2026-08-30 23:39:23 +07:00  
**Scope:** Standalone review of the attached proposal about context explosion in `impact` and `file-detail`  
**Code/source inspection:** Not performed by design  
**Verdict:** **CONFLICT — the problem and part of the direction are valid, but the proposal is not implementation-ready**

## Authority and evidence used

- `AGENTS.md` hard rules.
- Review target: `C:\Users\TAM NGUYEN\.codex\attachments\c0ce9630-65c9-4be8-8366-30d9731dc249\pasted-text.txt`.
- Active graph integrity authority: `docs/contracts/graph-accuracy-contract.md` and its linked roadmap.
- Existing file-detail compact plan, evidence, benchmark, actual-status, and Supervisor reports as historical implementation evidence, not as architecture authority.
- No `SPEC-MAP.md` or `Docs/SPEC/*` family exists in the current repository tree, so this report cannot issue a formal SPEC-alignment PASS. It evaluates internal coherence, contract safety, and consistency with the available active graph contract.

## Executive conclusion

The reported context-explosion problem is credible and worth solving. Three ideas are directionally sound:

1. compute impact reachability before shaping the response;
2. replace stringified JSON proof blobs with structured proof data;
3. use a compact, versioned representation for high-volume repeated facts.

The current proposal nevertheless combines three different changes without separating their contracts:

- **semantic scope reduction** — excluding facts outside a defined impact closure;
- **lossless representation compaction** — dictionaries, row tables, tuple ranges, structured proof;
- **public response migration** — replacing existing MCP fields and shapes.

Those changes have different correctness and compatibility requirements. Treating them as one “unique solution” makes the proposal unsafe.

## Parts that are reasonable

### 1. Projection after traversal is the right architectural order

The proposed sequence BFS → target-scoped projection → proof normalization → compact encoding is sound only if projection is a response-layer operation. It must not prune traversal, change risk, change impacted counts, change affected files/process discovery, or hide possible-impact boundaries.

### 2. Structured proof is better than stringified JSON

Parsing a JSON string stored in `evidence[].note` into a typed object removes escaping, improves validation, and makes fields addressable. This is a good change if every contract-relevant field remains present or losslessly retrievable.

### 3. Compact row tables can materially reduce repetition

The repository already has historical evidence that versioned dictionaries and row tables reduce `file-detail` payload size while preserving expanded compatibility. The measured reduction was material, though not uniformly 95%:

- compact model: `849,155 → 281,704` characters, about 66.8%;
- CLI: `1,282,361 → 282,588` characters, about 78.0%;
- HTTP: `615,355 → 282,949` characters, about 54.0%.

These are different fixtures from the proposal's target, so they validate the technique, not the claimed impact benchmark.

## Blocking findings

### [CRITICAL] The defined closure contradicts the zero-data-loss invariant

The proposal first requires all caller/callee nodes across the impact chain and all target-chain evidence, but the concrete rule later retains only the target plus callers at depth 1. That can omit depth ≥2 nodes, callees, relationships between closure nodes, unresolved facts at deeper nodes, and alternate paths.

The contract must define an exact closure such as:

`C(target, direction, maxDepth, traversalRelationKinds, graphSnapshot)`

It must also distinguish traversal edges from evidence-only edges and define cycles, minimum depth, parallel edges, multiple paths, and deterministic ordering.

### [CRITICAL] The example output itself demonstrates truncation

The proposed response reports:

- `processes_affected: 10` but supplies one process;
- `scopedUnresolved.total: 5` but supplies one item;
- `impactedCount: 2` while the shown `byDepth` contains one impacted node.

Without explicit pagination, omitted counts, or a statement that the example is non-normative and abbreviated, this violates the proposal's own “no mechanical truncation” and “100%” invariants.

### [CRITICAL] Proof normalization is lossy as specified

The old proof example contains `schemaVersion`, `fileHash`, a full range, source-site identity, resolution state, and proof detail. The replacement contains only `kind`, `sourceSiteId`, `status`, `line`, and `col`. Removing fields is not normalization; it is a contract change and possibly provenance loss.

The active graph contract requires record-level parity and says matching counts alone do not prove parity. Every old contract field must either map to the new schema or remain losslessly retrievable through a stable proof reference tied to the same graph snapshot.

### [HIGH] Target-scoping cannot claim complete impact while unresolved boundaries remain undefined

Filtering gaps only where `sourceSymbol == target/closure node` can miss an unresolved inbound source outside the resolved closure that may actually target the symbol. Dynamic dispatch, reflection, cross-language links, multiple candidates, confidence, and source sites without a stable owning symbol are also undefined.

The response needs explicit states such as `analysisCompleteness`, `possibleImpact`, and `unknownBoundary`. “Zero loss” can mean lossless relative to a specified indexed graph snapshot; it cannot mean runtime-complete blast-radius knowledge when unresolved evidence exists.

### [HIGH] The proposal is a breaking MCP contract with no migration path

The proposal changes at least these shapes:

- `evidence[]` to a single `proof` object;
- `fileLayer.relationships/unresolved` to scoped replacements;
- named objects to dictionary-indexed positional rows.

No root response version, `format=expanded-v1|compact-v2`, capability negotiation, dual-output window, deprecation policy, consumer inventory, or rollback rule is defined. Updating two assertion files is not enough to prove compatibility across CLI, MCP, resources, scripts, and downstream agents.

Historical repository evidence is especially relevant: the prior file-detail compact work added explicit `format`/`version`, dictionaries, schemas, row tables, parity tests, and deliberately preserved the MCP/agent file-context expanded contract. The new proposal must explicitly decide whether it is extending that contract or introducing a separately versioned impact contract.

### [HIGH] The declared dense schema and the “after” example are different designs

The specification requires `dict` plus positional vectors, while its final example still uses expanded named objects and contains no dictionary or row schema. The example therefore cannot substantiate the 8–12 KB estimate or define the consumer contract.

A positional format also needs canonical column metadata, types, enum tables, null/sentinel rules, dictionary layout, stable index assignment, deterministic sorting, malformed-index behavior, and decoder validation. A hybrid format may be better for an LLM: named objects for the small high-value impact core, compact tables only for bulky evidence.

### [HIGH] The numerical result is a forecast, not evidence

`215 KB → 8–12 KB` is arithmetically a 94.4–96.3% reduction, so “about 95%” is plausible as a target. It is not yet a measured result because the proposal itself lists the benchmark as future work.

The benchmark must hold constant:

- graph/index revision and target;
- direction, depth, and relation policy;
- compact versus pretty JSON serialization;
- exact field/value parity;
- tokenizer/model;
- repeated-run deterministic bytes.

It should separately attribute savings to scoping, proof normalization, and tabular encoding. Escaped quotes alone are unlikely to explain most of the claimed reduction.

### [MEDIUM] `file-detail` and `impact` are conflated

The title scopes both tools, but the proposed implementation is primarily an `impact` response redesign. `file-detail` already has a compact/full-detail contract and historical compatibility decisions. A new change should not reopen or silently replace that contract unless a separate current measurement proves a remaining file-detail problem.

## Required architecture decisions before implementation

1. **Closure semantics:** define target, direction, maximum depth, traversal relationship kinds, evidence-only kinds, cycle/multi-path behavior, and unknown/possible-impact handling.
2. **Projection boundary:** require full closure and risk/process discovery before any response filtering.
3. **Lossless meaning:** define field-level semantic equivalence for nodes, edge multiset, gaps/candidates, process-step hits, ranges, proof provenance, and counts.
4. **Contract versioning:** retain expanded v1 and introduce opt-in compact v2, or provide an equally explicit migration and capability-negotiation policy.
5. **Canonical compact schema:** choose either named compact objects, positional tables, or a hybrid; publish one machine-checkable schema and deterministic ordering rules.
6. **Evidence delivery:** decide whether all proof is inline or some proof is represented by stable, snapshot-bound resource handles. Either approach can be lossless; therefore “unique solution” is not justified.

## Acceptance oracle

For identical target, graph snapshot, direction, depth, and relation policy, require:

`decode(compact(project(expandedClosure))) == canonical(project(expandedClosure))`

The equality must cover field values and multiplicity, not only totals:

- exact node set;
- exact edge multiset, including parallel edges;
- minimum depth and alternate-path evidence;
- unresolved/ambiguous candidate set;
- process-step hits;
- proof/provenance fields;
- affected files, risk, and counts;
- deterministic byte output across repeated runs.

## Final verdict

- **Problem statement:** accepted as credible, but the quoted source-level root cause and measurements were not independently verified in this SPEC-only review.
- **Target-scoped response projection:** accepted in principle, only after full impact discovery and with a precise closure contract.
- **Structured proof:** accepted in principle, but the shown mapping is not lossless.
- **Dense compact representation:** viable as a versioned or negotiated format; not proven to be the only or best default MCP format.
- **Current proposal as an implementation specification:** **rejected / not ready**.
- **Residual ADR boundary:** closure semantics, compact impact contract versioning, and inline-versus-retrievable proof ownership require explicit Owner decisions.

No production code, tests, SPEC, plan, or runtime files were modified during this review.
