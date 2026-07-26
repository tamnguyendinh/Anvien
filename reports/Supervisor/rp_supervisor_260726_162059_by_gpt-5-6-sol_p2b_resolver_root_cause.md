# Supervisor Report: P2-B resolver/module first divergence

Verdict: PASS

## Metadata

- Report file: `rp_supervisor_260726_162059_by_gpt-5-6-sol_p2b_resolver_root_cause.md`
- Review time: 260726 162059 +07
- Reviewer: `gpt-5-6-sol`
- Repo/project: Anvien investigation against `E:\cheapapp.org`
- Scope reviewed: `reports/Investigation/20260726_155230_p2b_resolver_root_cause.md`; fixed ambient `Promise`/`Math.max`/`Math.min` sites and two fixed barrel-mediated calls
- Claim reviewed: the ambient misses first diverge because TypeScript library declarations are outside the resolver workspace, while the barrel calls first diverge because imported-definition lookup stops at physical definitions in the barrel and does not follow its re-export binding
- Authority used: current user boundary, `E:\Anvien\AGENTS.md`, target source, TypeScript compiler API, fresh target graph, real-source extraction/resolution probe, and current Anvien production source
- Related artifacts: P1-C compiler oracles, P2-B real-source IR/control probe, raw selected graph capture, fresh Supervisor reruns and impact captures

## Executive Summary

- Problem: determine whether P2-B traced the earliest wrong boundary for the five fixed compiler-resolvable source sites, rather than merely restating persisted graph symptoms.
- Decision: PASS for the explicitly bounded cases. Independent reruns reproduce the actual/control delta without changing target source or graph, and current source inspection closes the causal path from extracted fact to unresolved diagnostic.
- Required outcome: accepted as bounded resolver root-cause evidence; downstream command/projection behavior remains P2-C scope.

## Source-Level Clearance Notes

- `E:\Anvien\internal\analyze\analyze.go`: clear. Lines 235–343 and 799–863 show that scanner-selected target files are parsed to `ScopeIR` and passed into resolution; no TypeScript Program or standard-library declaration universe is added.
- `E:\Anvien\internal\resolution\indexes.go`: clear. Lines 140–150 build declaration indexes only from `ir.Definitions`; lines 280–314 assemble import bindings; lines 460–483 search only `defsByFile[targetFile]`; lines 638–653 require a receiver type binding for `Math`.
- `E:\Anvien\internal\resolution\resolve.go`: clear. Lines 165–183 take the member-call path, lines 211–217 emit unresolved calls, and lines 321–334 resolve non-primitive type annotations only through the workspace.
- `E:\Anvien\internal\providers\tsjs\imports.go`: clear. Lines 73–101 emit a `reexport` import fact for the barrel; extraction is not the failing boundary in this case.
- `E:\Anvien\internal\providers\tsjs\types.go` and `references.go`: clear. The selected `Promise` type reference and `Math` member calls are emitted into the real-source IR.
- `E:\Anvien\internal\graphhealth\diagnostics.go`: clear. Lines 233–258 contain Go-oriented builtin/qualifier classification and otherwise fall through to `in_repo_unresolved`; this is downstream metadata classification, not the ambient resolver's first miss.

## Evidence Checked

Passed:

- Target graph remained SHA-256 `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090` before and after all Supervisor reruns.
- Fresh TypeScript 5.9.3 ambient oracle: `Promise` resolves to `lib.es*.d.ts`; `Math.max` and `Math.min` resolve to `lib.es5.d.ts`; selected-line diagnostics are empty.
- Fresh barrel oracle: both consumer aliases resolve through `index.ts` to `read-admin-commercial-config.ts:10`; selected-line diagnostics are empty.
- Fresh rerun of the direct real-source Go probe: actual barrel chain reports `ImportsResolved=10`, `ImportUsesEmitted=1`, `ResolvedCalls=37`, `UnresolvedReferences=542`; direct-declaration control reports `10`, `3`, `39`, `540`. The two fixed call diagnostics disappear and two `CALLS` edges to the unique declaration appear only in the control; ambient diagnostics remain.
- Raw/IR checks: the barrel has `ImportReexport`; consumers have named imports and free-call sites; the actual graph has consumer→barrel and barrel→declaration file/import facts but lacks consumer→declaration call/use facts.
- Target inventory: zero tracked `.d.ts` files and zero `.d.ts` File nodes; the compiler declarations used by the oracle live outside the target graph.
- Fresh upstream impact spot checks: `resolveImportedDef` CRITICAL (1 module, 16 processes), `buildWorkspace` CRITICAL (5 modules, 28 processes), `resolveTypeAnnotation` CRITICAL (4 modules, 35 processes), and `classifyDiagnostic` HIGH (3 modules, 3 processes). These are blast-radius warnings only.
- Verification freshness: current Anvien source/index and target graph/HEAD stated in the investigation report.

Failed:

- None within the bounded causal claim.

Not run:

- No full-repository differential replay with a modified resolver; the report proposes no fix and explicitly uses a five-file isolation control.
- No global ambient-name, namespace/static-member, multi-hop/wildcard re-export, or cycle inventory.
- No remediation, production edit, build, test update, detect-changes, or commit.

## Invariant Closure

- affected invariant: for the selected compiler-resolvable sites, identify the earliest stage where a correct source fact becomes an unresolved graph fact, and separate that stage from downstream diagnostic projection.
- sibling surfaces checked: target source, TypeScript symbol resolution, TSJS extraction, ScopeIR import/type/call facts, workspace construction, import binding, call/type resolution, graph diagnostics, persisted graph, and a direct-import control.
- residual unverified same-invariant surfaces: product policy for ambient declaration representation, arbitrary re-export chains, wildcards/cycles, and repository-wide prevalence remain explicitly outside the claim; no narrower evidence is used to infer them.

## Overall Evaluation

The report identifies two distinct causal families and supports each with a direct differential, source trace, and persisted graph evidence. It correctly rules out parser failure, stale graph, path-resolution failure, and graph persistence loss. It also avoids conflating the later Go-oriented diagnostic label with the initial ambient resolver miss. The bounded P2-B root-cause report is accepted.
