# P2-A Investigation Report: Scanner File-Omission Root Cause

Status: **CONFIRMED BOUNDED ROOT CAUSE**

## Claim

Trace the eight P1-A TS/JS File-node omissions to the earliest proven divergence in Anvien without designing or applying a fix.

## Source-to-Graph Symptom

At target HEAD `a869876ab6262dacde6cd5d432d099a91852a646`, an independent Git manifest contains `895` eligible TS/JS paths while the fresh graph contains `887` matching File nodes. The graph is missing exactly eight source paths and has zero graph-only TS/JS paths. The bounded source/graph report is `reports/Investigation/20260726_144919_p1a_file_inventory.md`.

## First Divergence

The first divergence is directory discovery, before any file candidate, parser fact, File node, or database row exists:

1. `internal/ignore/constants.go:5-78` defines global ignored directory names including `env` (`:21`), `target` (`:40`), and `logs` (`:59`).
2. `internal/ignore/constants.go:219-235` normalizes a relative path and returns ignored when **any path segment** matches that global map.
3. `internal/ignore/matcher.go:41-56` calls `ShouldIgnorePath` before evaluating repository `.gitignore`/`.anvienignore` rules; a hardcoded match returns immediately.
4. `internal/scanner/scan.go:55-81` calls `matcher.Ignored` for each directory and returns `filepath.SkipDir` when true.
5. Consequently, nested legitimate source directories named `target`, `env`, or `logs` are pruned before lines `85-90` can visit or append their files as candidates.

## Fresh Reproduction

The read-only probe `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2a-ignore-probe.go` loaded the real target ignore matcher and checked all eight paths. For each path, both `parentIgnored` and `fileIgnored` were `true`. The output is `p2a-ignore-probe.json` in the same Anvien `.tmp` directory.

The three affected subtree families are:

- `app/(account)/dashboard/settings/sessions/target/...`
- `lib/env/...`
- `modules/email/server/logs/...`

`git check-ignore -v` produced no rule for the paths or their parent directories, and `E:\cheapapp.org\.anvienignore` is absent. The exclusion therefore does not originate from target repository ignore rules.

## File-Detail and Impact Evidence

The Anvien self-graph was refreshed before graph-based owner/impact commands.

- `internal/ignore/constants.go` file-detail: 16 symbols, 9 inbound references, 2 linked tests, file risk `high`; related source/test files include `internal/ignore/matcher.go`, `internal/scanner/scan.go`, `internal/scanner/selection.go`, and matcher/gitignore tests.
- `ShouldIgnorePath` exact-symbol upstream impact: affected file count `1`, impacted symbol `Matcher.Ignored`, command-level risk `LOW`; the containing file still has `high` file risk.
- `internal/scanner/scan.go` file-detail: 56 symbols, 153 inbound references, 50 linked tests, file risk `high`.
- `WalkRepositoryPaths` exact-symbol upstream impact: risk `CRITICAL`, impacted count `5`, affected modules `4`, affected processes `39`. This is a blast-radius warning, not a prohibition against a future scoped fix.

Raw artifacts:

- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2a-ignore-file-detail.json`
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2a-scanner-file-detail.json`
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2a-impact-should-ignore.json`
- `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2a-impact-walk-repository.json`

## Causal Determination

**Confirmed wrong for the eight measured paths:** global path-segment pruning treats ordinary nested application directories named `target`, `env`, and `logs` as universally disposable. `filepath.SkipDir` removes the entire subtree before file candidacy, which directly explains all eight missing File nodes.

This is one bounded scanner cause. It does not prove global prevalence for every hardcoded directory name, and it does not explain extraction, resolution, projection, command, process, or semantic discrepancies.

## No Remediation

No production source or test was changed. No fix design is approved. A future remediation slice must preserve intended build/cache exclusions while avoiding false pruning of legitimate nested source directories, and it must treat the `CRITICAL` `WalkRepositoryPaths` blast radius as a scope warning requiring full scanner/analyzer regression.
