# Child 03 P3-B2 Catch Contexts — Coder Impact/Boundary Report

- Status: `BLOCKED_ON_ORCHESTRATION_PLAN_UPDATE`
- Handoff target: Orchestration main
- Role: Coder; no Supervisor acceptance claim
- Recorded: `2026-08-15 00:20:37 +07:00`
- Repository: `E:\Anvien`
- Branch / HEAD / `origin/master`: `master` / `01f160e6e28ad74c1f379ce5ea47e643a5a14652` / `01f160e6e28ad74c1f379ce5ea47e643a5a14652`
- Open slice: Child 03 `P3-B2` only

## Outcome

Fresh source and Anvien evidence select a three-file P3-B2 implementation boundary:

1. `internal/providers/tsjs/definitions.go` — catch-clause dispatch and catch-leaf Definition/OwnedDefID/local-Binding emission.
2. `internal/providers/tsjs/scopes.go` — one exact catch-clause `ScopeBlock` candidate so the catch parameter has a real catch-local lexical owner.
3. `internal/providers/tsjs/extract_test.go` — the focused real-`Extract` catch matrix and transition of the existing P3-B1 sibling assertion that currently expects catch declarations to remain absent.

The active plan currently says lexical scope construction is inspect-only. Because correct catch-local scope is impossible without the exact `scopes.go` change, the coder stopped before all production/test edits as required. Orchestration must update/authorize this exact boundary before implementation resumes.

No production file, test file, governance file, temporary artifact, stage, commit, push, final detect, or Supervisor action was created by this lane.

## Authority and complete reads

Read in full before impact work:

- `AGENTS.md`
- `.agents/skills/working-rules/SKILL.md`
- `.agents/skills/coder/SKILL.md`
- active graph-accuracy roadmap
- all four Child 03 ledgers: plan, evidence, benchmark, and actual status
- `internal/providers/tsjs/extract.go`
- `internal/providers/tsjs/nodes.go`
- `internal/providers/tsjs/definitions.go`
- `internal/providers/tsjs/scopes.go`
- `internal/providers/tsjs/types.go`
- `internal/providers/tsjs/references.go`
- `internal/providers/tsjs/binding_patterns.go`
- `internal/providers/tsjs/extract_test.go`
- `internal/providers/tsjs/scope_id_test.go`
- `internal/providers/tsjs/legacy_scope_captures_test.go`
- `internal/providers/tsjs/definition_position_inputs_test.go`
- `internal/scopeir/facts.go`
- `internal/scopeir/ir.go`
- `internal/scopeir/sort_keys.go`
- `internal/scopeir/scope_tree.go`

Current source has `scopeir.BindingContextCatch`, but the TSJS provider has no catch dispatch/emission branch and no catch scope candidate. Before this slice, the only focused catch assertion is in `TestExtractParameterBindingPatternsPreserveShadowingAndSiblingContexts`, where `caught` is required to emit no catch leaf or Variable Definition.

## Invariant Family Map

- Family: TSJS catch-clause declaration bindings and catch-local lexical ownership.
- Authority / SSOT: Child 03 P3-B2 plan goal plus the accepted P3-A binding-leaf contract and accepted P3-B/P3-B1 source behavior.
- Input surface: pinned Tree-sitter `catch_clause` AST.
- Primary path: `Extract` scope collection/build -> `walkKind` -> catch dispatch -> accepted `extractBindingPattern` -> BindingLeaf/diagnostic collection -> one Variable Definition + one catch-owned DefID + one local Binding per legal leaf.
- Lexical boundary: one `ScopeBlock` whose range is the `catch_clause`, parented by the existing deterministic scope builder. This range contains both the parameter and body, so body facts resolve within the same catch-local environment and nested function scopes remain children.
- Sibling surfaces to preserve: identifier variables, variable patterns, callable parameters, direct formal-parameter ownership, 15 type-only-label zero facts/diagnostics, TypeScript unparenthesized arrows, explicit `this`, JavaScript controls, imports, type inference, assignments, loops, references, graph projection/resolution, Vue/resolution outputs, and accepted golden artifacts.
- Forbidden fallback: no enclosing function/module ownership for catch bindings, no name-based rescue, no duplicate identifier branch, no type invention, no assignment/reference or exception-flow modeling.
- Locked/deferred: P3-B2A, P3-C/P3-C2, Pn slices, Child 04+, final detect, staging, commit, push, and Supervisor verdict.

## E3-P3B2-IMPACT1

### Fresh graph basis

- `anvien --help`: exit `0`.
- `anvien analyze --force`: exit `0`, duration `51.2s`.
- Scanned / parsed code / failed: `1,647 / 689 / 0`.
- Graph: `100,756` nodes / `140,552` relationships.
- Indexed/current commit: `01f160e6e28ad74c1f379ce5ea47e643a5a14652`; stale `false`.
- Graph artifact: `E:\Anvien\.anvien\graph.json`.

### Pinned grammar evidence

The pinned TypeScript `v0.23.2` and JavaScript `v0.25.0` `node-types.json` contracts both define `catch_clause.parameter` as optional and permit exactly `identifier`, `array_pattern`, or `object_pattern`; `body` is a required `statement_block`. TypeScript additionally permits an optional `type` field. Therefore:

- `catch {}` is a legal zero-binding case and must emit zero leaves/diagnostics;
- identifier and pattern catch parameters share one catch adapter;
- the adapter must ignore catch type inference and must not invent type facts.

### File-detail

All commands used `--repo E:\Anvien --json` on the fresh graph and returned current/non-stale data.

| File | Symbols | Related files | In / out / local | Unresolved | Linked tests | Relationship risk |
|---|---:|---:|---:|---:|---:|---|
| `internal/providers/tsjs/definitions.go` | 95 | 22 | 6 / 82 / 12 | 208 | 3 | HIGH |
| `internal/providers/tsjs/scopes.go` | 35 | 18 | 11 / 44 / 11 | 81 | 4 | HIGH |
| `internal/providers/tsjs/extract_test.go` | 345 | 24 | 13 / 167 / 101 | 0 | 3 | LOW |

HIGH here is a relationship/blast-radius scope warning. It requires surgical edits and preservation regression; it is not an edit prohibition after plan authorization.

### Upstream file impact

| File | Risk | Direct impacts | Affected files | Processes / flows |
|---|---|---:|---|---:|
| `definitions.go` | MEDIUM | 8 | `definitions.go` 7, preserve-only `types.go` 1 | 0 / 0 |
| `scopes.go` | MEDIUM | 8 | `scopes.go` 6, `definitions.go` 2 | 0 / 0 |
| `extract_test.go` | LOW | 0 | none | 0 / 0 |

### Exact canonical-UID symbol impact

Each invocation used `anvien impact symbol --uid <canonical UID> --repo E:\Anvien --direction upstream --json` and returned LOW risk, `0` impacted symbols, `0` affected modules, and `0` affected processes.

1. `collector.emitDefinitionKind`

   `Method:file38:internal/providers/tsjs/definitions.goname28:collector.emitDefinitionKindarity1:2occurrence73:def:internal/providers/tsjs/definitions.go#14:0:Method:emitDefinitionKindscope63:scope:internal/providers/tsjs/definitions.go#14:0-90:1:Functionowner0:`

2. `collector.collectScopeCandidateForKind`

   `Method:file33:internal/providers/tsjs/scopes.goname38:collector.collectScopeCandidateForKindarity1:2occurrence78:def:internal/providers/tsjs/scopes.go#33:0:Method:collectScopeCandidateForKindscope58:scope:internal/providers/tsjs/scopes.go#33:0-42:1:Functionowner0:`

3. `TestExtractParameterBindingPatternsPreserveShadowingAndSiblingContexts`

   `Function:file39:internal/providers/tsjs/extract_test.goname70:TestExtractParameterBindingPatternsPreserveShadowingAndSiblingContextsarity1:1occurrence129:def:internal/providers/tsjs/extract_test.go#854:0:Function:TestExtractParameterBindingPatternsPreserveShadowingAndSiblingContextsscope66:scope:internal/providers/tsjs/extract_test.go#854:0-936:1:Functionowner0:`

New private catch-only helpers and new focused test functions do not yet have graph UIDs. They must remain inside these three selected owners.

## Exact proposed post-plan editable boundary

### `definitions.go`

- Edit only `collector.emitDefinitionKind` to dispatch `catch_clause`.
- Add private catch-only helpers that:
  - return immediately for an absent optional `parameter`;
  - call the accepted walker exactly once with `BindingContextCatch`, `Construct=catch_clause`, and `Pattern=parameter`;
  - append returned leaves/diagnostics without changing collector/result plumbing;
  - create one `NodeVariable` Definition per leaf using the accepted leaf Range/SelectionRange;
  - attach exactly one OwnedDefID and one `BindingLocal` to the exact catch `ScopeBlock`;
  - emit no declared/return type and call no type-inference path.
- Preserve all existing variable and parameter branches byte-semantically.

### `scopes.go`

- Edit only `collector.collectScopeCandidateForKind`.
- Add `catch_clause -> addScopeCandidate(scopeir.ScopeBlock, node)`.
- Do not change sorting, parent selection, ID format, containment helpers, module/class/function behavior, or general block-scope modeling.

### `extract_test.go`

- After production code only, add focused real-`Extract` tests for:
  - TypeScript identifier catch;
  - supported nested object/array/default/rest paths and exact modifiers/provenance/ranges/selections;
  - optional `catch {}` zero leaves/diagnostics/definitions;
  - JavaScript catch preservation;
  - catch scope parent/range and call `InScope` boundary;
  - same-name outer/catch shadowing with distinct DefIDs and exact single ownership/binding;
  - zero duplicate/false sibling declarations.
- Update only the catch expectation inside `TestExtractParameterBindingPatternsPreserveShadowingAndSiblingContexts`: `caught` must become exactly one catch-context leaf/Definition/OwnedDefID/local Binding while assignment `written` remains zero and every P3-B1 assertion remains intact.

## Required plan update

The P3-B2 scope boundary must explicitly authorize:

- `internal/providers/tsjs/definitions.go` for exact catch dispatch/emission;
- `internal/providers/tsjs/scopes.go` for the one-symbol catch `ScopeBlock` candidate change;
- `internal/providers/tsjs/extract_test.go` as the focused real-`Extract` test owner and exact stale catch expectation transition.

The walker, ScopeIR contracts, collector/result plumbing, references, types, imports, variables, parameters, loops, and downstream graph/resolution code remain inspect-/preserve-only.

## Governance and Git boundary

The four Orchestration-owned files retained their opening hashes after all reads and impacts:

| Governance path | SHA-256 |
|---|---|
| roadmap | `50B86CABEFD0CAB05DECAE157F906A0D773B8B7E696BC3E7B4DFA42231254554` |
| Child 03 plan | `B82A5A6BC24B688DDB621880898ABA7039592E0710350403731E9CE3F542D23C` |
| Child 03 evidence | `B432D351EFE8E2170F4D4436DC0DB75D252BBDFC11B9A5D389A59072A1A1C7B1` |
| Child 03 actual status | `EA91F8C1DFDEF883773F3C48602B27E720F042B9204837DF84CA54D150E13AE2` |

No staged path exists. At this checkpoint, the only pre-existing modified paths are those exact four governance files; this report is the sole coder-created untracked path. No production/test hash changed.

## Validation state

- Production implementation: not started; blocked by plan boundary.
- Tests: not edited or run for candidate behavior.
- Full build: not run because no implementation candidate is authorized.
- Real Extract/ScopeIR catch boundary: not yet available.
- Product regression: not run.
- `E3-P3B2-SRC1`, `BUILD1`, `TEST1`, and `BOUNDARY1`: pending.
- `E3-P3B2-REVIEW1`, `DETECT1`, and `COMMIT1`: reserved; not run/claimed.
- Handoff state: precise blocker, not `READY_FOR_SUPERVISOR`.

## Resume condition

Resume only after Orchestration updates the active P3-B2 plan/ledgers and explicitly authorizes the exact three-file/three-existing-symbol boundary above. On resume, recheck governance hashes/current bytes and Git status before production code.
