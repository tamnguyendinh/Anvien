# Supervisor Report: P2-A TypeScript extraction and graph identity root causes

Verdict: PASS

## Metadata

- Report file: `rp_supervisor_260726_164714_by_gpt-5-6-sol_p2a_ts_identity.md`
- Review time: `260726 164714 +07`
- Reviewer: `gpt-5-6-sol`
- Repo/project: Anvien investigation against `E:\cheapapp.org`
- Scope reviewed: `reports/Investigation/20260726_161748_p2a_ts_identity_root_cause.md`, the independent red-team report `reports/Investigation/20260726_161952_p2a_identity_root_cause.md`, and their current raw probes
- Claim reviewed: for the three fixed target files, six array bindings are lost before a `DefinitionFact`, two same-name local pairs collide at graph identity, and 21 source exports lose visibility metadata; the report identifies the earliest Anvien source boundary for each loss
- Authority used: current user boundary, `E:\Anvien\AGENTS.md`, target source and graph, direct TypeScript/ScopeIR probes, current Anvien source, and fresh Anvien impact evidence
- Related artifacts: `p2a_identity_source_graph_diff.json`, `supervisor_scopeir_rerun.json`, `p2a_anvien_root_cause_trace.json`, `p2a_target_state_after.json`, and the P1-B Supervisor PASS report

## Executive Summary

- Problem: decide whether the P2-A TypeScript identity/extractor trace explains the bounded P1-B source-to-graph mismatches rather than merely repeating graph symptoms.
- Decision: PASS for the explicitly bounded three-file cases. Fresh reruns reproduce the six missing pattern bindings, two graph identity collision groups, and 21 represented-but-unmarked exports; direct source inspection closes each first-divergence path.
- Required outcome: accept as bounded root-cause evidence only. This does not authorize remediation or a repository-wide TypeScript accuracy claim.

## Source-Level Clearance Notes

- `E:\cheapapp.org\modules\email\server\operations\email-operations-observability.ts`: clear for the bounded claim. The current source contains `time` locals at lines 207/214, `now` locals at lines 262/501, and the six array bindings at lines 503–509. The graph rerun contains one `Variable` node for `time` (line 207) and one for `now` (line 262), and no variable nodes for the six pattern names.
- `E:\cheapapp.org\modules\release-distribution\server\release-distribution-publication-state.ts`: clear for the export subset. Three source exports are represented in graph definitions and none has an export/visibility flag.
- `E:\cheapapp.org\modules\commercial-config\server\admin-commercial-config\read-admin-commercial-config.ts`: clear for the unique-name control and export subset. The function at line 10 is represented with the expected range, but its export metadata is absent; the adjacent `Promise` gap is outside this P2-A identity verdict.
- `E:\Anvien\internal\providers\tsjs\definitions.go:64-68`: clear. A `variable_declarator` whose `name` is not an `identifier` returns before `addDefinition`, matching the six array-pattern losses.
- `E:\Anvien\internal\providers\tsjs\definitions.go:95-112`, `internal/resolution/indexes.go:814-824`, and `internal/graph/types.go:96-104`: clear. Range-bearing facts are created, graph IDs omit range/scope, and duplicate IDs replace the existing graph node.
- `E:\Anvien\internal\scopeir\facts.go:3-18` and `internal/resolution/emit.go:158-180`: clear. `DefinitionFact` can carry `Visibility`, but the TSJS construction leaves it empty and the emitter only serializes a non-empty value.

## Evidence Checked

Passed:

- Fresh target identity check: HEAD `a869876ab6262dacde6cd5d432d099a91852a646`; graph SHA-256 `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`; graph inventory `1,359` files / `84,807` nodes / `114,125` relationships. The graph hash remained unchanged after the review probes.
- Independent TypeScript AST/graph rerun (`node .tmp/cheapapp-graph-root-cause-restart/p2a-identity/p2a_identity_probe.mjs`): six source array bindings have zero graph matches; the selected source exports total `17 + 3 + 1 = 21`, with `17/17`, `3/3`, and `1/1` represented but zero graph rows flagged exported; the selected graph locals are `time:207` and `now:262`.
- Independent ScopeIR rerun (`go run .tmp/cheapapp-graph-root-cause-restart/p2a-identity/scopeir_probe`): the email file has `169` extracted definitions, both `time` facts and both local `now` facts remain range-bearing, six pattern names are absent, and exactly two graph-ID collision groups are reported. The probe reports empty `Visibility` for the selected definitions.
- Direct source inspection of the target files and the five Anvien owner paths listed above; the source excerpts are also hash-pinned in `p2a_anvien_root_cause_trace.json`.
- Fresh Anvien self-graph owner checks after `anvien analyze --force`: the exact `graphIDForDef`, `Graph.AddNode`, `emitDefinitionNodes`, and `exportedSymbol` UIDs carry CRITICAL workflow warnings. These warnings describe blast radius only and do not change the bounded verdict.
- Target boundary: `E:\Anvien\p0-rc-c-fixture` is absent (`FIXTURE_REMOVED`); all probes read the real target path and write only under `E:\Anvien\.tmp`.

Failed:

- None within the bounded P2-A identity/extractor claim.

Not run:

- No repository-wide TypeScript binding/export inventory beyond the fixed three files.
- No policy decision for overloads, namespace members, default/export aliases, or the desired identity key.
- No production edit, remediation test, full build, detect-changes, or commit; this is an investigation review only.

## Invariant Closure

- affected invariant: source declaration/binding identity and export visibility must survive extraction, scope IR, graph identity, and graph emission without conflating parser facts with persisted-node identity.
- sibling surfaces checked: target AST, raw target graph nodes/properties, direct TSJS ScopeIR, graph-ID construction, `Graph.AddNode`, visibility emission, and the bounded file-detail summaries.
- residual unverified same-invariant surfaces: other destructuring patterns, export syntaxes, overload/namespace identity, and repository-wide prevalence remain outside the stated claim and are not required to accept this bounded slice.

## Review Caveat

The collision mechanism is accepted, not a claim that one particular declaration is the universal winner. `Graph.AddNode` is last-write replacement, and the current sorted emission run leaves the line-207/line-262 nodes visible; a future remediation must define an injective identity/order policy explicitly rather than rely on the observed winner.

## Overall Evaluation

Both P2-A reports make the same bounded causal claim, and the independent reruns reproduce it from the real target without changing target source or graph. The evidence identifies the earliest loss for array bindings, identity collision, and export visibility, while preserving the unresolved scope and remediation boundary. PASS is limited to these fixed cases.
