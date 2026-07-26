# Supervisor Review — Anvien Graph Identity and TypeScript Resolution Correctness v2

- Review time: `2026-07-27 01:24:18 +07:00`
- Verdict: `REJECT`
- Review scope: the frozen five-artifact plan set only. No production implementation was accepted, authorized, or modified by this review.

## Frozen snapshot reviewed

| Artifact | SHA-256 |
|---|---|
| `2026-07-26-anvien-graph-identity-resolution-v2-plan.md` | `AFF4DAE5F9F3DF34628EE0586D4C5CED02F2741F66F151A48DD0F2E28126C51E` |
| `2026-07-26-anvien-graph-identity-resolution-v2-evidence.md` | `08C1F870A69989A40AB36D847B657FA2ED98D2D7C925180887BD65CCC15DD262` |
| `2026-07-26-anvien-graph-identity-resolution-v2-benchmark.md` | `627910B8E72B63217289AA1FCA80F5508096A8A37993C94178DA25B6AE65997A` |
| `2026-07-26-anvien-graph-identity-resolution-v2-actual-status.md` | `04969DCD31C2A20400CD457CB2EC5D5CD7B68BCB865080DF53071A6C74D663A8` |
| `index-reader-matrix.md` | `A6FE5148341E048425E6240B403D21F37941299319FB194A24B5CBE5E4F97409` |

The hashes were recomputed immediately before this report and match the supplied manifest.

## Blocking finding

### B1 — The plan still violates the planner-template contract

The planner skill requires the standard plan to follow all four templates exactly (`[SKILL.md](E:/Anvien/.agents/skills/planner/SKILL.md:30)`). The plan template requires a structured Scope Boundary (`[plan.template.md](E:/Anvien/.agents/skills/planner/templates/plan.template.md:94)`–`[plan.template.md](E:/Anvien/.agents/skills/planner/templates/plan.template.md:100)`) and a structured Acceptance block with separate Source, Runtime/UI, DB/data, Behavior test, Cleanup/quarantine, Evidence IDs, and Actual-status fields (`[plan.template.md](E:/Anvien/.agents/skills/planner/templates/plan.template.md:142)`–`[plan.template.md](E:/Anvien/.agents/skills/planner/templates/plan.template.md:152)`).

Independent parsing of the frozen plan found:

- `102` unique checklist items and `99` eligible implementation/closure slices;
- all `99` eligible slices have the twelve explicit Pre-flight fields, and all `132` eligible numbered work steps have UI-flow, DB/data-flow, render-location, Mini-QA, and evidence-target fields;
- only `47/99` eligible slices use a structured Scope Boundary; `52/99` still compress it into one prose line;
- only `49/99` eligible slices use a structured Acceptance block; `50/99` still compress Acceptance into one prose line.

For example, `P1-C0A` has a compact Scope Boundary at `[plan.md](E:/Anvien/docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md:497)` and a compact Acceptance at `[plan.md](E:/Anvien/docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-plan.md:524)`. The Acceptance line does not explicitly state the required Runtime/UI, DB/data, Behavior test, Cleanup/quarantine, or Actual-status fields. The same pattern remains across the P1-C0A/P1-C0B/P1-D1..P1-D3 and P3–P7 slices listed by the structural scan.

`E0-P0A-TEMPLATE3` in the evidence ledger (`[evidence.md](E:/Anvien/docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-evidence.md:73)`) proves only the Pre-flight and per-work-step fields; it does not prove the mandatory Scope Boundary and Acceptance blocks. Consequently the current R5 freeze claim (`[actual-status.md](E:/Anvien/docs/plans/2026-07-26-anvien-graph-identity-resolution-v2/2026-07-26-anvien-graph-identity-resolution-v2-actual-status.md:135)`) is not sufficient evidence of template-complete plan readiness.

This is a release-blocking planning defect, not a cosmetic preference: without the explicit acceptance fields, each slice lacks independently auditable runtime/data/behavior/cleanup gates; without the explicit scope fields, ownership and preserve/out-of-scope boundaries are not machine-checkable under the required contract.

## Checks that cleared review

- P2 was materially decomposed: the frozen plan has `42` independently ordered P2 slices, including separate inventory, JSON/Ladybug/native/fallback, CLI/MCP/cache/embedding, HTTP/Web, registry/group, lease/GC, and failure-matrix owners. The P2 phase rule and each ordered list match the checklist exactly.
- Checklist-to-trace binding is exact: `101` non-P0 checklist IDs, `101` unique trace rows, and zero order differences.
- Evidence closure is complete for the frozen files: plan `747` unique references, benchmark `94`, and actual-status `53`; each has `0` missing declarations in the evidence ledger.
- Reader matrix checks pass independently: `R01..R195` contiguous, no duplicate IDs or anchors, all `195` source paths/function tokens present, and surface counts `S0=4`, `S1=5`, `S2=3`, `S3=56`, `S4=37`, `S5=40`, `S6=30`, `S7=4`, `S8=2`, `S9=17`, `S10=18`, `S11=14`.
- No unresolved template-brace placeholders, work markers, trailing whitespace, or bare `UI flow check: N/A.` / `Render location check: N/A.` lines remain.
- The current graph refresh is honestly recorded as `1,504` scanned Files, `676` parsed code files, `0` failed/unsupported/unknown, `84,532` nodes, and `123,372` relationships (`E0-P0A-GRAPH5`); the ledger explicitly distinguishes it from the historical `1,505`/`84,558`/`123,398` snapshot.
- The target boundary remains correct: `E:\cheapapp.org` is treated in place, its graph remains under the target `.anvien`, and no target source/fixture/report was copied into Anvien. The scanner `env`/`target`/`logs` omission remains explicitly quarantined and out of scope.

## Source-reality clearance

The planned root causes remain grounded in the current Anvien source:

- TypeScript variable extraction returns before creating a fact for non-identifier binding names (`[definitions.go](E:/Anvien/internal/providers/tsjs/definitions.go:64)`–`[definitions.go](E:/Anvien/internal/providers/tsjs/definitions.go:68)`), while `addDefinition` does not populate export visibility (`[definitions.go](E:/Anvien/internal/providers/tsjs/definitions.go:100)`–`[definitions.go](E:/Anvien/internal/providers/tsjs/definitions.go:111)`).
- Graph identity is generated from label, cleaned file path, qualified/name, and arity, without scope/range identity (`[indexes.go](E:/Anvien/internal/resolution/indexes.go:814)`–`[indexes.go](E:/Anvien/internal/resolution/indexes.go:823)`).
- Imported definitions are looked up directly in the target file definition bucket (`[indexes.go](E:/Anvien/internal/resolution/indexes.go:460)`–`[indexes.go](E:/Anvien/internal/resolution/indexes.go:475)`), supporting the planned barrel/re-export work.
- Diagnostic classification still uses target-text and qualifier heuristics (`[diagnostics.go](E:/Anvien/internal/graphhealth/diagnostics.go:233)`–`[diagnostics.go](E:/Anvien/internal/graphhealth/diagnostics.go:257)`), supporting the planned structured ambient/external outcome contract.

## Required resubmission gate

Before requesting another Supervisor review, expand every compact eligible Scope Boundary into the template’s four explicit fields and every compact eligible Acceptance into the template’s seven explicit fields (use explicit `N/A` reasons where applicable). Then rerun the structural audit with a new evidence ID, refresh the actual-status row and all four ledgers, rerun the frozen-hash/closure/matrix checks, and provide a new five-artifact SHA-256 manifest. Do not weaken the current P2 split, target boundary, or evidence denominator while making that correction.
