# Main Verification — P3-C2 BLOCKED handoff and observability boundary

Verdict: `PASS` for the validity of the QA `BLOCKED` handoff only. This is not `E3-P3C2-REVIEW1`, does not accept P3-C2, and does not open Pn-A.

## Metadata

- Review time: `2026-08-20 14:48:39 +07:00` (`Asia/Bangkok`)
- Reviewer role: Orchestration Main using the zero-trust Supervisor review guard; review/route only, no implementation or target access
- Repository: `E:\Anvien`
- Open slice: Child 03 `P3-C2` only
- Claim reviewed: the durable QA result is correctly classified `BLOCKED` because the persisted target graph proves the graph-and-later subset while the normal CLI cannot directly expose the required ScopeIR fact/path/owner-local-binding layer
- Authority: current Owner delegation, `AGENTS.md`, orchestration and working-rules skills, complete Child 03 plan/ledgers, roadmap, `docs/contracts/graph-accuracy-contract.md`, and accepted P3-C predecessor evidence
- QA artifact: `reports/QA/rp_qa_260820_143739_by_gpt-5_p3c2_target_binding_validation.md`

## Artifact and Git verification

- QA report: `19,701` bytes / `290` lines / SHA-256 `2109A1C9451A4D910D9BDFEFD3991F8C9FB061621784D275468DA38E15C19487`; path and hash match the visible lane handoff.
- HEAD / parent: `656a0445ff3e25b6225b994cdaf7cf1b35eb665c` / `a569b8674fefdaa757cf7fdf63f454caf7925215`.
- Accepted P3-C hashes remain exact:
  - `internal/resolution/resolve.go`: `C1FF5C515D401ECAD4FBF93C271DF4AC19101B2F5410D4174F7F598502BBC96A`
  - `internal/resolution/p3c_binding_occurrence_test.go`: `9BAE2F63575C313B5F5F8EF4C265360BCB38F8D2F616CDCDC8942C57A080CE7F`
  - `internal/lbugload/p3c_binding_occurrence_persistence_test.go`: `C704B45FD350F2A1B064D79E78B4DC99F6378D44358B4E86149A09FA38D4A850`
- No tracked modification or staged path exists. Before this report, the only untracked paths were the incoming Main rotation handoff and the new QA report.
- Proposed `cmd/binding-contract-probe/main.go` and `main_test.go` do not exist and are not tracked.

## Source-level verification

- `internal/providers/tsjs/extract.go:20-39` returns a production `scopeir.ScopeIR`; `collector.result` places `Definitions`, `BindingLeaves`, and `Scopes` into that value.
- `internal/scopeir/facts.go:3-38`, `81-91`, and `126-136` expose the exact Definition, binding-leaf path/provenance, owner `OwnedDefIDs`, and local Binding metadata required by P3-C2.
- `internal/cli/command.go:190-205` sets both `WriteGraphSnapshot=true` and `ReleaseScopeIRsAfterResolution=true` on the normal analyze path.
- `internal/analyze/analyze.go:116-122` excludes `ScopeIRs` from JSON with `json:"-"`; `351-355` clears ScopeIR and parsed-IR state after resolution.
- `cmd/graph-accuracy-probe/main.go` is a graph-snapshot comparison command. It requires graph inputs/output, calls `graphaccuracy.Run`, and can copy/write graph artifacts; adding ScopeIR extraction would mix responsibilities.
- The report's observability claim is therefore source-backed: persisted graph/Ladybug/Cypher cannot directly prove `BindingLeafFact.Path`, matching Definition facts, owner `OwnedDefIDs`, or owner-local `BindingLocal` rows.

## Fresh Anvien evidence and blast radius

`anvien --help` preceded graph use. A fresh self-index with both forbidden-tree exclusions passed at current HEAD:

```text
anvien analyze . --force --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**"
files: scanned=1110 parsed_code=624 failed=0
graph: nodes=80194 relationships=119137
indexed/current commit: 656a0445ff3e25b6225b994cdaf7cf1b35eb665c
```

The `+1 file / +17 nodes / +17 relationships` relative to the QA report is the newly created QA report; it does not change HEAD or code ownership.

Current owner evidence:

| Existing owner | File detail | Upstream file impact | Decision |
| --- | --- | --- | --- |
| `internal/providers/tsjs/extract.go` | `36` symbols, `22` inbound, `32` outbound, `1` flow, `6` tests, HIGH, current/non-stale | CRITICAL: `21` impacted symbols / `10` files / `1` flow | consume only; do not edit for validation |
| `cmd/graph-accuracy-probe/main.go` | `40` symbols, `0` inbound, `8` outbound, `4` flows, HIGH, current/non-stale | LOW: `2` impacted symbols / `1` file / `1` flow | wrong responsibility for ScopeIR fact observation |
| `internal/graphaccuracy/graphaccuracy.go` | `300` symbols, `142` inbound, `3` outbound, `11` flows, `11` tests, HIGH, current/non-stale | CRITICAL: `107` impacted symbols / `32` files / `1` flow | graph-comparison owner; do not edit |

HIGH/CRITICAL are blast-radius warnings. They support isolation into a new command; they do not prohibit necessary work.

## Target evidence classification

Main did not access `E:\cheapapp.org`, as explicitly required by the incoming authority. The QA artifact was checked for internal completeness and exact traceability rather than independently rerunning target commands.

- `E3-P3C2-ORACLE1`: the report records six names, paths `[0]..[5]`, exact ranges/selections, one lexical owner, and six expected downstream targets without copying source bodies.
- Persisted subset of `E3-P3C2-TARGET1`: the report records normal built analyze `1359/887/0`, graph `90,899/121,868`, six distinct Variables, six `DEFINES`, six exact binding `ACCESSES` endpoints, and two zero-row gap queries.
- `E3-P3C2-BOUNDARY1`: the report records exact pre/post HEAD, status hash, tracked diff object, untracked-manifest hash, oracle-source hash, analyzer-owned `.anvien` hashes, and equality results. No contradictory evidence is present in the lane handoff.
- These target facts are not promoted to P3-C2 acceptance because direct ScopeIR fact/path evidence remains missing and Main did not independently rerun the target.

## Invariant closure and route decision

- Affected invariant: direct conservation proof across source oracle -> `BindingLeafFact`/Definition/owner-local Binding -> persisted graph occurrence -> exact lexical endpoint -> zero binding-caused gap.
- Cleared subset: oracle, persisted occurrences/`DEFINES`, exact endpoints, gap outcome, and target preservation are fully recorded by the QA lane.
- Residual same-invariant surface: direct ScopeIR fact/path/provenance/owner-local-binding observation for the six target names.
- Handoff verdict: the QA lane is correctly `BLOCKED`; it is not FAILED and no production repair should open.
- Route: one visible coder lane inside P3-C2 may implement only a reusable metadata-only `cmd/binding-contract-probe` plus focused synthetic test. P3-C2 QA then resumes against the real target and remains subject to a separate visible Supervisor gate.

## Authorized implementation boundary for routing

- Editable: `cmd/binding-contract-probe/main.go`, `cmd/binding-contract-probe/main_test.go`, and one timestamped coder report only.
- Purpose: read one explicitly named TypeScript file, parse through the production parser, call `tsjs.Extract`, filter explicit local names, and emit deterministic metadata-only JSON to stdout.
- Required facts: one variable-context BindingLeaf with exact range/selection/path/provenance, one matching Variable Definition, exact owner `OwnedDefIDs`, and exactly one owner-local `(BindingLocal,name,DefID)` per selected name.
- Forbidden: source-body output; target write; `--out`; edits to providers, ScopeIR, analyze, CLI, resolution, graphaccuracy, existing probes, plans/ledgers, or accepted reports; target access by the coder lane; production repair or semantic fallback.
- Test: independent synthetic TypeScript only, code first then tests, deterministic output, fail closed on missing/duplicate/ambiguous fact-owner relationships.
- Validation: holder-clean full build after final bytes, focused command tests, repo-native affected regression, cleanup, durable coder handoff; no detect/stage/commit/push/self-acceptance.

## Overall evaluation

The QA report is a valid, evidence-backed `BLOCKED` handoff. Its graph and boundary observations are useful retained evidence, but they are narrower than complete P3-C2 acceptance. The smallest responsible route is an isolated two-file metadata probe implemented by a coder lane, followed by resumed target QA and independent Supervisor review.
