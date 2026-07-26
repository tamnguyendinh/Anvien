# Cheapapp Graph Root-Cause Causal Synthesis (Restart)

Status: bounded causal synthesis complete; no remediation authorization

## Metadata

- Review time: `2026-07-26 16:53:17 +07`
- Investigator/reviewer: `gpt-5-6-sol`
- Target repository: `E:\cheapapp.org`
- Target HEAD: `a869876ab6262dacde6cd5d432d099a91852a646`
- Target graph: `E:\cheapapp.org\.anvien\graph.json`
- Target graph SHA-256: `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`
- Target graph inventory: `1,359` files, `84,807` nodes, `114,125` relationships
- Anvien source HEAD: `107e0157ec072e54b44246719da3e7accf76e1cb`
- Plan: `docs/plans/2026-07-26-cheapapp-graph-root-cause-restart/`
- Authority: current user boundary, `E:\Anvien\AGENTS.md`, current target source/graph, independently rerun probes, and current Anvien source

## Executive conclusion

The restart does not show one universal “graph projection bug.” It shows several independent first divergences at different pipeline boundaries:

1. directory discovery prunes legitimate nested source directories before a File candidate exists;
2. the TypeScript extractor rejects array binding patterns;
3. graph identity omits range/scope and collides distinct local definitions;
4. TypeScript export visibility is not populated in the provider fact;
5. the resolver workspace has no TypeScript ambient/lib declaration universe;
6. imported-definition lookup stops at a barrel's physical definitions instead of following its re-export binding.

For the fixed downstream cases, `context`/`impact` and `file-detail` mostly expose the graph facts they receive. They do not recover upstream missing calls. The database/native projection changes selected representation (`nil step -> 0`, scalar -> string, and a reduced unresolved DTO) without a measured cardinality loss. The process control is a bounded no-change observation, not evidence of process completeness or product semantics.

This synthesis is intentionally bounded. It is not a global accuracy percentage, a complete TypeScript conformance result, a product execution-model verdict, a performance diagnosis, or approval to edit code.

## Boundary and baseline

- The target was analyzed in place. Its normal graph/index remains at `E:\cheapapp.org\.anvien`; the target was not copied into `E:\Anvien`.
- Reports, ledgers, probes, and temporary files are under `E:\Anvien`; no investigation artifact was written into the target. The user-removed `E:\Anvien\p0-rc-c-fixture` remains absent.
- The target had pre-existing Git changes. The investigation preserved them and did not edit target source. The supported target analyze had an acknowledged operational side effect: ignored `AGENTS.md`/`CLAUDE.md` timestamps may be touched; no content-delta claim is made without a pre-run baseline.
- The current graph hash and target HEAD stayed unchanged through the read-only P2-A/P2-B/P2-C reruns. Historical reports from the earlier context-contaminated checkpoint are evidence pointers only; their different graph hash/counts are not mixed into this synthesis.

## Causal matrix

| ID | Bounded symptom / oracle | Earliest confirmed divergence | Graph/output effect | Classification | Evidence |
|----|--------------------------|-------------------------------|--------------------|----------------|----------|
| C1 | 895 eligible TS/JS source paths vs 887 graph File nodes; eight exact paths missing, zero graph-only paths | `ShouldIgnorePath` treats any `env`, `target`, or `logs` path segment as ignored; `WalkRepositoryPaths` calls `filepath.SkipDir` before candidacy | Entire nested subtrees never reach parser/File-node creation | **CONFIRMED WRONG — bounded scanner slice** | `E1-P1A-FILE1`, `E1-P1A-FILE2`, `E1-P1A-CMP1`, `E2-P2A-SRC1`, `E2-P2A-SRC2`, `E2-P2A-CAUSE1`, `E2-P2A-REPORT1` |
| C2 | Six legal array-binding names (`messageRows`, `attemptRows`, `eventRows`, `providerEventRows`, `readinessRows`, `suppressionRows`) exist in TS AST but have no graph definition | TSJS `variable_declarator` handling returns when `name.Kind() != identifier` at `internal/providers/tsjs/definitions.go:64-68` | No `DefinitionFact` or local binding is created; downstream uses become gaps | **CONFIRMED WRONG — bounded extractor slice** | `E1-P1B-SRC1`, `E1-P1B-CMP1`, `E2-P2A-TSIR1`, `E2-P2A-REPORT2`, `E2-P2A-REVIEW1` |
| C3 | Four distinct local facts (`time` x2, `now` x2) survive ScopeIR but only two graph nodes remain | `graphIDForDef` uses file/name/arity and omits range/scope; `Graph.AddNode` replaces duplicate IDs | Distinct source identities collapse at graph persistence; the observed winner is order-dependent and not a valid identity policy | **CONFIRMED WRONG — bounded identity slice** | `E1-P1B-CMP1`, `E2-P2A-TSIR1`, `E2-P2A-IDENTITY1`, `E2-P2A-REVIEW1` |
| C4 | 21 selected top-level exports are represented as definitions but zero graph rows carry export/visibility metadata | TSJS `addDefinition` does not fill `DefinitionFact.Visibility`; emitter serializes visibility only when non-empty | Exported declarations appear present but unexported to `file-detail`/consumers | **CONFIRMED WRONG — bounded metadata slice** | `E1-P1B-CMP1`, `E2-P2A-VIS1`, `E2-P2A-REPORT2`, `E2-P2A-REVIEW1` |
| C5 | TypeScript resolves `Promise`, `Math.max`, and `Math.min` with zero target-line diagnostics, while Anvien records gaps | The resolver workspace indexes only target `ScopeIR` definitions; TypeScript ambient/lib declarations are not loaded | Correct source facts become unresolved diagnostics; later graph-health labels are Go-oriented metadata, not the original cause | **CONFIRMED WRONG — bounded ambient resolver slice** | `E1-P1C-ORACLE1`, `E1-P1C-ANVIEN1`, `E2-P2B-ORACLE1`, `E2-P2B-IR1`, `E2-P2B-CAUSE-A1`, `E2-P2B-REVIEW1` |
| C6 | TypeScript follows both barrel aliases to the unique declaration; Anvien resolves file imports but emits neither consumer `CALLS` edge | `resolveImportedDef` searches only physical definitions in the barrel and does not follow its re-export binding; `TargetDef == nil` skips the consumer scope binding | Consumer calls fall through to unresolved/fallback resolution; direct-import control restores both calls | **CONFIRMED WRONG — bounded barrel-binding slice** | `E1-P1C-SRC2`, `E1-P1C-ANVIEN2`, `E1-P1C-ORACLE2`, `E2-P2B-RESOLVE1`, `E2-P2B-CAUSE-B1`, `E2-P2B-REVIEW1` |
| C7 | `context`/`impact` expose existing wrapper/barrel/file edges but do not show the two missing consumer calls; bounded process arrays are empty | Commands scan configured relationships already present in the graph; they do not re-resolve imports or synthesize calls | Downstream command output mirrors upstream incompleteness; no separate command-layer loss was measured | **BOUNDED VALID — present-edge traversal; upstream graph remains wrong** | `E1-P1E-CMD1`, `E1-P1E-CMD2`, `E1-P1E-CMD3`, `E1-P1E-CMP1`, `E2-P2C-CMD1`, `E2-P2C-SRC1`, `E2-P2C-TARGET1`, `E2-P2C-REVIEW2` |
| C8 | Seven canonical source-site facts retain one raw gap node, one edge, one Cypher row, and one `file-detail` sample | `file-detail` deliberately builds a ten-field unresolved DTO; DB CSV maps absent `Relationship.Step` to `0`; native tuple reads stringify scalars | Shape/nullability/type representation changes are observable, but selected cardinality remains one-to-one | **BOUNDED REPRESENTATION NON-PARITY; cardinality valid for selected facts** | `E1-P1D-RAW1`, `E1-P1D-PROJ1`, `E1-P1D-PROJ2`, `E1-P1D-CMP1`, `E1-P1D-REVIEW1`, `E2-P2C-PROJ1`, `E2-P2C-SRC2`, `E2-P2C-REVIEW2` |
| C9 | Actual process control `3,771/662/2,761/0/0`; +2-call control `3,773/662/2,761/0/0` | Process producer is heuristic/capped; this exact control did not alter surviving traces | No-change is a configuration-specific observation; it cannot prove completeness, monotonicity, sensitivity, or product meaning | **UNRESOLVED — bounded derived observation only** | `E2-P2C-DERIVED1`, `E2-P2C-DERIVED2`, `E2-P2C-SRC1`, `E2-P2C-REVIEW2` |
| C10 | Target/graph/artifact ownership remains stable across the restart | Direct target path and repo-local `.anvien` boundary are explicit; investigation artifacts stay in Anvien | Prevents target contamination from being mistaken for graph behavior | **CONFIRMED BOUNDARY; not an accuracy result** | `E0-P0A-R2-OWNER1`, `E0-P0A-R2-BOUNDARY1`, `E2-P2B-BOUNDARY1`, `E2-P2C-BOUNDARY1`, `E2-P2A-REVIEW1` |

## Layer-by-layer interpretation

### 1. Discovery is upstream of every later fact

C1 occurs before parsing and before a File node exists. It cannot be repaired or explained by `file-detail`, Cypher, resolver, or command projections. The eight measured paths are a bounded proof of false directory pruning, not a repository-wide prevalence count.

### 2. TypeScript identity has three independent losses

C2 is an extractor omission. C3 is an identity-injectivity failure after extraction; the ScopeIR probe proves that the parser/provider retained the four local facts before graph emission. C4 is a metadata propagation omission: declaration nodes survive, but the export fact is absent. These causes must not be collapsed into a single “parser bug.”

### 3. Resolver correctness depends on its declaration universe and module binding model

C5 is a missing language authority/input boundary. C6 is a non-transitive re-export binding boundary. The differential probe isolates C6 without changing source or production code: imports remain `10`, while import-use edges and resolved calls change `1 -> 3` and `37 -> 39`; ambient gaps remain. This separates parser/path success from resolver binding failure.

### 4. Downstream commands are not independent oracles for missing upstream facts

C7 is important for interpretation: exact traversal can be faithful while the graph is wrong. `context` and `impact` cannot recreate absent `CALLS`; an empty process array only proves no selected direct `STEP_IN_PROCESS` edge was supplied. It does not prove that no product process exists.

### 5. Projections can change shape without changing bounded fact cardinality

C8 distinguishes three representation boundaries: reduced `file-detail` DTO fields, `nil Step -> 0` in relationship CSV, and native scalar-to-string conversion. The seven selected facts remain one-to-one in the checked surfaces. Whether nullable/type-preserving output is a public contract is unresolved and requires owner authority.

### 6. Derived process output remains heuristic evidence

C9 is intentionally not promoted to a correctness claim. The exact control and five reruns show no output change for this graph/configuration. The producer filters calls, caps candidates/traces, deduplicates subsets, and emits only surviving traces; no accepted product execution model or completeness authority was supplied.

## Classification ledger

### Confirmed wrong, bounded

- C1 scanner pruning of legitimate nested directories.
- C2 array-binding extraction omission.
- C3 range-free graph identity collision and duplicate-node replacement.
- C4 lost TypeScript export visibility.
- C5 missing ambient/lib declaration workspace.
- C6 non-transitive barrel re-export binding.

### Bounded valid or preserving behavior

- Raw/Cypher/`file-detail` cardinality parity for the seven selected source-site facts (C8's shape changes are recorded separately).
- `context`/`impact` traversal of relationships that actually exist (C7).
- Unique-name control `readAdminCommercialConfig` and direct-import control behavior within the fixed cases.

### Unresolved or authority-blocked

- Global prevalence of any cause outside the measured files/paths.
- Complete TypeScript compiler semantics, ambient policy, namespace/static-member modeling, wildcard/multi-hop/cyclic re-exports, and all destructuring/export forms.
- Global `context`, `impact`, rename, community, process, and semantic completeness.
- Product meaning of Features, Architectural Boundaries, Processes, communities, and layer labels; no accepted SPEC/product authority was used for those claims.
- Public contract requirements for nullable `step` and JSON scalar types.
- Remediation design, regression acceptance, performance diagnosis, and any global accuracy percentage.

## Historical and evidence-integrity corrections

- The earlier context-contaminated checkpoint and its graph hash/counts are not mixed into this synthesis.
- The first P2-C Supervisor review rejected stale derived-process wording (`E2-P2C-REJECT1`). The corrected report, five reruns, and latest PASS all use `3,771/662/2,761/0/0` versus `3,773/662/2,761/0/0` and make no monotonicity/sensitivity claim.
- The plan-ledger audit found stale pending-status prose; the current R10 refresh and evidence IDs now distinguish historical snapshots from current P2-A/P2-B/P2-C PASS records. P3 synthesis remains a plan-level gate.

## Verification and non-actions

Freshly checked for this synthesis:

- target graph hash/count and target HEAD;
- TypeScript AST/graph diff and ScopeIR identity reruns;
- real-source resolver differential (`10/1/37/542` versus `10/3/39/540`);
- target process control and five prior reruns;
- seven-fact raw/projection parity rerun;
- direct source line traces and current Anvien file-detail/impact evidence;
- targeted existing tests: `go test ./internal/mcp ./internal/filecontext ./internal/lbugload`.

Not run because this plan contains no implementation change:

- full production build, remediation tests, `anvien detect-changes`, and commit;
- any target-source or target-graph edit;
- any global accuracy or performance benchmark.

## Synthesis decision

The causal record is internally coherent for the fixed bounded slices. It traces each confirmed discrepancy to a first Anvien boundary, distinguishes downstream faithful traversal from upstream loss, preserves representation changes and counts, and leaves unresolved product/global questions explicit. The synthesis is ready for an independent final Supervisor review, but it does not close graph correctness or authorize a fix.
