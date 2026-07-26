# Supervisor Report: P1-D Projection Parity

Verdict: PASS

## Metadata

- Report file: `rp_supervisor_260726_154849_by_gpt-5-6-sol_p1d_projection_parity.md`
- Review time: `260726 154849 +07:00`
- Reviewer: `gpt-5-6-sol`
- Repo/project: target `E:\cheapapp.org`; investigation workspace `E:\Anvien`
- Scope reviewed: P1-D raw graph, Cypher, and `file-detail` projection comparison
- Claim reviewed: for seven canonical source-site facts selected from five visible source occurrences, persisted raw graph and supported read-only projections preserve one-to-one presence/cardinality; reported field differences are projection-shape differences rather than duplicate physical source sites.
- Authority used: owner boundary and `E:\Anvien\AGENTS.md`; fresh target graph identity; target source; raw graph JSON; current Anvien CLI output.
- Related artifact: `E:\Anvien\reports\Investigation\20260726_154047_p1d_projection_parity.md`

## Executive Summary

- Problem: the report could have mistaken repeated gap-node/edge/projection representations and same-line companion facts for duplicate source occurrences.
- Decision: PASS for the bounded seven-ID parity claim. An independent raw JSON scan reproduced one gap node, one `HAS_RESOLUTION_GAP` relationship, and zero `SourceSite` nodes for every selected canonical ID. A fresh Cypher query reproduced one row and the expected projection default (`step: "0"`) for the Promise case. The expanded `file-detail` shape and the report's captures are consistent with the raw records.
- Required outcome: accepted as a bounded P1-D result; do not generalize it to all graph/projection facts or treat projection shape omissions as resolver correctness.

## Source-Level Clearance Notes

- `E:\cheapapp.org\modules\commercial-config\server\admin-commercial-config\read-admin-commercial-config.ts:13`: source occurrence and canonical Promise source-site ID are explicit; raw/projection count is one.
- `E:\cheapapp.org\modules\email\server\operations\email-operations-observability.ts:191`: nested `Math.max` and `Math.min` calls have distinct canonical IDs; no line-only duplicate was counted.
- `E:\cheapapp.org\modules\admin-operations\server\commercial-config\save-admin-commercial-config-mutation.ts:142`: call and type-reference companions are distinct fact families; each has one raw node/edge.
- `E:\cheapapp.org\modules\admin-operations\server\commercial-config\read-admin-commercial-config-route-view.ts:32`: same companion distinction and one-to-one raw/projection counts.
- No production Anvien source was edited or accepted as changed by this review.

## Evidence Checked

Passed:

- Fresh graph identity: target HEAD `a869876ab6262dacde6cd5d432d099a91852a646`, graph SHA-256 `DD3A3A987E95C91787E7C70EF68BA5D2F953D24A32D35DEFE38B8E8E606FF090`, `84,807` nodes and `114,125` relationships.
- Independent raw JSON scan (run during this review): all seven exact `sourceSiteId` values returned `nodeCount: 1`, `relCount: 1`, `sourceSiteNodeCount: 0`, zero null properties, and no raw `step` key.
- Independent `anvien cypher` query for the Promise gap: `row_count: 1`, `HAS_RESOLUTION_GAP`, `confidence: "1.000000"`, and `step: "0"`.
- Current expanded `file-detail` command: Promise unresolved sample carries the exact source-site ID, line 13/column 3, and the expected classification/actionability/status.
- Report captures and scripts under `E:\Anvien\.tmp\cheapapp-graph-root-cause-restart\` reproduce the seven-ID inventory, field mapping, and unsupported-boundary probes.
- Target tracked status was not changed by this review; no analyze was rerun.

Failed:

- None for the bounded claim.

Not run:

- No global graph/projection audit, database write test, or remediation validation; those are outside P1-D and cannot be inferred from this PASS.

## Invariant Closure

- affected invariant: canonical source-site identity and fact multiplicity across raw graph and read-only projections.
- sibling surfaces checked: raw `graph.json`, `file-detail`, Cypher `ResolutionGap`/`CodeRelation` projections, source-site node absence, same-line companion facts, and unsupported Cypher schema boundaries.
- residual unverified same-invariant surfaces: all facts outside the seven selected IDs; resolver classification correctness; downstream process/semantic completeness. These remain explicitly outside this bounded acceptance.

## Overall Evaluation

The report's counting key is technically sound: `sourceSiteId + factFamily` distinguishes companion facts and prevents line-only overcounting. The raw graph and current projections support the exact bounded parity statement. Projection omissions (`filePath`, range/hash/edge metadata) and Cypher's absent-`step` default are real shape boundaries and are correctly not promoted to persistence loss. PASS is limited to this slice; P1-D does not accept any resolver or global graph-accuracy conclusion.
