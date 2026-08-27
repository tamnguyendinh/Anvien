# Child 06A D002 residual attribution

## Terminal verdict

`D002_RESIDUAL_ATTRIBUTION_BLOCKED`

Architect remains locked. The next owner is Main task `01a0415f-7e14-7193-9620-70b09d021b79`.

## Scope and decision rule

This is the read-only Child 06A P2-A attribution result for `B2-P2A-A001-D002 resolve_accesses`. It uses the completed checkpoint and bounded inspection of the current production owner and its direct synchronous emission helpers. It does not combine the targets, treat CPU samples as elapsed time, reopen A002/A003/A004, select a fix, or make an architecture decision.

The complete verdict required all three facts together:

1. one materially common exact residual cause on both Cheapapp and Restaurant Manager;
2. that cause's exact current production owner; and
3. its complete synchronous downstream graph/outcome/diagnostic/reference path.

The required conjunction fails at fact 1. The source proves the exact owners and paths of the observed branches, but the retained profiles do not prove one distinct materially common residual cause outside already accepted A002/A003 and rejected A004 ownership.

## Preserved measurement basis

CPU samples below are causal attribution only. They are not elapsed-time proof, are not percentages of the wall measurements, and are not averaged or combined across targets.

| Target | Profile identity | D002 wall basis | Parent / analyzer / process wall basis | Target inventory |
|---|---|---:|---:|---:|
| Cheapapp | 65,363 bytes; SHA-256 `806FAD79B12C567618778A8232E8521991882CE6C745AC2EBBE7B09ECC267107` | 9.380783200s | 20.472602300s / 93.531974900s / 95.630648200s | 26,042 accesses; 887 files |
| Restaurant Manager | 56,397 bytes; SHA-256 `4272FFAB0ABE4C14C470A139A31A601F31323B1373E0FA2986A09935456FA3E2` | 2.254679300s | 20.850792800s / 98.020546700s / 101.096911900s | 50,554 accesses; 1,234 files |

The retained `resolveAccess` CPU attribution remains target-specific:

| Target | Retained causal samples under `resolveAccess` |
|---|---|
| Cheapapp | `resolveAccess` cumulative 8.50s; `sort.SliceStable` 6.23s; `appendExportBindingEvidence` 5.61s; `sortExportBindingEvidence` 5.22s; `exportBindingEvidenceOrderFor` 5.19s; `encoding/json.Unmarshal` 5.03s; `emitUnresolvedReference` 1.41s; diagnostic appender 1.09s. |
| Restaurant Manager | `resolveAccess` cumulative 1.95s; `emitUnresolvedReference` 1.33s; diagnostic appender 1.07s; `decodeStructuredResolutionOutcome` 0.81s; JSON decoder 0.71s; outcome collector `record` 0.26s; `emitReference` 0.21s; resolution-outcome marshal 0.21s. |

## Residual attribution finding

Cheapapp's dominant exact child route is export-binding evidence projection and deterministic sorting:

`resolveAccess` -> `appendExportBindingEvidence` -> `sortExportBindingEvidence` -> `sort.SliceStable` -> `exportBindingEvidenceOrderFor` -> `json.Unmarshal`.

That route is owned by `internal/resolution/export_binding_proof.go` (`appendExportBindingEvidence`, lines 97-166; `sortExportBindingEvidence`, lines 231-246; `exportBindingEvidenceOrderFor`, lines 248-276). It is the rejected A004 direction and is not a distinct D002 common owner. The retained Restaurant Manager attribution does not show this route as a material child.

Restaurant Manager's dominant exact child route is unresolved outcome and diagnostic projection:

`resolveAccess` -> `emitUnresolvedReference` -> `recordRepositoryUnresolvedOutcome` -> outcome collector `record` / resolution-outcome marshal -> diagnostic appender -> diagnostic normalization -> `decodeStructuredResolutionOutcome` -> JSON decoder -> diagnostic bucket append -> graph node write-through.

Cheapapp also contains retained samples in `emitUnresolvedReference` and the diagnostic appender, but that overlap is the already accepted A002/A003 ownership. It cannot be relabeled as a distinct D002 residual owner. Conversely, Restaurant Manager's retained `emitReference`, outcome collector, and marshal samples do not provide a matching material Cheapapp attribution outside those excluded directions.

Therefore the broad `resolveAccess` envelope is common, but a broad envelope is not an exact residual cause. Access and file counts alone also do not establish a CPU cause.

## Exact current production ownership

The current access-resolution envelope is `internal/resolution/resolve.go::resolveAccess`, lines 630-741. Its current graph evidence carries HIGH file risk and CRITICAL symbol impact: 6 symbols, 4 files, 4 modules, and 32 processes. This is blast-radius evidence only; this lane makes no source edit.

The observed branch owners are exact:

- Access routing and terminal branch selection: `internal/resolution/resolve.go::resolveAccess` (lines 630-741).
- Internal resolved reference/outcome carrier: `internal/resolution/emit.go::emitReference` (lines 137-180), `internal/resolution/outcome.go::recordRepositoryResolvedOutcome` (lines 256-283), and `resolutionOutcomeCollector.record` (lines 83-106).
- Repository-unresolved outcome/diagnostic carrier: `internal/resolution/emit.go::emitUnresolvedReference` (lines 182-225), `internal/resolution/outcome.go::recordRepositoryUnresolvedOutcome` (lines 285-305), and `internal/graphhealth/diagnostics.go::diagnosticAppender.appendToNode` (lines 61-89).
- Structured diagnostic decoding: `internal/graphhealth/diagnostics.go::normalizeDiagnosticMetadata` (lines 275-287), `structuredResolutionDiagnosticPolicy` (lines 354-376), and `decodeStructuredResolutionOutcome` (lines 378-501).
- Export-binding evidence sorting: `internal/resolution/export_binding_proof.go::appendExportBindingEvidence`, `sortExportBindingEvidence`, and `exportBindingEvidenceOrderFor` at the line ranges above.
- TypeScript lookup carrier: `internal/resolution/resolve.go::recordTypeScriptLookup` (lines 926-977), `internal/resolution/outcome.go::recordTypeScriptOutcome` (lines 331-370), and `emitTypeScriptOutcomeDiagnostic` (lines 372-397).

These exact owners do not prove an exact owner for a distinct common D002 cause because the distinct common cause itself is absent from the retained evidence.

## Complete synchronous downstream paths

### Entry and branch selection

`ResolveBoundInto` iterates each file's `ir.Accesses` and calls `resolveAccess` (`internal/resolution/resolve.go`, lines 91-101). For each access, `resolveAccess` selects read/write kind, resolves an unqualified binding or qualified member/imported member, optionally checks the TypeScript standard-library boundary, and terminates through one of the carrier paths below.

### Missing source or repository-unresolved access

`resolveAccess`
-> `emitUnresolvedReference`
-> increment unresolved metric
-> `recordRepositoryUnresolvedOutcome`
-> `resolutionOutcomeCollector.record`
-> clone + validate + source-site map insertion
-> `marshalResolutionOutcome` / `json.Marshal`
-> construct `graphhealth.Diagnostic` carrying the encoded outcome
-> `diagnosticAppender.appendToNode`
-> graph source-node lookup
-> `normalizeDiagnosticMetadata`
-> `structuredResolutionDiagnosticPolicy`
-> `decodeStructuredResolutionOutcome`
-> streaming JSON token/decode validation
-> cached diagnostics lookup
-> `appendNormalizedDiagnostic` (bucket merge or stable sort)
-> write `diagnostics` property
-> `graph.AddNode`
-> diagnostic or unattributed metric.

The early `sourceForScopeOrFile` failure uses the same unresolved carrier with an empty source graph ID; the diagnostic appender rejects the missing endpoint and the unattributed metric is incremented.

### Internally resolved access

`resolveAccess`
-> optional retained semantic proof
-> optional `appendExportBindingEvidence`
-> projected proof construction + JSON marshal
-> `sortExportBindingEvidence`
-> comparator `exportBindingEvidenceOrderFor` + JSON unmarshal
-> evidence merge
-> `emitReference`
-> `recordRepositoryResolvedOutcome`
-> `resolutionOutcomeCollector.record`
-> clone + validate + source-site map insertion + JSON marshal
-> `ReferenceIndex.add` into source-scope and target-definition buckets
-> build `graph.Relationship`
-> `emitRelationship`
-> semantic edge merge/replace or `graph.AddRelationship`
-> resolved-reference and resolved-access metrics.

### TypeScript standard-library terminal access

When repository resolution fails without a repository ownership block:

`resolveAccess`
-> `lookupTypeScriptGlobal` or `lookupTypeScriptMember`
-> `recordTypeScriptLookup`.

For capability-unavailable, profile-excluded, or meaning-mismatch terminals:

`recordTypeScriptLookup`
-> `recordTypeScriptOutcome`
-> `resolutionOutcomeCollector.record` + JSON marshal
-> `emitTypeScriptOutcomeDiagnostic`
-> the same diagnostic appender / structured decoder / graph-node write-through path above.

For an externally resolved terminal, `recordTypeScriptLookup` records the authority result, external site, and resolved-external outcome. Synchronously later in the same `ResolveBoundInto` call, after all access iterations:

`finalizeTypeScriptAuthorityResults`
-> `emitTypeScriptExternalSymbols`
-> external graph-node emission
-> `emitTypeScriptExternalReference`
-> authority JSON marshal
-> `emitRelationship`
-> graph relationship merge/add
-> `ReferenceIndex.add`
-> resolved metrics.

If the TypeScript lookup returns no handled terminal, `resolveAccess` continues to the repository-unresolved carrier.

### Outcome finalization and carrier projection

After the access loop, `ResolveBoundInto` synchronously calls outcome collector `finalize` and `projectResolutionOutcomes` (`internal/resolution/resolve.go`, lines 103-120). Projection marshals finalized outcomes, adds outcome evidence to graph relationships and both reference-index buckets, and validates unresolved diagnostic carriers for source-site consistency. This post-loop projection is part of the production synchronous flow, but the retained per-`resolveAccess` causal samples do not make it a distinct common D002 cause.

## Exact blocker

The one missing fact is:

> The retained evidence does not show one same synchronous child symbol, outside accepted A002/A003 and rejected A004 ownership, with material CPU attribution under `resolveAccess` on both Cheapapp and Restaurant Manager.

Without that cross-target fact, no distinct materially common exact residual cause exists in evidence, and consequently no exact production owner can be assigned to such a cause. D002 cannot open Architect.

## Change and validation boundary

- Authorized and produced path only: `reports/Investigation/rp_child06a_d002_residual_attribution.md`.
- No source, test, plan/ledger, script, target repository, or other report was edited.
- No staging or commit is authorized.
- Exact report-path status: `?? reports/Investigation/rp_child06a_d002_residual_attribution.md`; the authorized report is the sole path in the scoped status result.
- `git diff --check -- reports/Investigation/rp_child06a_d002_residual_attribution.md`: PASS (exit 0).
- `git diff --cached --name-only`: empty; staged set proven empty.
