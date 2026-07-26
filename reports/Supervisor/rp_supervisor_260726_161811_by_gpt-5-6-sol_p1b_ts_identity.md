# Supervisor Report: P1-B TypeScript identity comparison

Verdict: PASS

## Metadata

- Report file: `rp_supervisor_260726_161811_by_gpt-5-6-sol_p1b_ts_identity.md`
- Review time: 260726 161811 +07
- Reviewer: `gpt-5-6-sol`
- Repo/project: Anvien investigation against `E:\cheapapp.org`
- Scope reviewed: `reports/Investigation/20260726_155807_p1b_ts_identity.md` and its bounded three-file evidence
- Claim reviewed: the selected source files contain six destructured bindings and four distinct same-name local declarations that the graph does not preserve one-for-one, while 21 selected exports exist as graph definitions but lack export/visibility metadata
- Authority used: current user boundary, `E:\Anvien\AGENTS.md`, target source, fresh target graph, TypeScript compiler API, and current Anvien source
- Related artifacts: P1-B oracle script/output, fresh Supervisor rerun, three fresh `file-detail` captures, target graph

## Executive Summary

- Problem: decide whether the P1-B report's bounded source-to-graph discrepancy claim is factually correct and sufficiently evidenced.
- Decision: PASS. A fresh independent rerun reproduces all three mismatch families against the unchanged graph hash, and direct source inspection explains why the comparison is semantically valid.
- Required outcome: accepted for the bounded three-file comparison only; root-cause ownership remains a separate P2-A review.

## Source-Level Clearance Notes

- `E:\cheapapp.org\modules\email\server\operations\email-operations-observability.ts`: clear for the bounded claim. The TypeScript AST rerun finds six array-binding names at lines 503–509, two `time` declarations at 207/214, and two `now` declarations at 262/501. Raw graph selection finds no six pattern names and one `Variable` node for each same-name pair.
- `E:\cheapapp.org\modules\release-distribution\server\release-distribution-publication-state.ts`: clear. Three source exports have matching graph definitions and zero selected visibility properties.
- `E:\cheapapp.org\modules\commercial-config\server\admin-commercial-config\read-admin-commercial-config.ts`: clear. The exported function at line 10 is present as one graph definition while `file-detail` reports `exportedSymbolCount=0`.
- `E:\Anvien\internal\providers\tsjs\definitions.go`: clearance evidence for interpretation only. `DefinitionFact` is created at lines 100–111 without setting `Visibility`; variable declarators with a non-identifier name are rejected at lines 64–68.
- `E:\Anvien\internal\resolution\indexes.go` and `E:\Anvien\internal\graph\types.go`: clearance evidence for interpretation only. Graph identity omits range at lines 814–824, and duplicate IDs replace existing nodes at lines 96–104.

## Evidence Checked

Passed:

- Fresh `anvien status` from `E:\cheapapp.org` with process-local Git trust: indexed and current commit both `a869876`, status up-to-date.
- Graph SHA-256 before and after the independent rerun: `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090` both times.
- Fresh rerun of `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\p1b-identity-oracle.mjs`: six source destructured names and zero graph matches; one graph `time` and one graph `now`; source export totals 17 + 3 + 1; zero selected definition visibility/export fields.
- Fresh `anvien file-detail ... --repo cheapapp-accuracy-direct --json` for all three files: exit 0, parsed, `stale=false`, and `exportedSymbolCount=0` for the selected files.
- Direct source inspection of the target and relevant Anvien owner paths.
- Verification freshness: current for target HEAD and the graph hash stated above.

Failed:

- None within the bounded claim.

Not run:

- No global TS/JS binding/export inventory; the reviewed claim explicitly does not make one.
- No remediation, build, test, detect-changes, or commit; this is investigation/review only.
- TypeScript 6.0.3 was not installed in the target. The oracle used 5.9.3; the reviewed facts are syntax/declaration and exact source-range facts, not version-sensitive type-behavior claims.

## Invariant Closure

- affected invariant: every bounded source binding/export fact must be compared to the persisted graph without conflating parser facts, graph identity, or presentation limits.
- sibling surfaces checked: independent TypeScript AST, raw graph nodes/properties, fresh `file-detail`, and the relevant extractor/identity source boundaries.
- residual unverified same-invariant surfaces: broader destructuring forms, overloads, export aliases/default exports, and repository-wide prevalence are outside the report's explicit scope; they do not block acceptance of the bounded claim.

## Overall Evaluation

The report is appropriately narrow and its three conclusions are independently reproducible. It does not overstate global prevalence, does not treat command projection as source truth, and does not claim a fix. The bounded P1-B comparison is accepted; P2-A must separately own and review the causal trace.
