# Anvien-Go vs Anvien-main Analyze Benchmarks

Date: 2026-05-16

This report compares both engines on two repos:

| Scenario | Repo | Profile |
|---|---|---|
| Anvien self analyze | `E:\Anvien` | Go-heavy / Go runtime home repo |
| Website analyze | `E:\Website` | TypeScript-heavy / Next.js-style app repo |

Raw benchmark payloads are embedded in `reports\benchmark\2026-05-16-120908-anvien-self-analyze.json` under `scenarios.<name>.runs`.

## Headline Summary

| Scenario | anvien | anvien-main | Go speedup | Go nodes | main nodes | Go relationships | main relationships |
|---|---:|---:|---:|---:|---:|---:|---:|
| Anvien | 13,588.1 ms | 63,215.2 ms | 4.65x | 19,984 | 17,539 | 47,622 | 39,019 |
| Website | 20,763.6 ms | 70,428.1 ms | 3.39x | 26,081 | 18,607 | 48,163 | 34,055 |

## Commands

### Anvien / anvien

~~~powershell
.\anvien-launcher\server-bundle\anvien.exe analyze E:\Anvien --force [redacted removed argument] --no-stats --benchmark-json reports\benchmark\2026-05-16-120908-anvien-self-analyze.json --benchmark-label 2026-05-16-120908-anvien-self-analyze
~~~

### Anvien / anvien-main

~~~powershell
node E:\anvien-main\anvien\dist\cli\index.js analyze E:\Anvien --force [redacted removed argument] --no-stats --benchmark-json reports\benchmark\2026-05-16-120908-anvien-main-self-analyze.json --benchmark-label 2026-05-16-120908-anvien-main-on-anvien
~~~

### Website / anvien-main

~~~powershell
node E:\anvien-main\anvien\dist\cli\index.js analyze E:\Website --force [redacted removed argument] --no-stats --benchmark-json reports\benchmark\2026-05-16-120908-anvien-main-on-website.json --benchmark-label 2026-05-16-120908-anvien-main-on-website
~~~

### Website / anvien

~~~powershell
.\anvien-launcher\server-bundle\anvien.exe analyze E:\Website --force [redacted removed argument] --no-stats --benchmark-json reports\benchmark\2026-05-16-120908-anvien-on-website.json --benchmark-label 2026-05-16-120908-anvien-on-website
~~~

## Anvien Detail

| Metric | anvien | anvien-main | Difference |
|---|---:|---:|---:|
| Total analyze time | 13,588.1 ms | 63,215.2 ms | anvien is 4.65x faster |
| API graph nodes | 19,984 | 17,539 | Go +2445 |
| API graph relationships | 47,622 | 39,019 | Go +8603 |
| Files scanned | 676 | 678 | Go -2 |
| Files parsed / parseable | 520 | 520 | 0 |
| Parse phase | 5,586.7 ms | 12,003.0 ms | main is 2.15x slower |
| DB / Ladybug load | 4,902.5 ms | 17,539.5 ms | main is 3.58x slower |
| Resolved calls | 12,792 | 585 | Go +12207 |
| Resolved references | 26,541 | 1,089 | Go +25452 |

### Anvien Coverage Note

`anvien-main` reports Go scope extraction as unavailable for the Go files in `E:\Anvien`:

| Field | Value |
|---|---:|
| Go parseable files | 406 |
| Go ScopeIR AST reused files | 0 |
| Go no-hook files | 406 |
| Go ScopeIR coverage | 0% |
| Go scope resolution reference sites | 0 |

So this scenario is useful for wall-clock and graph-size comparison, but it is biased toward `anvien` for Go resolution quality.

## Website Detail

| Metric | anvien | anvien-main | Difference |
|---|---:|---:|---:|
| Total analyze time | 20,763.6 ms | 70,428.1 ms | anvien is 3.39x faster |
| API graph nodes | 26,081 | 18,607 | Go +7474 |
| API graph relationships | 48,163 | 34,055 | Go +14108 |
| Files scanned | 1,870 | 1,870 | Go 0 |
| Files parsed / parseable | 998 | 998 | 0 |
| Parse phase | 6,316.1 ms | 27,611.0 ms | main is 4.37x slower |
| DB / Ladybug load | 4,057.5 ms | 14,604.8 ms | main is 3.60x slower |
| Resolved calls | 11,303 | 6,037 | Go +5266 |
| Resolved references | 20,849 | 10,232 | Go +10617 |
| Resolved accesses counter | 3 | 755 | not graph-equivalent; see audit below |
| Resolved type references | 9,543 | 3,440 | Go +6103 |
| Graph `ACCESSES` edges | 3 | 3 | same final graph edge count |
| Graph `HAS_PROPERTY` edges | 3 | 3 | same final graph edge count |
| Graph `Property` nodes | 5,222 | 3 | Go emits many standalone TS property nodes |

### Website Coverage Note

`anvien-main` reports full ScopeIR AST reuse for the JavaScript and TypeScript files in `E:\Website`:

| Language | Parseable files | AST reused | ScopeIR coverage | Reference sites |
|---|---:|---:|---:|---:|
| JavaScript | 7 | 7 | 100% | 523 |
| TypeScript | 991 | 991 | 100% | 73572 |

This makes the Website scenario a better cross-runtime comparison for a TypeScript-heavy repo. On this repo, `anvien` is still 3.39x faster and emits a larger graph.

### Website ACCESSES Audit

The `Resolved accesses` row was audited because the `3` vs `755` number is not a final graph-edge delta.

| Check | anvien | anvien-main | Interpretation |
|---|---:|---:|---|
| Final graph `ACCESSES` edges | 3 | 3 | No Go-vs-main deficit in emitted `ACCESSES` edges for this benchmark. |
| Final graph `HAS_PROPERTY` edges | 3 | 3 | Both final graphs only owner-link three properties. |
| Final graph `Property` nodes | 5,222 | 3 | Go emits many standalone TS property nodes, mostly from object/type shapes. |
| Access resolution counter | 3 | 755 | Not directly comparable; main's value is `scopeResolutionResolvedAccesses`, not final `ACCESSES` edge count. |

Audit input: existing embedded `anvien-main` benchmark payload under `scenarios.websiteTypescriptHeavy.runs.anvienMain`, plus the current `E:\Website\.anvien\graph.json` after `anvien` analyze. `anvien-main` was not rerun.

Root cause classification:

- The benchmark report should not use `resolvedAccessesDeltaGoMinusMain=-752` as a graph-quality conclusion.
- The real follow-up is TS property/member semantics: `anvien` emits standalone `Property` nodes for many TypeScript object/type shapes, but only definitions with `OwnerID` enter the `ownerMembers` index.
- `resolveAccess` resolves through `resolveMember(... propertyLabels())`, so unowned object/type-literal properties cannot become `ACCESSES` edges.
- Code pointers: `internal/providers/tsjs/definitions.go:48`, `internal/resolution/indexes.go:147`, and `internal/resolution/resolve.go:200`.

Next action: create a focused TS property ownership/access semantics task if higher `ACCESSES` coverage is required. Measure it against final graph `ACCESSES`/`HAS_PROPERTY` facts, not against the `anvien-main` internal access counter.

## Worktree Note

`coder.md` was already untracked before this benchmark work and was not modified.
