# Coder Report: Child 04 P4-C2 TypeAlias compatibility rejection repair

Status: **READY_FOR_QA**

## Metadata

- Report time: `2026-08-21 13:11:29 +07:00` (`Asia/Bangkok`).
- Role: bounded Coder rejection-repair lane.
- Repository: `E:\Anvien`.
- Git basis: branch `master`, HEAD `310502a88849fe75f86a45a987ba21490d19dbe2` (`docs(orchestration): fix skill name frontmatter`).
- Accepted predecessor: P4-C commit `c99c4070b66e7a96be8c9fa2721a0335a1f94877`.
- Rejection authority: `reports/Supervisor/rp_supervisor_260821_123030_by_gpt-5_child04_p4c2_durable_retry_review1.md`, canonical SHA-256 `8E37F4B126ABB38A5DCE071C35937F59D9152EFA135B9245408CA138BEC781F7`.
- Historical blocker report retained unchanged: `reports/Coder/rp_coder_260821_130238_by_gpt-5_child04_p4c2_typealias_compatibility_repair_blocked.md`, `12,062` bytes, actual SHA-256 `A7CAAF9C5B16129AEC7E6F847B668077AE2FE1F11A8F1779970722C7875F2634`.
- Active scope: only the rejected direct exported TypeAlias definition compatibility and downstream FileContext invariant.
- Target `E:\cheapapp.org`: not accessed.
- Next owner: Main for QA routing. This report is a Coder handoff, not a QA or Supervisor acceptance verdict.

## Outcome

The exact production-first repair and test-after-code regression update are complete within the authorized two-file boundary. After Main safely released the editor-owned MCP handles that had blocked runtime replacement, the resumed canonical full build passed, the focused rejection test passed, the nearest resolution -> FileContext -> Ladybug boundary passed, and all four affected packages passed their full package tests.

The candidate is `READY_FOR_QA`. Fresh target-specific `21/21` positive, `11/11` negative, Graph JSON/Ladybug record parity, and zero-state comparison remain QA-owned and were deliberately not run by this lane.

## Exact rejected invariant

- P001-P014 and P018 are direct exported TypeAlias definitions. Definition compatibility must be `isExported=true`, and production FileContext must consequently return `exported=true`.
- Their Export facts must remain `typeOnly=true` with `meanings=[type]`; definition compatibility must not synthesize runtime-value eligibility.
- Six Function positives, 11 negative controls, access/export separation, Graph JSON/Ladybug parity, zero orphan/diagnostic state, and zero Child 05 terminal/resolved/public-API state must remain unchanged.

## Invariant Family Map

- Family: local source export fact -> direct source-export membership -> definition `isExported` compatibility -> FileContext exported-symbol result.
- Authority source: a `ScopeIR.ExportFact` carrying `LocalDefID` after `exportProjectionNodes` has rejected source re-export (`TargetRaw != nil`), missing definition, and cross-file definition states.
- Sibling surfaces: Function/value direct exports, TypeAlias/type-only direct exports, definitions without export facts, anonymous default expressions, source re-exports, access visibility, Export-node type/meaning fields, Ladybug persistence, and FileContext canonical-field precedence.
- Forbidden fallback: runtime-value eligibility deciding whether a local source definition is directly exported.
- Forbidden expansion: FileContext/Ladybug production changes, terminal/barrel/alias-chain/cycle/ambiguity/public-API state, target access, or Child 05 work.
- Stale artifact repaired: `TestP4CProjectsExportFactsAndRuntimeCompatibility` encoded the rejected TypeOnly compatibility expectation as false.

## Confirmed root cause and production repair

`internal/resolution/emit.go::exportProjectionNodes` was inside the validated local-definition branch but populated `directExportDefIDs` only when `exportFactIsRuntime(fact)` returned true. That helper rejected `TypeOnly` and type-only meaning facts, conflating runtime eligibility with source-export membership.

Production-first change:

- `internal/resolution/emit.go:374` now records every validated local export definition in `directExportDefIDs` independently from `TypeOnly` and meaning lanes.
- The now-unused `exportFactIsRuntime` helper was removed; it had no other caller.
- `exportGraphNode` is unchanged, so `typeOnly`, `meanings`, source provenance, access separation, and the absence of Child 05-derived fields remain intact.
- A definition without a local Export fact is never inserted into `directExportDefIDs`, preserving negative controls.
- Function/value facts were already eligible and remain members; the repair only adds previously suppressed validated type-only local definitions.
- `internal/filecontext/context.go`, `internal/lbugload/csv.go`, and `internal/lbugschema/schema.go` remain byte-untouched and continue to consume or persist canonical fields.

## Test-after-code repair

Only after production code was changed, `internal/resolution/p4c_export_projection_test.go` was updated:

- the runtime-positive fixture uses `Function`;
- the rejected type-only fixture uses `TypeAlias`;
- direct TypeAlias definition compatibility expects `isExported=true`;
- its Export node must still expose `typeOnly=true` and exactly `meanings=[type]`;
- private/public visibility separation checks, local-definition endpoint checks, and zero terminal/public-API/resolved-target field checks remain active.

## Exact candidate diff and identities

| Path | Diff | Bytes | LF | SHA-256 | Git blob |
|---|---:|---:|---:|---|---|
| `internal/resolution/emit.go` | `+1/-15` | `26,772` | `815` | `B1B8F22AEFE94F72483CD284B7E267FE3897412C04D581AFAB7F9FE30F0B2867` | `3c6ede9c93531a634db32b8b0100c38bde0ffaeb` |
| `internal/resolution/p4c_export_projection_test.go` | `+23/-6` | `9,247` | `203` | `F575F16D20D8F95C347E142BCBBEE96E18DA84D23E1433E00B6C2545132E0162` | `ee9076e20adea437222e3c2df8cc28e9ad61e0ae` |

- Total: `24` insertions / `21` deletions across exactly two tracked repair files.
- Resume identity gate matched both files byte-for-byte, LF-for-LF, SHA-for-SHA, and Git-blob-for-Git-blob against the historical blocker report.
- Final `git diff --check -- <two repair files>` exited `0`.
- Before this lane edited them, both files had zero working-tree diff and zero drift from accepted P4-C through then-current HEAD.

## Anvien freshness, ownership, and blast radius

These gates were completed before code and were not rerun during the build-blocker resume:

- Repo-local `anvien --help`: PASS.
- Initial fresh self graph: exit `0`; scanned/parsed/failed `1,915/736/0`; `114,487` nodes / `157,336` relationships.
- Required candidate refresh after an unrelated orchestration-skill HEAD advance: `1,915/736/0`; `114,511` nodes / `157,326` relationships; indexed/current commit `310502a88849fe75f86a45a987ba21490d19dbe2`.
- `emit.go` file-detail: `150` symbols, `43` inbound, `226` outbound, `77` local relationships, `13` linked flows, `17` linked tests, file risk `HIGH`, non-stale.
- `emit.go` upstream file impact: `CRITICAL`; `30` impacted symbols, `24` direct, `5` affected files, `1` affected flow.
- `emitDefinitionNodes`: `CRITICAL`; `6` impacted symbols, `1` direct, `4` modules, `34` processes.
- `exportProjectionNodes`: `CRITICAL`; `4` impacted symbols, `1` direct, `2` modules, `26` processes.
- `exportFactIsRuntime`: `CRITICAL`; `3` impacted symbols, `1` direct, `1` module, `14` processes.
- Focused test file-detail: `39` symbols, zero inbound/unresolved, risk `LOW`; test-function impact `LOW` with zero affected symbols/modules/processes.
- HIGH/CRITICAL results were handled as blast-radius warnings, not edit bans; the candidate stayed within two files.

## Cleared build-holder gate

Main reported that Restart Manager holders PID `11656`/`14924` and wrapper PID `14616`/`14120` were safely stopped, with `remaining=[]` and exclusive no-share open PASS at `2026-08-21 13:06:42 +07:00`.

The resumed Coder lane independently checked once at `2026-08-21 13:07:28 +07:00`:

- repo-local runtime processes: `[]`;
- verified build-related processes: `[]`;
- analyze lock: absent;
- exclusive no-share open of `E:\Anvien\anvien\bin\anvien.exe`: PASS;
- candidate identities: exact;
- index: empty.

No process was terminated by the resumed Coder lane.

## Build and verification evidence

### Canonical full build — PASS

Command:

```text
npm run full-build
```

Result: exit `0`.

- Repo-local runtime replacement succeeded in postinstall, prepare, and main build paths.
- Built CLI version: `1.2.8`.
- Web production build: `2,943` modules transformed; Vite build `25.60s`.
- Chunk-size and mixed dynamic/static-import messages were warnings only.
- Final canonical analyze: scanned/parsed/failed `1,917/736/0`; graph `114,546` nodes / `157,361` relationships at `E:\Anvien\.anvien\graph.json`.

### Focused rejection regression — PASS

Command:

```text
go test ./internal/resolution -run '^TestP4CProjectsExportFactsAndRuntimeCompatibility$' -count=1 -v
```

Result: exit `0`; exact test PASS; package `ok` in `0.180s`.

This test proves at the real resolver graph boundary that:

- direct TypeAlias definition compatibility is true;
- its Export fact remains `typeOnly=true`, `meanings=[type]`;
- the Function positive remains true;
- private/public access visibility remains independent;
- local-definition references have existing endpoints;
- no terminal/public-API/resolved-target field is synthesized.

### Nearest resolution -> FileContext -> Ladybug boundary — PASS

Command:

```text
go test ./internal/resolution ./internal/filecontext ./internal/lbugload ./internal/lbugschema -run '^TestP4C(ProjectsExportFactsAndRuntimeCompatibility|RejectsOrphanLocalExportFact|ExportedSymbolUsesCanonicalExportFieldBeforeAccessVisibility|ExportCSVAndLoaderPreserveFactFields|ExportSchemaPreservesFactAndSourceProvenanceFields)$' -count=1 -v
```

Result: exit `0`.

- Resolution: projection regression and orphan-local-export fail-closed tests PASS.
- FileContext: all five canonical/access cases PASS, including canonical false not inheriting public visibility, canonical true independent of private visibility, legacy fallbacks, and malformed canonical value fail-closed.
- Ladybug load: Export CSV fields, definition compatibility carrier, loader rows, and File -> Export pair PASS.
- Ladybug schema: full Export fact/source-provenance column contract and File -> Export relation pair PASS.

### Full affected-package regression — PASS

Command:

```text
go test ./internal/resolution ./internal/filecontext ./internal/lbugload ./internal/lbugschema -count=1
```

Result: exit `0`; all four packages PASS (`0.284s`, `0.951s`, `0.845s`, `0.890s`).

## Behavior interpretation

- Direct exported TypeAlias source membership now yields definition `isExported=true` at the production resolver boundary.
- The TypeAlias Export fact remains explicitly type-only/type-meaning; no runtime-value field or terminal resolution state was added.
- Function positives use the same validated local-membership path as before and remain true; no Function-specific code changed.
- Definitions without a local Export fact remain absent from the membership map. The FileContext canonical-false and fail-closed cases pass, supporting preservation of all target negative-control semantics.
- Access visibility is unchanged and independent from compatibility.
- FileContext consumes the corrected canonical value without production edits.
- Ladybug CSV/schema/loader owners remain unchanged and their nearest parity contracts pass.
- Orphan local-definition references fail closed, and the projection test finds no forbidden Child 05 field.
- The exact six target Functions and exact 11 owner-qualified target controls require fresh QA comparison; this is the next ordered gate, not a residual Coder-scope gap.

## E2E Verification

```text
E2E Verification:
  [PASS] Compiled: npm run full-build -> exit 0; CLI 1.2.8; final analyze 1,917/736/0
  [PASS] Runtime boundary: resolver graph -> FileContext canonical reader -> Ladybug CSV/schema/loader
  [PASS] Happy path: Function and direct type-only TypeAlias definitions -> isExported=true
  [PASS] Edge cases: typeOnly/meaning preserved; access separation; orphan fail-closed; canonical false/malformed false; no Child 05 state
```

No benchmark row is created: this repair changes no benchmarkable performance/capacity system. Build and test timings are validation evidence.

## Worktree and boundary preservation

- Index is empty.
- Current tracked modifications are exactly the five pre-existing Child 04 living documents plus this lane's two repair files.
- The concurrent orchestration-skill commit is preserved and untouched by this lane.
- All pre-existing untracked Oracle/QA/Supervisor/Architect/Planner/Main provenance remains preserved.
- The historical `BLOCKED` report is retained unchanged; this `READY_FOR_QA` report is the only new durable artifact created by the resume.
- The disposable `.tmp/p4c2_lock_monitor.ps1` remains absent.
- Existing empty `.tmp/p4c-tests` parent remains untouched; it contains no evidence artifact.
- No edit occurred in FileContext, Ladybug CSV/schema/loader, ScopeIR/provider owners, other tests/goldens, plan/ledger documents, target files, or Child 05/later-slice owners.
- No target analyze/comparison, QA, Supervisor, Anvien detect-changes, stage, commit, push, reset, checkout, or broad cleanup occurred.
- No evidence-bearing artifact originated in `.tmp`.

## Residual scope and handoff

Residual unverified surfaces within the authorized Coder repair boundary: **none**.

Status: **READY_FOR_QA**.

Main should route the existing sealed P4-C2 oracle and immutable QA workflow to fresh target validation. QA must prove exact `21/21` positives, including all 15 TypeAliases and six Functions; exact `11/11` negative controls; FileContext exported results; Graph JSON/Ladybug record-level parity; zero duplicate/orphan/diagnostic/forbidden Child 05 state; and target/config preservation. Coder does not authorize or perform that external gate.

## Report identity

- Encoding/line ending: UTF-8 / LF.
- Canonical SHA-256 basis: replace the 64-character value after `canonical SHA-256:` with 64 ASCII zeroes.
- Bytes / LF / canonical SHA-256: `14218` / `215` / canonical SHA-256: `2B3A9B04801DE8F76D79017172EA8E251B4A2975F1772E0A7BD1E278DE1997F6`.
