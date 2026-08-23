# Supervisor Report: P6-D Graph-Health Projection Resubmission

Verdict: REJECT

## Metadata

- Report file: `rp_supervisor_260823_064551_by_gpt-5_p6d_graph_health_projection_resubmission.md`
- Review time: `2026-08-23 06:45:51 +07:00`
- Reviewer: `gpt-5`
- Repo/project: `E:\Anvien`, branch `master`
- Scope reviewed: exact repaired seven-path P6-D candidate on HEAD `0d9803777880d8ecca09cae90592d4728a231160`, parent `2acd357db32ac0c2bd3c3968a950c773acf3000d`
- Claim reviewed: whether the repaired graph-health projector closes the prior six-field nested-authority/outer-outcome finding and whether the complete P6-D slice can be accepted
- Authority used: current Owner delegation; `AGENTS.md`; full `working-rules`, `supervisor`, and `planner` skills; all four active Child 06 ledgers; sealed Main handoff; coder resubmission; prior independent P6-D report; accepted P6-C3 anchor and current source/Git/runtime/Anvien evidence
- Related artifacts: Main handoff SHA-256 `2C8A05B9318EBB68CAF4CF4F1F155E31AA8627EA3BF76CA65654EC183A60F675`; coder report SHA-256 `6E8AB2DBC96B43C16FDC7102FCA89D5096CD76D6B597D162C732FED533F72420`; prior independent report SHA-256 `D015CBD4847261DD50C9033DFA487657629910B296041744EF0BCAF5E108F2AE`; accepted P6-C3 report SHA-256 `DBA2111E7739AEB95E3E74CC29A5CA79F122908E34A8D5B75953D55CADA72E0F`
- Role boundary: independent visible Supervisor, review-only; no code/test/ledger repair, staging, commit, branch, stash, reset, checkout, target access, cleanup, network, install, or package script
- Next owner: Main
- Material correction: amended at `2026-08-23 09:26:53 +07:00` under direct Owner instruction. The original report incorrectly labeled raw `go build -buildvcs=false ./...` as the repository's canonical "full build." Source inspection now establishes that `scripts/full-build.ps1` is the canonical pipeline. No build or analyze command was rerun for this correction; only the existing report, retained E-only evidence, and repository source were read.
- Concurrent-state note: while this amendment was being sealed, Main advanced live `master` externally to HEAD `c28660118fc606ea3e19ad2f6cfb206a768f46f1` / parent `1838be3074a508d442481fc8720ce3afa6745404f`. `git diff 0d980377..c2866011 -- scripts/full-build.ps1 package.json` is empty, so the canonical-build correction applies to the reviewed anchor. The live seven-path worktree identity has meanwhile changed to `377,945` bytes / canonical SHA-256 `E4FF42DE8AA1941010AFFF5DD38B6789AE71168E3EBF9F59D8EECF735D5BD044`; this report remains a verdict on the exact historical candidate named in `Scope reviewed` and does not accept the newer bytes.

## Executive Summary

- Problem: the prior review proved that graph-health accepted a nested TypeScript authority object even when six identity groups drifted from the outer P6-C3 outcome. Full P6-D additionally requires a built current runtime, native Ladybug and built-reader parity, exact target `3/3`, an independent oracle, and target pre/post boundary evidence.
- Local repair decision: cleared. Current `internal/graphhealth/diagnostics.go:410-441` enforces exact equality for file path, file hash, all four range coordinates, site kind, requested name, and requested meaning. The sole test owner contains six isolated hostile mutations within a complete `26/26` hostile denominator. All six now fail closed to `unclassified/review`; positive missing/rejected authority states, legacy behavior, and generic Graph JSON/ResolutionGap/summary parity remain green.
- Full-slice decision: not accepted. The canonical repository full build was not run. `scripts/full-build.ps1:18-52` defines a pipeline of `npm install`, `npm run build`, `npm install -g .`, global runtime discovery/version checks, launcher build, and `anvien analyze . --force`. The review boundary prohibited network, install, package scripts, global inventory, and outside-E effects, so the correct gate state is `NOT RUN / BLOCKED`, not "full build exit 1." The raw Go compile diagnostics also failed and produced no current product runtime. Native Ladybug, built readers, target `3/3`, oracle, and target pre/post evidence therefore remain unrun or unproved.
- Analyze decision: unresolved. `anvien analyze . --force` is part of the canonical full-build pipeline at `scripts/full-build.ps1:50-52`. This review ran a separate graph-refresh invocation, not the canonical pipeline: its first invocation remained live beyond bounded waits and was terminated after its session handle was lost; exactly one retry completed. That retry proves current graph refresh only. It does not prove the canonical full-build analyze step, and the first invocation's abnormal liveness/performance was not benchmarked or root-caused.
- Procedural decision: not clean. Three internal lanes were opened before a later Main guard prohibited them and were interrupted without their output being read. Live E-only evidence proves no observable Git/candidate/quarantine/top-level-temp drift, but cannot prove that those lanes never accessed target/C/network or never wrote an ignored file. A later lock-inspection command also exposed host launcher metadata outside E in output; no such path was opened or followed, but this crossed the name/PID-only evidence boundary.
- Required outcome: keep P6-D unchecked, unstaged, and uncommitted. Main must preserve the local repair, route the committed P6-C3 concern, resolve dependency/runtime authority, complete every blocked product/target gate, and use a clean independent re-review before acceptance.

## Acceptance Gate Matrix

| Gate | Current result | Full P6-D effect |
|---|---|---|
| Six nested/outer equality groups | cleared in source and `6/6` isolated hostile rows | local prior blocker closed |
| Complete hostile/mapping/legacy matrix | `3 + 2 + 26 + 1 + 1`, all command exits `0` | local projector evidence cleared |
| Full `internal/graphhealth` regression | exit `0` | affected package regression cleared |
| Generic Graph JSON/ResolutionGap/summary parity | exit `0`; exact `3` standard/non-actionable and `0` in-repo/analyzer-gap | in-memory/generic boundary cleared only |
| Canonical repository full build (`scripts/full-build.ps1`) | `NOT RUN / BLOCKED`; prohibited install/network/package-script/global actions | required gate missing |
| Raw Go repository compile diagnostic | `go build -buildvcs=false ./...` exit `1` after `6,600 ms`; this is not the canonical full build | diagnostic failure only; cannot satisfy full-build gate |
| Affected production build | exit `0` after `13,759 ms` | local compile cleared |
| Raw current-CLI build diagnostic | exit `1` after `283 ms`; output binary absent | current-runtime gate missing |
| Canonical full-build analyze step | not run inside the canonical pipeline | required gate missing |
| Separate graph-refresh analyze | first invocation had unresolved liveness; one allowed retry exited `0` | graph freshness only; not full-build/runtime proof |
| Native Ladybug / built readers | not run because no current runtime exists | required gate missing |
| Target `Promise`, `Math.max`, `Math.min` | `0/3` runtime proof; target not accessed | required gate missing |
| Independent oracle | not run | required gate missing |
| Target pre/post boundary | not established | required gate missing |
| Procedural isolation | two incidents, neither provably clean | acceptance integrity gate missing |

## Blocking Findings

### [CRITICAL] Canonical full-build, built-current-runtime, and target acceptance chain is absent

File: `docs/plans/2026-07-26-anvien-graph-accuracy-multi-plan/2026-07-28-06-ambient-external-resolution-and-diagnostics/2026-07-28-06-ambient-external-resolution-and-diagnostics-plan.md:448`

Issue: full P6-D requires the repository-defined full-build pipeline, current product-runtime truth, native persistence/reader truth, exact target outcomes, oracle comparison, and target pre/post preservation. The canonical full build was not run. The review instead mislabeled a raw Go compile diagnostic as "full build." The raw diagnostic could not build the current CLI, and the target remained locked and unaccessed.

Evidence:

- Canonical source truth: root `package.json` maps `full-build` to `scripts/full-build.ps1`. That script executes `npm install` (`18-20`), `npm run build` (`22-24`), `npm install -g .` (`26-28`), `Get-Command anvien` and `anvien version` (`30-35`), launcher build (`42-44`), another runtime version check (`46-48`), and `anvien analyze . --force` (`50-52`).
- Boundary truth: the Owner delegation prohibited network, install, package scripts, global executable inventory, and outside-E access/effects. Therefore Supervisor had no authority to execute the canonical script. Correct handling was to inspect it and record the full-build gate as `NOT RUN / BLOCKED`; relabeling a narrower command was invalid.
- Rule-13 preflight found `E:\Anvien\.anvien\analyze.lock` absent and `0` relevant Go/compiler/linker process names before the narrower Go diagnostics; no process required termination.
- Raw Go diagnostic only: with fresh E-only `HOME`, `USERPROFILE`, `TEMP`, `GOCACHE`, `GOMODCACHE`, and `GOPATH`, plus `GOPROXY=off`, `GOSUMDB=off`, and `GOTOOLCHAIN=local`, `go build -buildvcs=false ./...` exited `1` after `6,600 ms`. It did not invoke `scripts/full-build.ps1` or `anvien analyze` and must not be represented as the repository full build.
- Exact retained dependency evidence from that raw diagnostic is `17` zero-byte module lock files: `github.com/spf13/cobra@v1.10.2`; `github.com/knights-analytics/hugot@v0.7.2`; `github.com/tree-sitter/go-tree-sitter@v0.25.0`; Dart `github.com/UserNobody14/tree-sitter-dart@v0.0.0-20260508020638-507c5546dc73`; Swift `github.com/flamingoosesoftwareinc/tree-sitter-swift@v0.0.0-20260212012612-56ffc4e2dcc9`; and grammar modules for Kotlin `v1.1.0`, C# `v0.23.5`, C `v0.24.2`, C++ `v0.23.4`, Go `v0.25.0`, Java `v0.23.5`, JavaScript `v0.25.0`, PHP `v0.24.2`, Python `v0.25.0`, Ruby `v0.23.1`, Rust `v0.24.2`, and TypeScript `v0.23.2`. `GOPROXY=off` prevented resolution; no module content was downloaded.
- Exact source-confirmed fixture boundaries reached by raw `./...`: `anvien/test/fixtures/lang-resolution/go-method-enrichment` contains packages `animal` (`animal.go`) and `main` (`app.go`) in one directory; `anvien/test/fixtures/sample-code` contains `simple.go` (`package main`) beside standalone `simple.c`. These are raw Go package-discovery failures, not evidence that the P6-D graph-health production package failed; `./internal/graphhealth` built successfully.
- The original raw terminal stderr was not preserved as a separate durable artifact, so this report does not present reconstructed text as a verbatim transcript. The module lock identities and fixture layouts above are the strongest retained E-only attribution.
- After a second clean preflight, `go build -buildvcs=false ./internal/graphhealth` exited `0` after `13,759 ms`.
- A separate current product build, `go build -buildvcs=false -o <fresh-review-root>\runtime\anvien.exe ./cmd/anvien`, exited `1` after `283 ms` on the same offline-unavailable modules. The runtime output contains `0` files / `0` bytes.
- The Go messages saying `downloading` did not establish network use: `GOPROXY=off` rejected every lookup, and the fresh `gomodcache` contains only `17` zero-byte lock files with no downloaded module content.
- Existing packaged/global binaries and all quarantined caches were excluded as stale or unauthorized evidence.
- With no current runtime, native Ladybug, CLI/MCP/HTTP built-reader parity, target pre-state/analyze/query/health/file-detail, target `3/3`, oracle, and post-state were not run.

Why this blocks acceptance: plan lines `477-490`, the active ledger gates `E6-P6D-BUILD1`, `PARITY1`, `TARGET1`, `ORACLE1`, and `BOUNDARY1`, and the Owner delegation require these exact gates. The canonical full build was never exercised, and a truthful `BLOCKED` handoff does not satisfy it.

Fix direction:

1. Main/Owner must first resolve the authority conflict for the canonical script. Either explicitly authorize and contain every `scripts/full-build.ps1` effect (`npm install`, package build, global install/discovery, launcher build, and final analyze), or provide a repository-owned hermetic full-build entrypoint that performs the same product steps using only approved E-only inputs/outputs. Supervisor must not silently replace the canonical pipeline.
2. Provision exact dependencies from an explicitly authorized source into a new non-quarantined E-only dependency/cache boundary. A zero-byte lock-only cache is insufficient. Network fetch, reuse of protected caches, and stale-runtime substitution remain forbidden unless Owner explicitly changes those boundaries.
3. If raw `go build ./...` remains an intended sub-gate of the canonical contract, Main/Coder must make fixture discovery build-safe under approved plan authority: separate the `animal` and `main` packages; isolate the standalone C sample from the Go package; or establish an authority-approved nested-module/`testdata`/explicit-package boundary. Merely ignoring these errors cannot count as a passing full build. This is outside Supervisor repair authority and may be outside P6-D ownership, so Main must route it rather than edit opportunistically.
4. Run the actual canonical full build from current bytes and retain step-by-step command, exit, elapsed-time, runtime identity, and final analyze evidence. Only after it exits `0` may the built current runtime be used for native Ladybug, built readers, target `3/3`, oracle, and pre/post gates.

Re-review evidence required: canonical `scripts/full-build.ps1` (or Owner-approved equivalent) exit `0` with every step recorded; successful current runtime identity; final analyze completion and elapsed/liveness disposition; native Ladybug and built-reader records; exact three source-site outcomes/proofs with `0/3` analyzer gaps; independent TypeScript oracle; target HEAD/worktree/source/index/graph pre/post inventory.

### [HIGH] Canonical analyze gate has unresolved liveness/performance evidence

File: `scripts/full-build.ps1:50`

Issue: `anvien analyze . --force` is part of full build. The review failed to identify that linkage before judging. A separately invoked graph refresh remained live beyond bounded waits, lost its session handle, and required exact-PID termination. One allowed retry completed, but no benchmark artifact or accepted elapsed-time baseline was captured. The Owner reports the invocation took roughly two to three times normal; current durable evidence cannot confirm the ratio or attribute it to source, cold-cache state, runtime identity, lock/session handling, or another cause.

Evidence:

- Canonical script lines `50-52` make analyze a required full-build step.
- Initial standalone invocation: PID `17176` remained alive through bounded intervals, was terminated, and was explicitly excluded from graph evidence.
- Retry: exit `0` at `2,071` scanned / `763` parsed / `0` failed / `117,906` nodes / `164,481` relationships / `20,310` dependency edges / `469` unresolved projections.
- Missing evidence: no canonical full-build execution, no newly built current runtime, no retained initial session output, no `--benchmark-json` artifact, and no accepted same-condition baseline. Therefore neither "normal" nor "codebase regression" is proved.

Why this blocks acceptance: a successful retry proves graph freshness, not that the full-build analyze stage is reliable or performs within its accepted envelope. The symptom cannot be normalized away, and it cannot honestly be assigned to P6-D code without controlled evidence.

Fix direction: after producing the current runtime through the authorized canonical pipeline, Main must rerun the canonical analyze stage with a retained session handle and record wall time, lock/PID lifecycle, runtime identity, repository identity, and complete output. Run an E-only `--benchmark-json` measurement under the same approved runtime/cache conditions and compare it with an accepted baseline. If the regression reproduces, open and route a separate performance/liveness repair slice; Supervisor must not repair it inside P6-D.

Re-review evidence required: canonical full-build analyze exit `0`, retained end-to-end output, current-runtime identity, benchmark JSON, comparable baseline, and a documented disposition of the first-invocation liveness symptom.

### [HIGH / PROCEDURAL] Pre-guard internal-lane activity cannot be proven clean

File: N/A

Issue: before Main supplied the explicit no-subagent guard, this visible task opened three read-only internal lanes named `ledger_audit`, `history_audit`, and `source_test_audit`. On receipt of the guard, all three were interrupted immediately. Their output was never read or used. However, available task metadata exposes only final state, not a complete auditable command history.

Evidence:

- The guard arrived before creation of the review temp root at `2026-08-23 06:21:44 +07:00`; the exact guard-message receipt timestamp is not exposed by current task evidence.
- Each interrupt returned prior state `running`; a later live task snapshot reported all three as `interrupted` and only `/root` as `running`.
- The first post-guard Git/identity snapshot exactly matched Main's pre-dispatch truth: status `61 = 5` tracked + `56` untracked, index `0`, diff-check `0`, reports `55 = 44 Investigation + 7 Supervisor + 4 coder`, and exact candidate `371,376` bytes / SHA-256 `758066DE6FDF37DCCE1EEB22DD4B2B8F577CE8A58CA739445F9AFE6F72B0D3BE`.
- All eight protected temp/quarantine roots matched Main's per-root file/byte/zero-byte counts and had `0` reparse points. A bounded E-only top-level `.tmp` creation-time check around the incident found only this review's later root.
- These facts prove no observable tracked/untracked candidate, known quarantine-root, or top-level temp-root mutation. They do not prove absence of target/C/network reads, ignored-file writes inside an existing root, or other non-persistent activity by the interrupted lanes.

Why this blocks acceptance: the Owner requires independent review to occur entirely in this visible task and explicitly forbids inferring clean when the lane activity cannot be fully proved. Missing negative proof cannot support acceptance.

Fix direction: do not reopen or inspect the interrupted lanes. For any future acceptance attempt, use a fresh visible Supervisor task with no internal/separate lanes from its first action and remeasure the candidate before work.

Re-review evidence required: a clean task-local audit trail showing no spawned lanes and no unexplained Git/filesystem delta.

### [HIGH / PROCEDURAL] Analyze-lock inspection exposed forbidden launcher metadata

File: `E:\Anvien\.anvien\analyze.lock`

Issue: the first fresh analyze invocation exceeded the tool yield window and its session handle was not retained. To identify its verified holder, the review read the E-only lock file; that output included a `command` field containing host launcher metadata outside E. The external path itself was not opened, followed, listed, statted, or used, but emitting it crossed the explicit name/PID-only evidence boundary.

Evidence:

- PID `17176`, process name `anvien`, was verified as the live holder. After bounded waits it remained alive and was terminated by exact PID; the invocation was excluded from evidence. No lock file was manually deleted.
- Per Main's guard, exactly one fresh analyze retry was started with a retained session handle. It acquired a new lock and completed exit `0` at `2,071` scanned / `763` parsed / `0` failed / `117,906` nodes / `164,481` relationships / `20,310` dependency edges / `469` unresolved projections.
- The successful retry fully supersedes graph state, but it does not erase the procedural output-boundary incident.

Why this blocks acceptance: the review boundary allowed process preflight by necessary name/PID only and prohibited executable-path/command-line enumeration outside E. Acceptance cannot call this review procedurally clean.

Fix direction: future review must retain every long-running session handle and inspect only lock existence plus verified PID/name. Do not read or print the lock's command metadata.

Re-review evidence required: a clean current graph sequence without path/command-line metadata exposure.

## Source-Level Clearance Notes

- `internal/graphhealth/diagnostics.go:224`: clear locally. Structured carriers are evaluated before prefilled or heuristic classification; structured-present invalid input cannot fall through to name inference.
- `internal/graphhealth/diagnostics.go:303`: clear locally. Repository unresolved maps mechanically to `in_repo_unresolved/analyzer_gap`; TypeScript standard-library unresolved/capability maps to `standard_library/non_actionable`; all other structured diagnostic states fail closed.
- `internal/graphhealth/diagnostics.go:327`: clear locally. Malformed-present JSON is distinguished from absence, preserving legacy fallback only for diagnostics without a structured marker.
- `internal/graphhealth/diagnostics.go:352`: clear locally. Flat diagnostic/outcome identity, lifecycle, proof, target/reason, and authority-presence checks remain intact.
- `internal/graphhealth/diagnostics.go:410`: clear for the prior P6-D finding. Lines `419-427` enforce the six required nested/outer equality groups exactly.
- `internal/graphhealth/p6d_outcome_projection_test.go:13`: clear locally. The complete owner was read, not inferred from test names: policy `3/3`, valid missing/rejected `2/2`, hostile `26/26`, legacy `1/1`, and parity `1/1` are explicit.
- `internal/graphhealth/p6d_outcome_projection_test.go:135`: clear locally. Six isolated rows change only nested file path, file hash, one complete-range member group, site kind, requested name, or requested meaning while the rest remains valid.
- `internal/graphhealth/policy.go` and `internal/graphhealth/resolution_gap_inputs.go`: preserve-only. Current source and fresh impact show no alternate adapter that bypasses normalized Diagnostic projection; no edit exists.
- Target-name check: `rg` over the production/test owners returned no `Promise`, `Math.max`, `Math.min`, or target-repository literal. Generic names only are used.
- `internal/resolution/outcome.go:195`: routed concern, not edited or reopened. The committed P6-C3 validator compares nested authority to outer outcome only on source-site ID, stage, and status family. Lines `352-368` construct the outer fields from the authority, but the validator still lacks the six equality checks. Fresh `validateResolutionOutcome` impact is CRITICAL: `12` impacted symbols / `6` files / `2` direct / `2` modules / `33` processes. Main must route this out-of-authority concern; it is not a license for P6-D to modify closed P6-C3.
- Four Child 06 ledgers and coder report: identity-clear and semantically truthful about the local repair versus the full-slice blockers. Supervisor did not update them.

## Evidence Checked

Passed / cleared:

- Zero-trust authority seals: all four supplied reports matched bytes, LF/CR, strict UTF-8/no-BOM, and SHA-256 expectations.
- Exact repaired candidate, both before and after validation: `371,376` bytes, canonical SHA-256 `758066DE6FDF37DCCE1EEB22DD4B2B8F577CE8A58CA739445F9AFE6F72B0D3BE`; every per-path identity matched.
- Git before and after validation: branch `master`; HEAD/parent unchanged; status `61`; tracked/untracked `5/56`; index `0`; diff-check `0`; sole non-report untracked path is the P6-D test.
- Affected build exit `0`.
- Hostile test exit `0`, all `26/26`; focused P6-D exit `0`; full `internal/graphhealth` exit `0`; exact generic parity exit `0`.
- Fresh current Anvien retry exit `0`; `diagnostics.go` file-detail HIGH at `154` symbols / `48` inbound / `70` outbound / `149` local / `187` unresolved / `1` flow / `19` tests.
- `validStructuredResolutionAuthority` impact LOW at `3` impacted / `1` file / `1` module / `0` processes. The containing-file result remains CRITICAL at `107` impacted / `31` files / `34` direct / `1` flow / `19` tests; neither layer is normalized away.
- Detect exits `0`: file layer HIGH with `5` tracked files / `201` changed symbols (`180` backend graph-health + `21` docs); separate summary risk `low`, `affected_count=0`; ResolutionGap `96` entities / `100` occurrences (`38` analyzer-gap / `58` non-actionable; `14` builtin / `38` in-repository / `44` standard-library); Resolution Health `0`; semantic fields complete `117,906/117,906`.
- All eight pre-existing protected roots remained byte/count-identical and at `0` reparse points after validation.

Failed / blocked:

- Canonical repository full build: `NOT RUN / BLOCKED`; the prior "exit 1" label is retracted.
- Raw Go repository compile diagnostic exit `1`; not a substitute for canonical full build.
- Raw current-CLI build diagnostic exit `1`; runtime output absent.
- Canonical analyze stage not run; separate first analyze invocation had unresolved liveness/performance and its retry cannot substitute for full-build proof.
- Procedural isolation is not provably clean because of the pre-guard lanes and lock-metadata exposure.

Not run:

- Canonical `scripts/full-build.ps1`: prohibited install/network/package-script/global-discovery effects exceeded Supervisor authority. Source was inspected after Owner correction; the script was not executed.
- Native Ladybug and built CLI/MCP/HTTP/Web readers: no current product runtime exists; stale/quarantined substitutes forbidden.
- Target pre-state, analyze, exact site queries, graph-health/file-detail, `3/3`, and post-state: target remained locked; it was not accessed to strengthen a negative claim.
- Independent TypeScript oracle against the current target environment: blocked behind the same runtime/target authority gate.
- Network retry, install, package scripts, dependency fetch, shared/quarantined cache reuse, and cleanup workaround: explicitly forbidden.
- UI/Playwright: fresh impact shows no frontend owner; UI is N/A for this slice.

## Anvien Blast Radius

- Current graph: `2,071/763/0`, `117,906` nodes, `164,481` relationships, `20,310` dependency edges, `469` unresolved projections; indexed/current commit matches HEAD and file-detail reports `stale=false`.
- `diagnostics.go`: file detail HIGH; containing-file impact CRITICAL `107/31/34/1/19` (`impacted symbols / files / direct / flows / linked tests`).
- `validStructuredResolutionAuthority`: LOW `3/1/1/0` (`impacted symbols / files / modules / processes`). This narrow result does not downgrade the containing-file warning.
- Committed `validateResolutionOutcome`: CRITICAL `12` impacted / `6` files / `2` direct / `2` modules / `33` processes.
- Detect: file layer HIGH at `5` files / `201` changed symbols; independent summary low at `affected_count=0`. Both layers are preserved.

## Invariant Closure

- Affected invariant: a graph-health diagnostic may project external/capability policy only from one internally consistent P6-C3 outcome whose nested P6-B authority belongs to exactly the same source identity and accepted status family.
- Local P6-D sibling surfaces checked: structured marker detection; flat diagnostic/outcome parity; all six authority/outcome identity groups; status/reason/catalog-proof matrix; prefilled-policy override; legacy no-marker fallback; generic Graph JSON; `SourceBackedResolutionGapInputs`; ResolutionGap node/relationship; second Graph JSON; summary; target-name absence.
- Local prior finding: closed on current bytes.
- Residual unverified same-invariant surfaces: native Ladybug and built readers; exact target outputs; independent oracle; target pre/post boundary. The committed P6-C3 validator omission is source-confirmed and routed out of authority.
- Full P6-D closure: incomplete because the product/runtime/target chain and procedural isolation are not established.

## Target-Never-Accessed Provenance and Limit

- Strongest permitted E-only provenance is the sealed Main handoff, coder report, prior independent report, four ledgers, current visible command set, Git/candidate identity, and absence of any target-produced artifact in the reviewed candidate/runtime root.
- This visible Supervisor did not stat, list, read, hash, run Git in, analyze, query, or otherwise reference the target filesystem path in any command.
- The target-never-accessed claim is therefore supported only to the limit of those E-resident records and this visible task's command evidence. It is not elevated to an absolute audit claim, especially because the three interrupted internal lanes do not expose a complete activity history. No target access was performed to strengthen the claim.

## Git, Candidate, Report, and Quarantine Truth

- Pre-report Git: status `61 = 5` tracked + `56` untracked; index `0`; diff-check `0`; untracked reports `55 = 44 Investigation + 7 Supervisor + 4 coder`; sole other untracked path is `internal/graphhealth/p6d_outcome_projection_test.go`.
- Candidate remains exactly seven paths. Canonical construction is ordinal/case-sensitive relative-path sort, rows `relative-path|bytes|SHA256`, terminal LF. Total and hash remain `371,376` / `758066DE6FDF37DCCE1EEB22DD4B2B8F577CE8A58CA739445F9AFE6F72B0D3BE`.
- This report is the only durable artifact created outside the one permitted temp root. Expected post-freeze delta is one new Supervisor report only; external measurement follows content freeze.
- Protected roots remained unchanged (`files / bytes / zero-byte files / reparse points`):
  - `p6c3-repair` `1,673 / 76,978,267 / 19 / 0`
  - `p6c3-supervisor-rereview-01a02a84` `1,662 / 93,162,861 / 18 / 0`
  - `p6c3-main-closure-01a02abb` `1 / 343 / 0 / 0`
  - `p6d-coder-01a02abb` `10,053 / 971,554,982 / 84 / 0`
  - `p6d-main-verification-01a02b25` `1,512 / 64,175,973 / 19 / 0`
  - `p6d-supervisor-review-01a02b4f-260823-01` `1,522 / 66,564,280 / 19 / 0`
  - `p6d-identity-repair-01a02b25-01` `1,515 / 64,214,232 / 19 / 0`
  - `p6d-main-verification-01a02b71-02` `1,505 / 64,115,581 / 18 / 0`
- Review-owned root `E:\Anvien\.tmp\p6d-supervisor-resubmission-01a02bc2-260823-01`: `1,512 / 64,197,503 / 19 / 0`. Its `gocache` is `1,487 / 64,115,238 / 1`; `gomodcache` is `17 / 0 / 17`; `runtime` is `0 / 0 / 0`. It is not product-runtime evidence and was not cleaned.

## Required Fix List For Resubmission

1. Preserve the exact local six-field production/test repair; no further P6-D code change is currently required for the prior finding.
2. Main/Owner must resolve authority for the actual `scripts/full-build.ps1` pipeline, including its install/global/analyze effects, or approve a hermetic equivalent that preserves the same product contract.
3. Provision the exact missing dependency set in a new permitted non-quarantined E-only boundary and route the two source-confirmed fixture-layout failures if raw Go `./...` remains a required sub-gate.
4. Run the canonical full build from current bytes to exit `0`; retain every step, current runtime identity, and final analyze output/timing. Investigate the analyze liveness/performance symptom with benchmark evidence rather than assuming its cause.
5. Prove native Ladybug and every affected built reader from that runtime, without stale binary or cache substitution.
6. With explicit target authority, capture target pre-state, run current analyze/read boundaries, prove exact `Promise`, `Math.max`, and `Math.min` outcomes (`3/3`) with `0/3` in-repository analyzer gaps, run the independent oracle, and capture post-state.
7. Route the committed P6-C3 validator omission as an out-of-authority concern; do not silently edit or reopen P6-C3 from P6-D.
8. Run the next acceptance attempt in a fresh visible Supervisor task with no internal lanes from the outset and with process evidence limited to required name/PID fields.
9. Reseal candidate/Git/report/quarantine/target provenance after all gates, keeping P6-D unchecked/unstaged/uncommitted until complete independent acceptance.

## Overall Evaluation

The repaired graph-health implementation is locally coherent and the exact prior six-field consumer defect is closed on current source and tests. Evidence quality is sufficient only for that bounded clearance. The original report's full-build classification was materially wrong: the canonical script was not run, its analyze linkage was missed, and raw Go diagnostics were overclaimed. Full P6-D remains incomplete because the canonical-build/current-runtime/analyze-liveness/native/built-reader/target/oracle/pre-post chain is absent, and the visible review cannot claim procedural cleanliness after the two disclosed incidents. The correct handoff is to Main with P6-D still unchecked, unstaged, and uncommitted.

Report content is re-frozen here after the direct Owner correction. The prior external seal is superseded. Bytes, LF/CR, strict UTF-8/BOM, SHA-256, and final Git/report delta must be measured externally without another build, analyze, Anvien graph, or detect loop.
