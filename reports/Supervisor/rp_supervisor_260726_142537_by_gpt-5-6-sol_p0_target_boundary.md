# Supervisor Report: P0 Direct Target Boundary

Verdict: PASS

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260726_142537_by_gpt-5-6-sol_p0_target_boundary.md`
- Review time: `2026-07-26 14:25:37 Asia/Bangkok`
- Reviewer: `gpt-5-6-sol`
- Repo/project: `E:\Anvien` with requested external target `E:\cheapapp`
- Scope reviewed: P0 target-path and Anvien-storage safety claim for the restart plan
- Claim reviewed: the restart cannot safely open graph analysis because the exact requested target is absent and the reviewed Anvien CLI stores analysis data repo-locally by default
- Authority used: current user scope, `AGENTS.md`, restart plan, Anvien CLI help, Anvien README/source, filesystem and registry metadata
- Related artifacts: `reports/Investigation/20260726_142153_p0_direct_target_boundary.md`

## Executive Summary

- Problem: the requested target path `E:\cheapapp` is not present, while the existing indexed repository is `E:\cheapapp.org`.
- Decision: PASS for the narrow blocker claim. The evidence is current, direct, and independently reproduced.
- Required outcome: do not start P1 graph analysis until the exact target path is available or the owner explicitly changes the target scope and storage boundary.

## Source-Level Clearance Notes

- `internal/repo/paths.go:8-34`: clear for the reviewed claim; `StoragePath(repoPath)` derives `<repo>\\.anvien` and no alternate output root is exposed in the reviewed help.
- `internal/repo/path_policy.go:14-44`: clear for the reviewed claim; the analyzer resolves the exact absolute path and rejects a missing path rather than mapping it to another repository.
- `README.md:98-103,422-457`: clear for the reviewed claim; documented graph/index storage is repo-local and the global registry is separate.

## Evidence Checked

Passed:

- Fresh PowerShell probe: `E:\cheapapp` does not exist; `E:\cheapapp.org` exists.
- Fresh registry read: `C:\Users\TAM NGUYEN\.anvien\registry.json` maps `cheapapp-accuracy-direct` to `E:\cheapapp.org` with storage `E:\cheapapp.org\\.anvien`.
- Fresh metadata read: `E:\cheapapp.org\\.anvien\\meta.json` reports `repoPath` `E:\cheapapp.org`.
- Fresh command help: no `--output-dir` or equivalent alternate graph-root flag appears in `anvien analyze --help`.
- Fresh negative flag probe: `anvien analyze --storage-path E:\\Anvien\\.tmp\\outside --help` returns `unknown flag: --storage-path`.
- Source inspection proves the normal analyze path can write repo-local `.anvien` storage, managed `.gitignore`, metadata, `AGENTS.md`, `CLAUDE.md`, and generated context/skill files (`internal/analyze/analyze.go:205-225,515-530`; `internal/cli/command.go:174-187`; `internal/cli/analyze_postrun.go:26-44,71-83`; `internal/aicontext/aicontext.go:61-91`; `internal/gitignore/managed.go:48-62`).
- Source inspection proves `ANVIEN_HOME` changes only the global registry location and registry entries normalize storage back to `StoragePath(repoPath)` (`internal/repo/paths.go:38-50`; `internal/repo/registry.go:57-68,118-125`).
- Independent red-team path check matches the main-agent result.
- Verification freshness: fresh/current; no target analysis command was run.

Failed:

- None for the narrow blocker claim.

Not run:

- `anvien analyze E:\cheapapp`: not run because the path is absent.
- Any graph/source comparison or root-cause slice: not run because P0 blocks them.
- `E:\cheapapp.org` analysis: deliberately not run or treated as a substitute.

## Invariant Closure

- affected invariant: target repository identity and read-only evidence boundary
- sibling surfaces checked: requested path, similarly named path, global registry entry, target-local meta, Anvien storage path source, CLI help
- residual unverified same-invariant surfaces: only an untracked alternate binary/configuration outside the reviewed source could differ; no such mode is evidenced in the current Anvien 1.2.8 CLI/source.

## Overall Evaluation

The evidence proves only that P0 is correctly blocked. It does not accept any prior graph finding, graph accuracy claim, remediation, or performance conclusion. The restart must remain paused until the target identity and safe analysis boundary are resolved by the owner.
