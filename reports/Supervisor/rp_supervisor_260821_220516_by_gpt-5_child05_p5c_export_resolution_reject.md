# Child 05 P5-C Supervisor Review

- Review mode: zero-trust, review-only, authoritative checkout `E:\Anvien`.
- Authority: branch `master`, HEAD `861000cb6b6e36ce105623f0dc8c093b089f61fa`.
- Accepted predecessor: P5-B `c1559df953a277b099009f8489576d00ed25aa58` (confirmed ancestor).
- Candidate claim: resolve aliases, re-exports, stars, cycles, ambiguity, meanings, namespace/member, and explicit-import no-global rescue.
- Coder handoff: `E:\Anvien\reports\coder\rp_coder_260821_213311_by_gpt-5_child05_p5c_export_resolution_ready_for_supervisor.md`.
- Handoff identity independently verified: `15,694` bytes / `232` LF / `0` CR / UTF-8 BOM absent / SHA-256 `3AA88C36D8AF4E839BADB31B4388EAB183EEE80ED1D70E6F56311C5019B4D28D`.

## Verdict

`REJECT`

P5-C does not yet preserve distinct-owner ambiguity through the semantic member-resolution surface. The defect is source-backed and reachable through the production `resolveImportedMember` consumer even though the current focused suite passes.

## Blocking invariant

Rejected invariant: **distinct-terminal ambiguity must select no definition, including namespace/member lookup, and every branch must retain a complete terminal or failure proof**.

Source evidence in `E:\Anvien\internal\resolution\export_resolution.go`:

1. `resolveSemanticImportedMember` obtains `ownerResult` at lines `219-223`, but does not gate on `ownerResult.Outcome`.
2. Line `224` calls `ownerResult.allProofs()`. That helper intentionally exposes terminal proofs from every candidate even when the aggregate owner result is `ambiguity`.
3. Lines `225-266` resolve a member independently from each ambiguous owner proof.
4. `ownerProducedMember` and `ownerHadTerminal` are shared across the whole import. Lines `267-273` add one missing-member failure only when **no** owner produced a member.
5. Therefore, for an import whose requested owner name resolves through stars to two distinct terminal owners A and B, where A has member `run` and B does not, A contributes one terminal member proof, B contributes no proof, and the missing branch is suppressed because `ownerProducedMember` is already true. `aggregateExportResolution` then sees one candidate and returns `terminal`; `definition()` selects A's member.
6. `E:\Anvien\internal\resolution\indexes.go:639-644` accepts that terminal from `resolveSemanticImportedMember`, so `resolveImportedMember` returns it. The production call/access consumers can consequently emit a relationship from an explicitly ambiguous receiver.

This is not merely a test-coverage concern. It converts a distinct-terminal owner ambiguity into a selected member terminal and drops source-backed failure provenance for the other owner. It violates the P5-C no-selection rule and the complete-proof requirement on the same authorized namespace/member surface.

## Source-level clearance

### `E:\Anvien\internal\resolution\export_resolution.go`

Not cleared. The core traversal itself has deterministic outcomes, explicit-over-star behavior, default exclusion, active-key cycle termination, meaning intersection, same-terminal proof dedupe, distinct direct-terminal ambiguity, and owned cloning. The blocking member composition at lines `219-273`, however, does not preserve an ambiguous owner result and does not record per-owner missing-member outcomes.

Final candidate identity checked: `25,380` bytes / `764` LF / `0` CR / no BOM / SHA-256 `5494CAE992429F4E2599B51FD9E2B21A39D0AD3128FDC19A337E541580BE9F04`.

### `E:\Anvien\internal\resolution\indexes.go`

Cleared for its bounded orchestration changes: phase 1 resolves each import's file candidates once, phase 2 builds the accepted P5-B tables once, and phase 3 resolves terminals and creates bindings without another path lookup. The redundant post-`resolveImports` table build is removed. The semantic-member dispatch correctly fail-closes against the legacy physical scan, but it consumes the flawed member result described above and therefore cannot clear the end-to-end invariant by itself.

Final candidate identity checked: `33,417` bytes / `1,213` LF / `0` CR / no BOM / SHA-256 `A44C94FA545FFBAEF7017D49C6529B1B4CFDE82C3BB244EC6F4DD57EBE4283F6`.

### `E:\Anvien\internal\resolution\resolve.go`

Cleared for the authorized seam. The tracked diff is limited to `resolveCall`; it gates the repository-global fallback for semantic explicit imports while permitting a distinct inner/local shadow. Generic `resolveName`, `resolveGlobalName`, and `resolveGlobalCallName` remain unchanged. No separate no-global blocker was found.

Final candidate identity checked: `20,799` bytes / `668` LF / `0` CR / no BOM / SHA-256 `047CCDAE8F451553FC159CCF5B9864FFCA60F597E25FF31BDF5494638EC5817C`.

## Test owner and coverage

`E:\Anvien\internal\resolution\export_resolution_test.go` identity checked: `25,477` bytes / `523` LF / `0` CR / no BOM / SHA-256 `09797F5B6AAC7F424AA380C8E9F758F08FE3F45CFD7C53DC6CEA4E317674E82E`.

The suite covers direct/alias identity, explicit-over-star, no star default, same-terminal dedupe, direct distinct-terminal ambiguity, pure/terminal cycles, meaning mismatch, unambiguous namespace/member lookup, no-global rescue, syntactic `IMPORTS`, and one non-semantic Go regression. It does not exercise member lookup whose owner result is already ambiguous, so the blocking branch is not contradicted by the green suite.

## Evidence checked

- Full `AGENTS.md`, working-rules skill, Supervisor skill, four living ledgers, immutable Coder report, all four candidate files, tracked diff, and accepted P5-B table owner were read.
- Branch/HEAD and P5-B ancestry matched authority. P5-C and P5-D remain unchecked; staging is empty.
- Tracked numstat matched exactly: `indexes.go` `43+/14-`; `resolve.go` `22+/2-`. `git diff --check` was clean.
- Preserve-only identities matched: `export_tables.go` `BA4C7F9F2D05F6715CA22195C1CFF64C1AACADC3D81E7CA3ED711C8EA8B63E19`; `export_tables_test.go` `A0395EF9030CB054B09C52AFBE048D2F8244BAD8A924EE492BB7AB312CD08AD8`; `import_resolution.go` `67C5F10E40E2DCD5850C56622A4D0ED3CBDFE15D84A9419971E97F8799876413`; `emit.go` `B1B8F22AEFE94F72483CD284B7E267FE3897412C04D581AFAB7F9FE30F0B2867`.
- Focused independent replay after an E-only free-lock check: `go -C E:\Anvien test ./internal/resolution -run '^TestP5C' -count=1 -v`; exit `0`, package `0.188s`.
- Canonical final-byte build evidence was checked, not repeated: literal `powershell -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`, cwd `E:\Anvien`, exit `0`; final source/test writes precede `anvien.exe` and `graph.json`.
- Current canonical outputs matched the handoff: `anvien.exe` SHA-256 `A3AFBA37318E4762E4F7045FB650A457F7A7ADE1ED4BE78C0D8A157841742A37`; `lbug_shared.dll` `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`; graph `ACB222619F0E15C888FEB22D4F6382A27B784E253F831679A90E6B751C5BDFE5`, with `115,800 / 159,445` recorded by the canonical run.
- Fixed-corpus artifact matched SHA-256 `7687C5197AE76DC74294FB948AB9E25124BCEF26E444230EB29A2771EC373796` and records `5,072 / 5,072` on `736` parsed-code files. Current graph directly contains `5,177` `IMPORTS`; exact graph counts independently matched P5-B `8+12+20=40` and P5-C `8+21+20=49`, supporting `5,177-(40+49)=5,088` and `5,161-89=5,072`.

## Evidence not run

- No repeated full build or analyze: final-byte output identities and chronology were intact, and there was no pre-review invalidation requiring a duplicate gate.
- No independent full resolution-package replay: the focused same-byte replay was run; the immutable handoff records the full package run.
- No `detect-changes`, stage, commit, push, reset, checkout, cleanup, target access, P5-D action, or C-worktree access.

## Required re-review evidence

1. Correct the member composition so an owner-level `ambiguity` remains non-terminal/no-selection and every owner branch contributes owned terminal/failure provenance; do not repair it with physical or global lookup.
2. Add a focused fixture with two distinct owner terminals reached through the accepted export tables, only one of which supplies the requested member. Assert the semantic member result is `ambiguity`, `definition()` returns false, `resolveImportedMember` returns false, no call/access relationship is emitted to the sole branch member, and both owner branches remain represented in proof/failure evidence.
3. Re-run the complete P5-C focused matrix and the full `internal/resolution` package on final bytes.
4. Because production/test bytes must change, provide one fresh exact canonical E-only build using `powershell -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1`, with source/output chronology, canonical output hashes, fixed-corpus preservation, exact candidate identities/diff, and a new immutable Coder resubmission report.

State after verdict: `IDLE/REJECT`. P5-C remains unchecked; detect/commit and P5-D remain locked.
