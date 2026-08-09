# Anvien Graph Accuracy Contract

Status: Draft campaign authority; implementation blocked until the plan-correction Supervisor gate passes

Authority: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-26-anvien-graph-accuracy-roadmap.md`

## Purpose

This contract governs the seven-child campaign whose outcome is a graph that represents the measured source facts more correctly and precisely than the baseline graph.

The problem report at `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien` supplies the bounded findings and target oracles. Its proposed architecture is DRAFT input, not implementation authority. The causal synthesis and bounded Supervisor PASS verify the measured findings; they do not prescribe the repair design.

## Runtime and command boundary

- Work is performed directly in the production codebase at the repository root.
- Commands keep their normal grammar, including `anvien analyze <path> --force`, `anvien query`, `anvien context`, `anvien impact`, and the existing API/MCP surfaces.
- The implementation and ordering used by analyze at the start of a slice are baseline evidence, not immutable architecture.
- A slice may change a pipeline component only when current source, file-detail, impact, and behavior evidence show that the change is needed for its acceptance target.
- Analyze remains a repeatedly invoked command. Each invocation must report success or a clear failure according to the behavior actually implemented and verified at its real boundary.
- Concrete source tree, command mode, output path, and storage behavior remain owned by current production source and the evidence of the owning slice.

## Canonical graph facts

The campaign must preserve these semantic distinctions regardless of the concrete representation selected after source inspection:

- A declaration is one source occurrence.
- A symbol is the language-level entity to which a declaration belongs.
- Same-name declarations in different lexical owners are distinct.
- Type, value, and namespace meaning must not be conflated when the language distinguishes them.
- A source range covers the declared construct; a selection range identifies the declaring token when the provider can supply it.
- A binding leaf, export entry, module request, resolution target, external outcome, and unresolved outcome remain distinguishable facts.

Child 01 must determine the concrete graph and ScopeIR representation from current source evidence before implementation. This contract does not require a particular node topology, hash function, or new file layout.

## Identity and range rules

- Identity inputs must be sufficient to distinguish the measured `time` and `now` declarations and other declarations with the same name in different lexical owners.
- Canonical inputs may include normalized repository-relative location, language, semantic name, meaning, lexical owner, declaration role, binding path, and source anchor when source evidence proves they are required.
- Absolute repository paths, traversal order, worker order, and random values must not make an otherwise identical analyze nondeterministic.
- Range and selection-range encoding, coordinate base, and interval semantics must be established from provider/source behavior and then preserved consistently.
- Declaration merging or overload sharing is allowed only when the provider supplies evidence that the declarations represent one semantic symbol.
- A range-only patch is insufficient if lexical owner or meaning is still lost.

## Mutation and integrity rules

- A distinct source occurrence must not disappear because another fact has the same computed identity.
- Conflicting canonical facts must not be silently replaced, skipped, or collapsed.
- Any intentional merge or enrichment must have an explicit, source-backed rule and must conserve the contributing occurrences and provenance required by downstream acceptance.
- Every emitted relationship must reference existing endpoints.
- Graph validation reports identity collision, occurrence loss, invalid range, and missing endpoint failures clearly.
- Child 01 changes only the owners proven by its P1-C/P1-D impact evidence. Relationship identity or a broad producer rewrite is not required unless that evidence demonstrates the need and the Child 01 plan is updated first.

## Persistence and reader rules

- Graph JSON and Ladybug must preserve the corrected identifiers, ranges, meanings, exports, resolution facts, and endpoints that their source records contain.
- Field parity is measured at record level; matching counts alone do not prove parity.
- No serializer or loader may silently discard a corrected record.
- Child 02 inventories readers from current source and changes only readers whose consumed fields are affected by corrected graph facts.
- A reader uses the explicit fields supplied by its real source contract. It must not reconstruct meaning from an opaque identifier when a corrected field is available.
- Repeated normal analyze and reader checks use the artifact locations and mechanisms implemented by the codebase under test. The campaign does not predetermine how those artifacts are written.
- An analyze or graph-bound command failure is surfaced as a failure at the normal command/runtime boundary; data from another artifact is not substituted as a successful result.

## Semantic ownership

| Child | Owned responsibility | Required terminal result |
|-------|----------------------|--------------------------|
| 01 | declaration/symbol identity, lexical owner, range inputs, occurrence conservation, proven collision owner | `time` and `now` represented `4/4`, deterministic and integrity-valid |
| 02 | corrected-field persistence and affected-reader consistency | Graph JSON/Ladybug parity, zero dropped corrected records, affected readers correct |
| 03 | TypeScript binding-pattern extraction | six bounded bindings and six downstream sites correct |
| 04 | TypeScript export semantics | 21 bounded direct exports correct |
| 05 | module export and re-export resolution | two bounded barrel calls resolve to the terminal symbols |
| 06 | ambient/external resolution outcomes and health projection | three bounded ambient sites have correct external or explicit capability outcomes |
| 07 | cross-surface acceptance only | all five target oracles, determinism, affected-surface parity, runtime, and measured performance pass |

Child 07 does not repair production code. A failed acceptance row returns to the Child that owns the failed invariant.

## File responsibility rule

- Every touched production or test file owns one primary semantic responsibility.
- A file may call many modules when every link serves that responsibility.
- New behavior is placed in its evidence-backed owner; broad `utils`, `common`, `helpers`, or `misc` owners are not created.
- An existing mixed file is split only when the active slice must touch the mixed responsibility and impact evidence justifies the split.
- Unrelated legacy refactoring is outside this campaign.

## Campaign acceptance

| Defect oracle | Baseline | Required result |
|---------------|---------:|----------------:|
| Same-name declaration identity | `2/4` | `4/4` |
| Binding-pattern declarations and downstream sites | `0/6` | `6/6` |
| Direct export facts | `0/21` | `21/21` |
| Barrel calls to terminal symbols | `0/2` | `2/2` |
| Ambient sites | `0/3` | `3/3` correct external/capability outcomes |

Campaign completion additionally requires deterministic repeated analyze results for the accepted canonical facts, zero unexplained occurrence loss, zero missing relationship endpoints, Graph JSON/Ladybug parity for corrected fields, correct behavior on every source-proven affected reader, full build and real-runtime validation, exact evidence, per-slice commits, and final Supervisor PASS.
