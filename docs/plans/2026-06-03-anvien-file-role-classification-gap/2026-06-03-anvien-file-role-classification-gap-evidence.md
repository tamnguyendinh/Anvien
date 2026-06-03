# Anvien File Role Classification Gap Evidence Ledger

Date: 2026-06-03

Status: Open

Companion files:

- Plan: [2026-06-03-anvien-file-role-classification-gap-plan.md](2026-06-03-anvien-file-role-classification-gap-plan.md)
- Benchmark ledger: [2026-06-03-anvien-file-role-classification-gap-benchmark.md](2026-06-03-anvien-file-role-classification-gap-benchmark.md)

## Evidence Rules

1. Record Anvien commands used to discover the owner and baseline.
2. Keep quantitative benchmark tables in the benchmark ledger.
3. For code changes, record impact/blast-radius before edits.
4. Preserve the distinction between raw unresolved, default-visible unresolved, and file-role classification.
5. Record failures and their handling.
6. Record `anvien detect-changes --repo Anvien --scope all` before each implementation commit.

## Evidence Template

Use this template for implementation phases:

```text
## E<n> - <Phase/Task>

Date:

Status:

Scope:

- ...

Source / command evidence:

| Check | Result |
|---|---|
| ... | ... |

Impact / blast radius:

| Target | Result |
|---|---|
| ... | ... |

Implementation evidence:

| File | Evidence |
|---|---|
| ... | ... |

Validation:

| Command | Result |
|---|---|
| ... | ... |

Failures / handling:

- ...

Detect changes:

| Command | Result |
|---|---|
| `anvien detect-changes --repo Anvien --scope all` | ... |

Commit:

- `<hash> <subject>`
```

## E0 - Baseline Analyze Refresh

Date: 2026-06-03

Status: completed

Command:

```powershell
.\anvien\bin\anvien.exe analyze --force
```

Result:

| Field | Value |
|---|---:|
| Files scanned | 821 |
| Parsed code files | 601 |
| Failed files | 0 |
| Indexed documents | 114 |
| Indexed metadata | 99 |
| Indexed analyzers | 0 |
| Indexed scripts | 4 |
| Indexed static | 3 |
| Unsupported language gaps | 0 |
| Unknown gaps | 0 |
| Graph nodes | 60974 |
| Graph relationships | 96624 |
| File projection files | 821 |
| File projection dependency edges | 15965 |
| Default-visible unresolved files | 336 |
| Raw unresolved files | 353 |
| Raw-only file difference | 17 |

Evidence interpretation:

- Analyze succeeded with no failed files.
- The `353 - 336 = 17` difference is not parse failure or unknown language inventory.
- The issue is classification expressiveness for known raw-only files.

## E1 - Raw-Only File Evidence

Date: 2026-06-03

Status: completed

Commands:

```powershell
.\anvien\bin\anvien.exe file-hotspots --repo Anvien --json --limit 0 --sort raw-unresolved --unresolved-only
.\anvien\bin\anvien.exe file-context internal/frameworks/frameworks.go --repo Anvien --json
.\anvien\bin\anvien.exe file-context internal/scopeir/sort_keys.go --repo Anvien --json
.\anvien\bin\anvien.exe file-context internal/group/types.go --repo Anvien --json
.\anvien\bin\anvien.exe file-context internal/repo/paths.go --repo Anvien --json
```

Raw-only file summary:

| Metric | Value |
|---|---:|
| Raw unresolved files | 353 |
| Default-visible unresolved files | 336 |
| Raw-only files | 17 |
| Raw-only source sites | 376 |
| Raw-only production source sites | 0 |
| Raw-only non-actionable source sites | 376 |

Representative samples:

| File | Raw | Classification | Sample targets |
|---|---:|---|---|
| `internal/frameworks/frameworks.go` | 209 | `builtin=2`, `standard_library=207` | `strings.ReplaceAll`, `strings.ToLower`, `strings.HasPrefix`, `float64` |
| `internal/scopeir/sort_keys.go` | 63 | `builtin=63` | `int`, `string` |
| `internal/group/types.go` | 16 | `builtin=16` | `int`, `map[string]string`, `map[string]RepoSnapshot` |
| `internal/repo/paths.go` | 13 | `standard_library=13` | `os.Getenv`, `os.UserHomeDir`, `filepath.Join` |

Evidence interpretation:

- The raw-only files are recognized source files.
- Their unresolved sites are all non-actionable builtin, standard library, or test-framework targets.
- The user-facing gap is that Anvien lacks a concise role label explaining these files as backend support/model/helper files.

## E2 - Owner Discovery

Date: 2026-06-03

Status: completed

Commands:

```powershell
.\anvien\bin\anvien.exe query "file classification appLayer functionalArea file role unresolved file projection" --repo Anvien
.\anvien\bin\anvien.exe context "ClassifyAppLayer" --repo Anvien
.\anvien\bin\anvien.exe context "ClassifyFunctionalArea" --repo Anvien
```

Owner evidence:

| Owner | Evidence |
|---|---|
| `internal/semantic/app_layer.go` | Query rank 1 for app-layer classification; `ClassifyAppLayer` owns existing app-layer path classification. |
| `internal/semantic/functional_area.go` | Query rank 2 for functional-area classification; `ClassifyFunctionalArea` owns existing functional-area path classification. |
| `internal/filecontext/context.go` | Query rank 2 file-layer owner; `FileSummary` owns file summary fields and unresolved buckets. |
| `anvien-web/src/components/FileMapPanel.tsx` | Query result and search evidence show file-summary Web consumer if role labels surface in UI. |
| `anvien-web/src/components/FileDetailPanel.tsx` | Query result and search evidence show file-detail consumer if role labels surface in UI. |

## E3 - Impact Baseline For Likely Classifier Edits

Date: 2026-06-03

Status: completed

Commands:

```powershell
.\anvien\bin\anvien.exe impact symbol "ClassifyAppLayer" --uid "Function:internal/semantic/app_layer.go:ClassifyAppLayer#1" --repo Anvien --direction upstream
.\anvien\bin\anvien.exe impact symbol "ClassifyFunctionalArea" --uid "Function:internal/semantic/functional_area.go:ClassifyFunctionalArea#1" --repo Anvien --direction upstream
```

Blast radius:

| Target | Risk | Impacted count | Affected files | Affected app layers | Affected functional areas |
|---|---|---:|---:|---|---|
| `ClassifyAppLayer` | CRITICAL | 6 | 4 | `backend=6` | `analyzer=1`, `cli=1`, `graph_health=1`, `unknown=3` |
| `ClassifyFunctionalArea` | CRITICAL | 4 | 4 | `backend=4` | `analyzer=1`, `cli=1`, `graph_health=1`, `unknown=1` |

Directly affected files from impact output:

- `internal/analyze/analyze.go`
- `internal/cli/command.go`
- `internal/graphaccuracy/access_candidate.go`
- `internal/semantic/app_layer.go`

Impact interpretation:

- Classifier edits are allowed but must be narrow.
- The plan should prefer additive file-role classification over broad rewrites of existing app-layer or functional-area behavior.

## E4 - Plan Creation

Date: 2026-06-03

Status: completed

Files created:

| File | Evidence |
|---|---|
| `docs/plans/2026-06-03-anvien-file-role-classification-gap/2026-06-03-anvien-file-role-classification-gap-plan.md` | Plan controller with goal, scope, requirements, phase checklist, and risk notes. |
| `docs/plans/2026-06-03-anvien-file-role-classification-gap/2026-06-03-anvien-file-role-classification-gap-evidence.md` | Evidence ledger seeded with baseline, raw-only file evidence, owner discovery, and impact baseline. |
| `docs/plans/2026-06-03-anvien-file-role-classification-gap/2026-06-03-anvien-file-role-classification-gap-benchmark.md` | Benchmark ledger seeded with raw/default unresolved and role-coverage baseline. |

Implementation evidence:

- No product code was edited in P0-A.
- Plan status is ready for implementation.
