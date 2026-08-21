# Supervisor Report: Child 05 P5-A Requested Import Meanings

Verdict: PASS

## Metadata

- Report file: `E:\Anvien\reports\Supervisor\rp_supervisor_260821_190000_by_gpt-5_child05_p5a_requested_meanings.md`
- Review time: `260821 190000 +07:00` (Asia/Bangkok)
- Reviewer: `gpt-5`
- Repo/project: `E:\Anvien` (authoritative checkout)
- Scope reviewed: uncommitted Child 05 P5-A candidate at HEAD `0aa49c87628c9e8b2041754515d6ebf0a930d55b`
- Claim reviewed: P5-A adds the canonical TS/JS requested-meaning and type-only import-input contract, owns it through ScopeIR normalization/order, and preserves module/path, re-export SSOT, other-language, and graph-count boundaries.
- Authority used: latest Main authorization, `AGENTS.md`, `working-rules` and `supervisor` skills, R2 plan, current R3 evidence/benchmark/actual-status ledgers, accepted Child 04 closure, current source, canonical E build addendum, and fresh E graph artifacts.
- Related artifacts:
  - Coder candidate: `reports/coder/rp_coder_260821_170956_by_gpt-5_child05_p5a_requested_meanings_ready_for_supervisor.md` — 15,237 bytes / 283 LF / SHA-256 `C2599B890BB75D0783DFAB6F643F42522C0D1E3C163E490BAE5A287B6F5C7968`.
  - Canonical E build addendum: `reports/coder/rp_coder_260821_1811_by_gpt-5_child05_p5a_canonical_e_full_build.md` — SHA-256 `77BC1E402AACD54BEDAF32EDBDB6458BF7CC0778E35C11842CFFAA7BD7887840`.
  - Fresh E analyze artifact: `E:\Anvien\.tmp\p5a-supervisor-e-analyze.json` — 9,378 bytes / 365 LF / SHA-256 `293458E24AF38982FC6E6B43B8D2539F4C2AA159861B71F16A0DB763EA20E0BD`.

## Executive Summary

- Problem: determine whether the sole open P5-A candidate is safe and complete enough for acceptance after the earlier invalid C:/X: build provenance was superseded.
- Decision: PASS. Source, scope, canonical E build provenance, focused and boundary regressions, unaffected-language classification, preservation hashes, and the three independent import denominators close the requested-input invariant.
- Required outcome: accepted for the Supervisor gate only. `E5-P5A-DETECT1` and `E5-P5A-COMMIT1` remain pending; P5-B/P5-C/P5-D, target access, and global-rescue work remain locked.

## Source-Level Clearance Notes

- `internal/scopeir/facts.go:189-206`: clear. `ImportFact` adds only `RequestedMeanings []ExportMeaning` and `TypeOnly bool`; dormant `Target*` fields remain present and untouched in meaning.
- `internal/providers/tsjs/imports.go:26-105,1327-1367`: clear. Source-written TS/JS forms populate the new fields; compatibility re-export/wildcard calls pass `nil,false`; side-effect-only statements still return without a fact.
- `internal/scopeir/ir.go:35-115,166-194`: clear. `Normalized` and `NormalizeOwned` deep-clone requested-meaning slices; `NormalizeInPlace` sorts and deduplicates them while retaining `TypeOnly`.
- `internal/scopeir/sort_keys.go:39-61`: clear. `compareImport` compares requested meanings and `TypeOnly` before existing target/ID tie-breakers.
- Focused test owners `internal/providers/tsjs/extract_test.go:92-170` and `internal/scopeir/scopeir_test.go:88-160`: clear. Tests cover the syntax matrix, side-effect/re-export boundaries, clone/canonicalization, deterministic ordering, and round-trip behavior.

The complete changed Go manifest is exactly six authorized files (the four production owners above plus the two focused test owners); no production or test path outside that manifest is changed. The preserve-only hashes are unchanged: `internal/resolution/indexes.go` = `AA19B9D543012309A90974089BACBD0A122594C7481FFB0790DE5C01F3D3D76B`; `internal/resolution/import_resolution.go` = `67C5F10E40E2DCD5850C56622A4D0ED3CBDFE15D84A9419971E97F8799876413`.

## Form-Matrix Verification

| Source form | Verified result |
|---|---|
| Plain default, named, or alias | exact imported/local/alias names; canonical `RequestedMeanings={value,type,namespace}`; `TypeOnly=false` |
| Statement-level type-only | `RequestedMeanings={type}`; `TypeOnly=true` |
| Inline type-only specifier | `RequestedMeanings={type}`; `TypeOnly=true` |
| Plain namespace | `ImportedName` is empty (no exported-name request); `RequestedMeanings={namespace}`; `TypeOnly=false` |
| Type-only namespace | `ImportedName` is empty; `RequestedMeanings={type}`; `TypeOnly=true` |
| Side-effect-only | no `ImportFact` emitted |
| Compatibility re-export imports | requested fields remain empty and `TypeOnly=false`; accepted `ExportFact` remains the sole re-export semantic SSOT |
| Non-TS/JS facts | source writers do not assign either new field, so they remain empty/false |

## Evidence Checked

Passed:

- Exact current candidate hashes:
  - `internal/scopeir/facts.go` — `4548FF2C312F329B0672A5CF5999F112F647CBAC6C060E489B08D6B7C1063646`
  - `internal/providers/tsjs/imports.go` — `538FD9723AB77C672DD35F7B5D17D01C565167757FA73FFCF23DE05E0A1640D4`
  - `internal/scopeir/ir.go` — `4BCE48A4E490810707708C0D199C3B837899347F19DD084DB6142F59DE7D6ED6`
  - `internal/scopeir/sort_keys.go` — `741F83729AB1CED8E09945A74779C10E18BD4BF526EA0AA03A228236FFEB06B2`
  - `internal/providers/tsjs/extract_test.go` — `CCA89257D3E428D39DB8BDB3E61136ABACE135A8F20D3B800FC5563B01CE6D59`
  - `internal/scopeir/scopeir_test.go` — `AC757B1FB75E529CAF29299620E80C0CB46144B2327C574D5B62AFD378249964`
- Fresh E graph artifact: scanned `1,949`, parsed code `736`, failed `0`; graph `114,842` nodes / `157,744` relationships; resolution `ImportsResolved=5,072`; `FinalizedImportsEmitted=5,072`.
- Direct persisted graph inspection of `E:\Anvien\.anvien\graph.json`: `IMPORTS=5,088`.
- Three-count proof on the same parsed-code corpus: baseline/post-change `5,072 / 5,072 / 5,088`, deltas `0 / 0 / 0` (physical target-file resolutions / resolver-emitted syntactic `IMPORTS` / persisted graph-wide `IMPORTS`).
- Canonical build evidence is E-bound: `powershell -ExecutionPolicy Bypass -File .\scripts\full-build.ps1`, cwd `E:\Anvien`, exit `0`, canonical outputs `E:\Anvien\anvien\bin\anvien.exe` and `E:\Anvien\anvien\bin\lbug_shared.dll`, final graph at `E:\Anvien\.anvien\graph.json`; no X:, `subst`, alias, copied output, or C-worktree provenance is used for this verdict.
- Verified output identities: `anvien.exe` 71,376,896 bytes, SHA-256 `4EE58BFB1B3252D9A12CF11153D79DCBF1328832D445DCE4CCE28992A470C63D`; `lbug_shared.dll` 20,230,656 bytes, SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`; runtime metadata reports source `E:/Anvien`, Windows AMD64, version `1.2.8`.
- Focused requested-meaning tests and full `internal/providers/tsjs` plus `internal/scopeir` pass; nearest real non-UI `internal/resolution`, `internal/analyze`, and `internal/cli` boundary passes; aggregate provider results contain only the already accepted out-of-slice C#/Dart `ACCESSES` parity baselines, while focused C#/Dart extractor tests pass.
- Blast-radius warnings were reviewed and scoped: `facts.go` CRITICAL (`867` symbols / `123` files), `imports.go` HIGH (`18` / `1`), `ir.go` CRITICAL (`610` / `76`), `sort_keys.go` CRITICAL (`32` / `5`), `scopeir.ImportFact` CRITICAL (`624` symbols / `73` files / `25` modules / `67` processes), and `ScopeIR.Normalized` CRITICAL (`21` / `6` / `4` / `15`). No warning was ignored; no preserve-only owner was edited.
- Current ledger identities at review: plan `719DD0CF1CA5442CA206409BBC6352C812EF524A006968A890648751907E2A1E`; evidence `030DD29C05F108223DDFBD687B17EC3E903D44B31D14405A97D49B39788F721A`; benchmark `D9F5982C914AE5E15E4F66C393D8BC58F206C5F1B6EF76042BBF310DF3DE1BB9`; actual-status `E2DC5F9F1FEDD7594AD0FBFA0D4F320EF3BABBD03F888635829ED24AD3F8F4EC`. The plan remains R2; P5-A is not self-ticked, and REVIEW/DETECT/COMMIT plus P5-B+ remain pending/locked.

Failed:

- None within the authorized P5-A invariant.

Not run:

- Full build was not rerun during this review because Main explicitly required use of the canonical E build evidence and prohibited another build. The canonical E script result and output identities were independently checked.
- `anvien detect-changes` and commit were intentionally not run; they are post-Supervisor workflow gates and remain pending by authority.
- Target `E:\cheapapp.org`, P5-B+, traversal/table/global-rescue surfaces, and resolver/path owners were intentionally not accessed or changed.

## Invariant Closure

- affected invariant: source-written TS/JS module-request semantics before export lookup, including requested meaning/type-only state and deterministic ScopeIR ownership.
- sibling surfaces checked: accepted Child 04 `ExportFact` SSOT, compatibility re-export imports, side-effect-only handling, all ScopeIR normalization paths, deterministic import ordering, every non-TS/JS production import writer, resolver/analyze/CLI boundary, unaffected C#/Dart extractors, and all three graph denominators.
- residual unverified same-invariant surfaces: none inside authorized P5-A. Export-table construction, re-export traversal, terminal binding, target proof, detect, and commit are separate locked gates, not residual P5-A obligations.

## Overall Evaluation

The candidate is source-correct, scope-bounded, and backed by authoritative E checkout build/runtime and fresh graph/count evidence. The previous C:/X: provenance is explicitly excluded; the canonical E script/output pair closes that process boundary. PASS applies only to P5-A Supervisor acceptance. Main may open the detect/commit workflow after this report; no later slice is implicitly opened.
