# Supervisor Report: Child 06A A006-M1 Measurement Handoff

Verdict: PASS

## Metadata

- Report file: `rp_supervisor_260827_071409_by_gpt-5_child06a_a006_m1_handoff.md`
- Review time: `260827 071409 +07:00`
- Reviewer: `gpt-5`
- Repo/project: `E:\Anvien`
- Scope reviewed: A006-M1 recovery build, sequential Cheapapp/Restaurant packets, existing measurement report, and worktree boundary
- Claim reviewed: `A006_M1_DIRECT_CALLEE_ATTRIBUTION_READY`
- Authority used: `AGENTS.md`, Child 06A `plan-rules.md`, four active ledgers, `E2-P2A-A006ARCH1`, and the A006 Architect measurement contract
- Related artifact: `E:\Anvien\reports\Investigation\rp_child06a_a006_m1_direct_callee_attribution.md`

## Executive Summary

- The initial unsupported `AttemptId=A006-M1` failure remains preserved as measurement-support history.
- The one authorized recovery build used unchanged overlay/script inputs, `AttemptId=A006`, and a new output root. Provenance independently exposes schema/attempt/build/root exit `1/A006/0/0`, exact `2/4/3` overlay/candidate/native identities, and matching output hashes.
- Cheapapp and Restaurant Manager each launched once and exited `0`. Cheapapp ended before Restaurant started; the exact gap is `261,011,279,700 ns`, capture overlap is `0`.
- Both packets contain `30/30` operations, `17/17` children, ten ordered groups, exact group-to-D001 conservation, nonnegative residuals, unchanged non-timing benchmark semantics, and A003-identical Graph JSON/stdout/stderr.
- Decision: accept the packet as measurement-only Architect input. It changes no accepted A003 baseline, D001 streak, checklist, disposition, production/test/script bytes, or target source.

## Source-Level Clearance Notes

- Canonical production/tests/scripts: clear — tracked diff is empty for these surfaces; the only worktree change at review is the assigned measurement report.
- Overlay `resolve.go` / `types.go`: clear for measurement-only use — provenance matches exact expected replacement hashes and the A006 Architect limited instrumentation to these two repo-local overlay copies.
- Target repositories: clear — each process packet records identical target HEAD/status before and after; Anvien staged set remains empty.

## Evidence Checked

Passed:

- Current report read through EOF; corrected exact cross-target gap `261,011,279,700 ns`; scoped `git diff --check` PASS.
- Recovery provenance: build exit `0`, repository clean at build boundary, exact `2` overlay mappings, `4` candidate-source hashes, `3` native hashes, executable SHA-256 `0362FE211C2E072988EF61B662FC4F0D2160437F520F4525CDA59DDE52FB57A7`.
- Cheapapp: launch/exit `1/0`; `30/30`; `17/17`; ten groups sum exactly `3,048,999,300 ns` = D001; parent/child/residual `19,018,779,500 / 19,001,125,700 / 17,653,800 ns`; denominator `27,890/887`.
- Restaurant Manager: launch/exit `1/0`; `30/30`; `17/17`; ten groups sum exactly `8,297,125,100 ns` = D001; parent/child/residual `18,605,489,700 / 18,582,446,900 / 23,042,800 ns`; denominator `86,030/1,234`.
- Independent normalized non-timing comparison against each accepted A003 benchmark is exact on both targets. Graph JSON, stdout, and stderr SHA-256 values also exactly match the accepted A003 identities.
- Verification freshness: current raw artifacts and current worktree were read during this review.

Failed:

- None after the one-line report gap correction. The preserved initial builder input failure is historical support evidence and created no target launch/value.

Not run:

- No build, target analyze, graph analyze, tests, benchmark rerun, or profile rerun. The review consumed the immutable just-produced packet and did not repeat the measurement gate.

## Invariant Closure

- Affected invariant: measurement-only identity, non-overlapping direct-callee attribution, exact D001 conservation, target separation, and A003 output/semantic parity.
- Sibling surfaces checked: recovery provenance; both process packets; both benchmark packets; both Graph JSON/stdout/stderr outputs; repository/target HEAD-status boundaries; scoped report diff.
- Residual unverified same-invariant surfaces: none required before Architect consumption. Instrumented elapsed/resource values remain attribution-only and cannot be promoted.

## Overall Evaluation

PASS. `A006-M1` is valid measurement input for one fresh A006 Architect. Planner and Coder remain locked until that Architect returns a production direction, and D001 streak remains `2`.
