# Child 03 P3-C — Binding occurrence graph projection and lexical resolution

Created: 2026-08-15 14:09:08 +07:00  
Role: Coder implementation candidate; this report does not self-accept the slice.  
Assigned slice: P3-C, “Project and resolve binding occurrences in the current graph.”

## Handoff summary

The P3-A through P3-B2A `BindingLeafFact` records now cross the graph boundary as distinct definition occurrences. References are resolved through the real lexical scope chain only when the resolved definition belongs to the exact accepted binding-occurrence set. The graph emits existing `ACCESSES` relationships with `proofKind=scope-binding`, `targetRole=binding`, resolved source-site identity, and read/write step metadata. Orphan, duplicate, owner-drift, and local-binding-drift inputs fail before graph mutation.

Production changed only in `internal/resolution/resolve.go`. Graph JSON and Ladybug persistence production remained unchanged because their existing schemas already preserve all required node IDs, endpoints, source-site fields, proof metadata, and ranges. Two focused test files cover resolution and persisted parity.

## Invariant Family Map

| Item | P3-C boundary |
| --- | --- |
| Family | Accepted binding leaf -> exact variable definition -> lexical owner/local binding -> graph node/`DEFINES` -> lexical reference -> Graph JSON -> Ladybug `CodeRelation` |
| Authority / SSOT | Active Child 03 plan and ledgers, campaign roadmap, and `docs/contracts/graph-accuracy-contract.md`, read in full before implementation |
| Production owner | `internal/resolution/resolve.go` |
| Sibling surfaces checked | Accepted TypeScript extraction facts; workspace definitions/scopes/bindings; existing emitter/reference index; graph JSON; Ladybug CSV/load/schema; real installed CLI analyze/query boundary |
| Inspect-only production | `internal/resolution/indexes.go`, `internal/resolution/emit.go`, `internal/lbugload/csv.go`, `internal/lbugload/load.go`, and Ladybug native read/load path |
| Forbidden fallback | Global or same-name rescue, inferred missing binding, new export/module/ambient semantics, assignment target promoted to declaration, persistence row skipping/fallback |
| Success oracle | Every accepted leaf has one distinct Variable node and one `DEFINES`; five reads plus one assignment write reach exact lexical targets; inner/outer `value` endpoints differ; zero binding gaps/orphans/drift; one import remains one import; Graph JSON/Ladybug parity is exact |
| Residual unverified surfaces | none inside P3-C; real-target validation remains intentionally locked to P3-C2 and was not accessed |

## E3-P3C-IMPACT1 — Mandatory pre-edit evidence

Fresh graph command used both exclusions before any graph-based owner selection:

```text
anvien analyze . --force --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**"
```

Pre-edit graph basis:

- Files: 1,100 scanned, 622 parsed code, 0 failed.
- Graph: 80,102 nodes, 118,375 relationships.
- Index was current at the then-current assignment HEAD.

Exact file/symbol evidence:

- `internal/resolution/resolve.go`: CRITICAL file blast radius; 61 symbols, 69 direct impacts, 37 affected files, 1 linked flow.
- `ResolveBoundInto`: CRITICAL; 83 impacts, 24 modules, 33 processes.
- `resolveCall`: CRITICAL; 26 impacts, 7 modules, 33 processes.
- `resolveAccess`: CRITICAL; 26 impacts, 7 modules, 33 processes.
- Samples selected from the impact surface were the public orchestration entry `ResolveBoundInto` and its two direct call/access dispatch owners. The CRITICAL rating was treated as a blast-radius warning: the implementation stayed in the one proven file and preserved all non-binding branches.
- `indexes.go`, `emit.go`, graph schema, and persistence production were proven compatible and retained as inspect-only owners.

Baseline real probe before production change:

- All 15 fixture definitions and all 15 `DEFINES` relationships were projected.
- Five accepted binding receivers had zero reference endpoints.
- Those five sites produced five unresolved records.

Final full-build refresh, also with both exclusions, produced 1,102 scanned / 624 parsed code / 0 failed and 80,064 nodes / 118,992 relationships. The count change includes the final P3-C source/tests; it is validation evidence, not a product benchmark.

## E3-P3C-SRC1 — Production implementation

Changed production file and symbols:

- `internal/resolution/resolve.go:52` — `ResolveBoundInto`
- `internal/resolution/resolve.go:135` — `buildBindingOccurrenceIndex`
- `internal/resolution/resolve.go:211` — `bindingOccurrenceKeyForDefinition`
- `internal/resolution/resolve.go:224` — `bindingOccurrenceKeyForLeaf`
- `internal/resolution/resolve.go:237` — `bindingOccurrenceIndex.resolve`
- `internal/resolution/resolve.go:250` — `bindingOccurrenceReferenceSource`
- `internal/resolution/resolve.go:260` — `emitBindingOccurrenceReference`
- `internal/resolution/resolve.go:311` — `resolveCall`
- `internal/resolution/resolve.go:465` — `resolveAccess`

Behavioral invariant implemented:

1. Before any graph mutation, every accepted leaf must match exactly one Variable definition by normalized file path, name, range, and selection range.
2. That definition must have exactly one lexical owner and exactly one local `(scope, name, defID)` binding.
3. A definition occurrence may be projected only once; orphan, duplicate, owner drift, and binding drift return an error.
4. Reference lookup follows `resolveScopedName` from the fact’s real `InScope`, then accepts the result only if its definition ID is in the validated occurrence index. This is lexical resolution, not global/name rescue.
5. Member-call receivers emit read `ACCESSES`; unqualified access facts emit read/write `ACCESSES`. Existing member/import-member handling remains unchanged.
6. Emitted relationships preserve exact source-site ID/range, target definition ID, resolved status, `scope-binding` proof, `binding` target role, evidence, and read/write step.

Tracked production diff: 219 insertions, 11 deletions in the single proven production owner. `git diff --check -- internal/resolution/resolve.go` passed.

## Tests added after production behavior

- `internal/resolution/p3c_binding_occurrence_test.go:11` — `TestP3CBindingOccurrencesProjectAndResolveLexically`
- `internal/resolution/p3c_binding_occurrence_test.go:123` — `TestP3CBindingOccurrenceProjectionFailsClosedOnOrphanOrDrift`
  - failure subcases: missing exact definition, missing local binding, and duplicate accepted leaf
- `internal/lbugload/p3c_binding_occurrence_persistence_test.go:20` — `TestP3CBindingOccurrencesPreserveGraphJSONAndLadybugLoadParity`

The shadowing oracle uses occurrence-exact range/selection matching. It does not use the older shared test helper that collapses same-name, same-qualified-name variables to the first graph node.

## E3-P3C-BUILD1 — Holder-clean full build

Pre-build holder evidence:

- `anvien doctor locks --repo . --json`: analyze lock `free`, lock file absent.
- The global editor-owned `anvien.exe mcp` process at PID 9176 held the binary replaced by `npm install -g .`; it was terminated.
- Recheck found no live `anvien`, Go, npm, esbuild, or Vite build process and no analyze lock.

Executed command:

```text
powershell -ExecutionPolicy Bypass -File .\.tmp\p3c-full-build.ps1
```

The temporary repo-local driver reproduced `scripts/full-build.ps1` step-for-step, set `TEMP`/`TMP` under `E:\Anvien\.tmp`, and changed only the final command to:

```text
anvien analyze . --force --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**"
```

Result: PASS, exit 0, 127.1 seconds.

- npm install: PASS, 0 vulnerabilities.
- package runtime build/global install: PASS.
- `anvien version`: 1.2.8.
- launcher Go/native build: PASS.
- Vite production build: PASS, 2,943 modules transformed.
- final excluded-tree graph refresh: PASS, 1,102 scanned / 624 parsed code / 0 failed; 80,064 nodes / 118,992 relationships.

Non-failing output: npm reported its existing allow-scripts warning, and Vite reported existing dynamic-import/chunk-size warnings. Neither changed the exit result. The temporary build driver was removed after evidence capture.

This slice is not benchmarkable under the campaign definition; the build/analyze timings above are validation evidence only.

## E3-P3C-TEST1 — Final-byte tests and regressions

Production compile checkpoint before focused test authoring:

```text
go test -count=1 -run '^$' ./internal/resolution
```

Result: PASS.

Post-build focused tests:

```text
go test -count=1 -run 'TestP3C' -v ./internal/resolution ./internal/lbugload
```

Result: PASS.

- Projection/shadowing test: PASS.
- Fail-closed test and all three subtests: PASS.
- Graph JSON/Ladybug persistence parity test: PASS.

Post-build invariant-family regression:

```text
go test -count=1 ./internal/resolution ./internal/lbugload ./internal/graph ./internal/lbugschema ./internal/providers/tsjs ./internal/scopeir
```

Result: PASS for all six packages.

Post-build static check:

```text
go vet ./internal/resolution ./internal/lbugload ./internal/graph ./internal/lbugschema ./internal/providers/tsjs ./internal/scopeir
```

Result: PASS, no output.

All commands used repo-local `TEMP`/`TMP`. A broad `go test ./...` was intentionally not used because the explicit task prohibition prevents traversing the excluded skill trees; affected and nearest-boundary packages were run directly instead.

## E3-P3C-BOUNDARY1 — Real graph and persisted-graph validation

A repo-local debug fixture containing an import, outer/inner shadowed array bindings, assignment-form destructuring, catch binding, for-of declaration binding, and parameter binding was analyzed by the final installed CLI. Both source and generated index were removed after evidence capture.

Analyze command:

```text
anvien analyze .tmp/p3c-boundary-repo --force --skip-git --name p3c-boundary-20260815 --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**" --json
```

Observed final result: 2 scanned / 2 parsed / 0 failed; 14 graph nodes; 21 relationships; file projection unresolved count 0.

Graph JSON oracle at `.tmp/p3c-boundary-repo/.anvien/graph.json` before cleanup:

```text
VariableNodes=7
BindingAccesses=6
ReadAccesses=5
WriteAccesses=1
Imports=1
Uses=1
ResolutionGaps=0
MissingEndpoints=0
```

Occurrence and assignment conservation:

- Seven accepted binding leaves persisted as seven distinct Variable node IDs and seven `DEFINES` endpoints: outer `rows`, outer `value`, inner `rows`, inner `value`, `caught`, `loop`, and `arg`.
- The assignment target `for ([value] of rows)` at line 6 produced no declaration node. It persisted only one step-2 write relationship to the existing outer `value` definition at line 4.

Lexical endpoint/shadowing oracle:

- Outer line-5 `value` read -> outer line-4 `value` definition.
- Line-6 assignment write -> outer line-4 `value` definition.
- Inner line-9 `value` read -> distinct inner line-8 `value` definition.
- `caught`, `loop`, and `arg` each reached their own exact lexical definition.
- Every edge carried a unique source-site ID, `proofKind=scope-binding`, `targetRole=binding`, correct step, and existing source/target endpoints.

Ladybug persisted queries:

```text
anvien cypher "MATCH (n:Variable) RETURN n.id AS id, n.name AS name, n.startLine AS startLine, n.startCol AS startCol ORDER BY n.startLine, n.startCol" --repo p3c-boundary-20260815
```

Result: 7 rows, with distinct outer/inner `rows` and `value` occurrence IDs.

```text
anvien cypher "MATCH (a)-[r:CodeRelation]->(b:Variable) WHERE r.type = 'ACCESSES' AND r.targetRole = 'binding' RETURN a.id AS sourceId, b.id AS targetId, r.step AS step, r.sourceSiteId AS sourceSiteId, r.proofKind AS proofKind, r.targetText AS targetText, r.startLine AS startLine, r.startCol AS startCol ORDER BY r.startLine, r.startCol" --repo p3c-boundary-20260815
```

Result: 6 rows, exactly 5 reads and 1 write, with the lexical endpoints listed above.

Additional persisted counts:

- File -> File `IMPORTS`: 1.
- File -> Variable `DEFINES`: 7.
- `ResolutionGap`: 0.
- `resolution-inventory`: 6 resolved references; unresolved, source-backed unresolved, unattributed unresolved, and every resolution-gap bucket all 0.

The focused persistence test additionally deep-compares the six relationships before/after Graph JSON round-trip, verifies one persisted `DEFINES` per leaf, compares every Ladybug CSV relationship column, and requires zero skipped rows, fallback inserts, fallback failures, or warnings.

Import behavior remained at the accepted oracle: `ImportsResolved=1`, `ImportUsesEmitted=1`, and `FinalizedImportsEmitted=1`; the real graph contained one `IMPORTS` and one `USES`. Therefore P3-C caused zero import-count delta.

## Known failures and out-of-scope observations

- Final P3-C build, focused tests, regressions, static checks, Graph JSON boundary, Ladybug queries, and resolution inventory have no failures.
- During boundary-fixture shaping, an extra explicit `helper()` call produced one pre-existing `unresolved_call` analyzer gap even though the import relationship resolved. Imported-call semantics are preserve-only and explicitly outside P3-C, so no production or test change was made for it. The final scoped fixture retained the import but omitted that unrelated call and reached zero gaps.
- Vite/npm warnings from the full build are non-failing existing build output, classified above.

## Git and workspace boundary

Assignment HEAD: `69089d17b6a248b2a7c03f45bfa0cfec3c78d7e4`.  
Current HEAD at handoff: `0014123f45ce5489043fea68c198ce0bc6548905` on `master`.

The assignment HEAD is an ancestor of the current HEAD. At 2026-08-15 14:05:33 +07:00, another session advanced HEAD with the doc-only commit `0014123f docs(orchestration): refine monitoring and documentation rules`. This coder lane issued no commit command and did not author or amend that commit; production/test working bytes remained the validated P3-C candidate.

P3-C candidate hashes:

| File | Lines | SHA-256 |
| --- | ---: | --- |
| `internal/resolution/resolve.go` | 599 | `93f7b75b27c4f7ebd9061248df876195d497bf9e6a6b4852516bc45db8f1e512` |
| `internal/resolution/p3c_binding_occurrence_test.go` | 346 | `af148e3ceb48a8dc5e0c5a1115685d63286c2c1e0b6e230104e67216ede7cbae` |
| `internal/lbugload/p3c_binding_occurrence_persistence_test.go` | 282 | `c704b45fd350f2a1b064d79e78b4dc99f6378d44358b4e86149a09fa38d4a850` |

Final unstaged/untracked boundary, confirmed after report creation:

```text
 M docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md
 M docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-actual-status.md
 M docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-evidence.md
 M docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-03-typescript-binding-pattern-extraction/2026-07-28-03-typescript-binding-pattern-extraction-plan.md
 M internal/resolution/resolve.go
?? internal/lbugload/p3c_binding_occurrence_persistence_test.go
?? internal/resolution/p3c_binding_occurrence_test.go
?? reports/coder/rp_coder_260815_140908_by_gpt-5_p3c_binding_occurrence_projection.md
```

The four plan/roadmap/evidence/status paths were pre-existing Orchestration-owned modifications and received no write from this lane. No other production/test path appeared. Build output and all P3-C debug fixtures/scripts were cleaned without removing valid evidence.

Excluded-tree disclosure: before the context handoff, one broad filename-list command printed three path names from the forbidden skill trees. No excluded file content was opened or read, no excluded result was used as evidence, and nothing under either tree was modified. Subsequent analyze, search, and final Git boundary commands used explicit exclusions. It would be inaccurate to claim that no excluded pathname was ever surfaced; it is accurate that the trees were otherwise untouched.

No `detect-changes`, Supervisor, staging, commit, push, subagent, target-repository access, plan/ledger edit, or excluded-tree content access occurred in this lane.

## Next responsible person

Orchestration main should verify this candidate boundary and open a separate visible Supervisor lane. This coder report is a handoff, not acceptance.

READY_FOR_SUPERVISOR
