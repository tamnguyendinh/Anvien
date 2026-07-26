# Supervisor Report: P2-C projection, command, and derived behavior

Verdict: PASS

## Metadata

- Report file: `rp_supervisor_260726_163556_by_gpt-5-6-sol_p2c_projection_command.md`
- Review time: 260726 163556 +07
- Reviewer: `gpt-5-6-sol`
- Repo/project: Anvien investigation against `E:\cheapapp.org`
- Scope reviewed: corrected `reports/Investigation/20260726_161241_p2c_projection_command_root_cause.md`, P1-D/P1-E fixed cases, and the P2-C raw/source/control evidence
- Claim reviewed: downstream `context`, `impact`, and `file-detail` preserve the bounded facts/edges supplied by the graph; two bounded DB/Cypher representation changes occur at CSV/native-read boundaries; process completeness remains unresolved
- Authority used: current user boundary, `E:\Anvien\AGENTS.md`, current target graph, Anvien source, current probes, targeted tests, and the prior Supervisor rejection/closure record
- Related artifacts: prior REJECT `rp_supervisor_260726_163219_by_gpt-5-6-sol_p2c_projection_command.md`; corrected P2-C report; five fresh process reruns; target/projection captures

## Executive Summary

- Problem: decide whether P2-C correctly separates downstream projection behavior from upstream graph loss and preserves the derived-process uncertainty boundary.
- Decision: PASS for the bounded scope. The stale process differential identified in the prior review was corrected; fresh reruns reproduce the corrected no-change control, and source/test/raw evidence closes the remaining bounded claims.
- Required outcome: accepted as bounded P2-C evidence; no global process/semantic or public-contract conclusion follows.

## Source-Level Clearance Notes

- `E:\Anvien\internal\mcp\context_index.go:12-77`: clear. `contextNeighborhood` scans exact graph relationships and adds process rows only for direct `STEP_IN_PROCESS` edges from the selected symbol.
- `E:\Anvien\internal\mcp\impact.go:554-631,943-999`: clear. Impact traverses configured existing edges and aggregates process memberships only from existing selected-source step edges; it does not resolve missing calls.
- `E:\Anvien\internal\filecontext\context.go:829-890` and `internal\filecontext\compact.go:169-190,489-507`: clear. The unresolved DTO has exactly the ten documented fields; the seven selected rows are retained.
- `E:\Anvien\internal\lbugload\csv.go:372-400`: clear for the representation finding. A nil `Relationship.Step` is initialized as integer `0` before CSV/database load.
- `E:\Anvien\internal\lbugnative\native_ladybugdb.go:132-165`: clear. Every native tuple value is converted through `tupleString` into a string row value.
- `E:\Anvien\internal\processes\processes.go:59-149,152-179,256-363`: clear for the uncertainty boundary. Calls are filtered, entry candidates/branching/traces are capped and deduplicated, and only surviving traces emit process-step edges; the control result is not treated as product truth.

## Evidence Checked

Passed:

- Corrected report now records the reproducible process control: actual `3,771 calls / 662 processes / 2,761 steps / 0 selected consumer memberships / 0 wrapper memberships`; +2-call control `3,773 / 662 / 2,761 / 0 / 0`.
- Five fresh independent runs in `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p2c\supervisor\process-run-1.json` through `process-run-5.json` reproduce those exact values.
- Fresh target boundary probe: graph SHA-256 before/after `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`; 84,807 nodes / 114,125 relationships; 662 Process nodes / 2,761 step edges; zero selected symbol/file process edges.
- Fresh P1-D parity rerun: each of seven canonical facts has one raw gap node, one raw relationship, zero SourceSite nodes, and one file-detail sample.
- Source trace confirms `context`/`impact` cannot synthesize the two P2-B-missing consumer calls and that empty process arrays match raw selected membership.
- Source trace confirms the first representation changes are `nil Step -> 0` in `relationshipCSVRow` and native scalar -> string in `tupleRow`/`tupleString`.
- Targeted validation rerun: `go test ./internal/mcp ./internal/filecontext ./internal/lbugload` passed for all three packages.
- Prior blocking finding is closed: corrected report, evidence ledger, benchmark ledger, and actual-status now agree on the no-change process control and avoid a monotonicity/sensitivity claim.
- Verification freshness: current target graph/source and fresh reruns; no target analyze was needed for this read-only slice.

Failed:

- None within the corrected bounded claim.

Not run:

- No full process/semantic completeness audit, product-flow authority review, or public Cypher contract decision.
- No remediation, source edit, build, detect-changes, or commit.

## Invariant Closure

- affected invariant: downstream projections must preserve the presence/cardinality of supplied graph facts, and representation changes must be distinguished from upstream fact loss; derived process output must not be promoted to a completeness claim.
- sibling surfaces checked: raw target graph, P1-D raw/Cypher/file-detail captures, P1-E context/impact captures, MCP/filecontext/lbugload source and tests, native adapter source, process producer/control, and corrected report/ledgers.
- residual unverified same-invariant surfaces: global process/community/semantic completeness and any contract requirement for nullable/type-preserving Cypher output remain explicitly unresolved and outside this PASS.

## Overall Evaluation

After the prior evidence-integrity rejection was corrected, the P2-C report is internally consistent and technically bounded. It establishes command fidelity for present edges and line-traces the two representation changes without mistaking them for missing records. The process control is reported as a no-change observation only, which is the strongest claim supported by current evidence. PASS is limited to this scope.
