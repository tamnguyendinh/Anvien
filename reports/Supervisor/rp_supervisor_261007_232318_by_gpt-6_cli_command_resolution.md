# Supervisor Report: Global Anvien CLI command resolution

Verdict: PASS

## Metadata

- Reviewer: gpt-6
- Repo: E:\Anvien
- Scope: Windows PowerShell global CLI registration and execution.
- Claim: `anvien` resolves and executes in a fresh PowerShell session from `G:\Restaurant_manager`.
- Authority: user troubleshooting request, README installation instructions, user prohibition on planner, and user confirmation that Anvien now works.
- Evidence: JSON companion with fresh command output, full-build output, runtime hash, and source references.

## Executive Summary

The initial failure occurred before Anvien executed: the canonical executable existed, but the global command was absent. The npm command shims and package junction now exist; the npm prefix was already on PATH. Fresh PowerShell resolves the shim and runs version 1.3.0, top-level help, and analyze help with exit code 0. The repository full-build script completed successfully, including runtime, Web/launcher, and forced Anvien index refresh. Outcome: accepted.

Global registration appeared during investigation; its creation is not attributed to this agent. No production code or plan was edited.

## Source-Level Clearance Notes

- `README.md:93,136`: global registration is an explicit manual installation step.
- `anvien/package.json:34-35`: executable mapping is correct and preserved.
- `scripts/full-build.ps1:42,53,69`: full build rebuilds runtime/launcher and refreshes the index. It does not perform global npm registration.
- Global npm shim and junction were read directly and resolve to `E:\Anvien\anvien\bin\anvien.exe`.

## Evidence Checked

Passed:

- Fresh PowerShell in the reported working directory: command discovery, version, help, analyze help; exit 0.
- Full build: exit 0. Complete output and current canonical binary hash are in the JSON companion.
- User independently confirmed commands are usable.
- Repository worktree was clean before this report was written.

Not run:

- Full `anvien analyze --force` in Restaurant_manager was not repeated after the user confirmed resolution; the accepted claim is command discovery/dispatch. The full build did successfully execute forced analysis of Anvien.
- Production tests, graph impact, and detect-changes: no production implementation was changed.

## Invariant Closure

- Affected invariant: global command discovery must resolve the current canonical CLI from arbitrary project directories.
- Sibling surfaces: PowerShell shim, npm global junction, persisted user PATH, top-level command dispatch, analyze dispatch, canonical binary.
- Residual unverified surfaces for this command-resolution claim: none.
- Existing PowerShell/IDE processes with independently modified PATH are outside this claim; the user's shell and a fresh shell both work.

## Overall Evaluation

Current source, global registration, fresh shell behavior, successful full build, and user confirmation support acceptance. The issue was an absent global installation step; no Anvien command-handler repair was needed.
