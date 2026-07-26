# P0 Investigation Report: Direct Target Boundary

Status: **BLOCKED**

## Scope

This report records the P0 safety and path-boundary check for the restarted graph investigation. It does not analyze a graph, accept prior findings, or substitute any repository path.

## Required Target

- User-specified target: `E:\cheapapp`
- Anvien workspace/report repository: `E:\Anvien`
- Forbidden substitution: `E:\cheapapp.org` (a distinct path)

## Evidence

1. `anvien --help`, `anvien analyze --help`, `anvien index --help`, and `anvien list --help` were read before any target analysis command.
2. `E:\Anvien\README.md:98-103` documents `anvien analyze .` as creating `<repo>/.anvien/` and registering the repository.
3. `E:\Anvien\README.md:422-457` documents the graph and projection storage under `<repo>/.anvien/`, with only the global registry under `~/.anvien/`.
4. `E:\Anvien\internal\repo\paths.go:8-34` defines `StorageDirName = ".anvien"` and derives graph/meta/lock paths from the analyzed repository path; no alternate output-root flag is present in the analyzed CLI help.
5. `E:\Anvien\internal\repo\path_policy.go:14-44` resolves and validates the exact absolute path supplied to analyze; it does not map `E:\cheapapp` to another directory.
6. A read-only filesystem check returned `TARGET_NOT_FOUND` for `E:\cheapapp`.
7. A read-only root listing found `E:\cheapapp.org`, but no `E:\cheapapp`; the name difference is material and no authorization exists to substitute it.
8. Independent red-team checks reproduced the same absence and path mismatch.
9. `E:\Anvien\internal\analyze\analyze.go:205-225,515-530` resolves the supplied path, derives repo-local storage paths, creates the storage/temp directories, and on `--force` removes/rebuilds repo-local graph output.
10. `E:\Anvien\internal\cli\command.go:174-187,293-309` exposes no storage-root flag and calls managed `.gitignore` maintenance for a Git target before analysis.
11. `E:\Anvien\internal\cli\analyze_postrun.go:26-44,71-83`, `internal\aicontext\aicontext.go:61-91`, and `internal\gitignore\managed.go:48-62` show post-analysis writes to target metadata/context files and `.gitignore` in the normal CLI path.
12. `E:\Anvien\internal\repo\settings.go:16-27` can create repo-local settings; `internal\repo\registry.go:57-68,118-125` derives and normalizes storage from the repository path rather than accepting an external storage root.

## Consequence

Running `anvien analyze E:\cheapapp --force` cannot produce the requested fresh graph because the exact target directory is absent. Running it against `E:\cheapapp.org` would violate the user's scope. Running it against a copied/checkpointed target would violate the no-copy boundary. Running analyze against a path that writes `.anvien`, `.gitignore`, metadata, `AGENTS.md`, `CLAUDE.md`, or generated context/skill files into the target would violate the read-only target requirement. The reviewed CLI/source provides no supported safe output mode.

## Decision

P0 is blocked pending the exact target path becoming available or an owner-approved correction to the target path. No graph command, source comparison, root-cause slice, or remediation has been started against a substitute repository.

## Artifacts

All artifacts for this check are in `E:\Anvien`; no file was created or changed in `E:\cheapapp` or `E:\cheapapp.org`.
