# Supervisor Report: Child 05 P5-C Ambiguous Owner Member Resubmission

Verdict: `PASS`

## Metadata

- Review time: `260821 223035` Asia/Bangkok (`UTC+07`).
- Reviewer: `gpt-5` Supervisor lane.
- Authoritative checkout: `E:\Anvien` only.
- Branch/HEAD: `master` / `861000cb6b6e36ce105623f0dc8c093b089f61fa`.
- Accepted predecessor: P5-B `c1559df953a277b099009f8489576d00ed25aa58`.
- Scope reviewed: only the previously rejected distinct-owner ambiguity invariant on semantic member composition; no full P5-C audit restart.
- Prior authority: `E:\Anvien\reports\Supervisor\rp_supervisor_260821_220516_by_gpt-5_child05_p5c_export_resolution_reject.md`, independently verified as `8,918` bytes / `81` LF / `0` CR / no BOM / SHA-256 `1B8E32DA433782116A70BE427962A12DD799FF26B431BF7982A2D32FCE3121F1`.
- Resubmission: `E:\Anvien\reports\coder\rp_coder_260821_222334_by_gpt-5_child05_p5c_ambiguous_owner_member_resubmission.md`, independently verified as `12,702` bytes / `236` LF / `0` CR / no BOM / SHA-256 `14A9ADB407527F12CF5F27484FF1E8C3E731AF174573D63D6895CDC607E4CC5F`; creation and last-write timestamps are identical.

## Claim reviewed

The repair must keep an ambiguous export owner non-terminal during member lookup, retain one owned terminal or missing-member proof for every distinct owner branch, make `definition()` and `resolveImportedMember` fail closed, emit no `CALLS`/`ACCESSES` to a sole surviving member branch, and introduce no physical/global fallback or scope drift.

## Decision

The rejected invariant is closed on current source and final bytes. The repair composes members per owner candidate, records the formerly dropped missing-owner branch, preserves owner-level ambiguity after aggregation, and leaves all previously cleared consumers and fallback helpers byte-identical. No required follow-up remains inside this reject-only review scope.

## Source-level clearance

### `E:\Anvien\internal\resolution\export_resolution.go` — clear

- Final identity: `26,015` bytes / `780` LF / `0` CR / no BOM / SHA-256 `566A69B965212D94E87E4D46703FC4E4C80E910FDC8657449D3914D5B868248F`.
- `resolveSemanticImportedMember` records owner ambiguity at lines `225-226` instead of discarding the aggregate outcome.
- Lines `228-265` preserve owner failures and iterate each candidate/proof branch directly; the rejected `ownerResult.allProofs()` flattening and shared `ownerProducedMember` suppression are absent.
- A definition owner without the requested member now contributes `missingOwnedMemberProof` at lines `247-249`.
- `missingOwnedMemberProof` at lines `283-300` deep-clones inherited proof data, clears terminal state, records the owner file/name, and appends a source-backed member hop with exact `MemberOwnerDefID`.
- Lines `272-275` aggregate all member/failure branches deterministically and retain `ambiguity` when the owner lookup was ambiguous. A sole surviving member candidate therefore remains provenance only and is not selectable.

### `E:\Anvien\internal\resolution\indexes.go` — inherited clearance preserved and consumer rechecked

- Identity remains exactly `A44C94FA545FFBAEF7017D49C6529B1B4CFDE82C3BB244EC6F4DD57EBE4283F6` (`33,417` bytes / `1,213` LF).
- Lines `639-644` call the semantic member resolver and return false inside the `handled` branch when `definition()` fails. The legacy physical scan at lines `646+` is therefore unreachable for this semantic ambiguity.

### `E:\Anvien\internal\resolution\resolve.go` — inherited clearance preserved and emission paths rechecked

- Identity remains exactly `047CCDAE8F451553FC159CCF5B9864FFCA60F597E25FF31BDF5494638EC5817C` (`20,799` bytes / `668` LF).
- Member calls with a non-empty explicit receiver use `resolveImportedMember` at lines `419-429`; the repository-global fallback at lines `431-440` is restricted to an empty receiver and cannot rescue this case.
- Explicit-receiver access uses the same fail-closed member consumer at lines `552-565` and has no physical/global fallback after failure.

Preserve-only P5-B/path/emit identities also remain exact: `export_tables.go` `BA4C7F9F2D05F6715CA22195C1CFF64C1AACADC3D81E7CA3ED711C8EA8B63E19`; `export_tables_test.go` `A0395EF9030CB054B09C52AFBE048D2F8244BAD8A924EE492BB7AB312CD08AD8`; `import_resolution.go` `67C5F10E40E2DCD5850C56622A4D0ED3CBDFE15D84A9419971E97F8799876413`; `emit.go` `B1B8F22AEFE94F72483CD284B7E267FE3897412C04D581AFAB7F9FE30F0B2867`.

## Focused regression clearance

`E:\Anvien\internal\resolution\export_resolution_test.go` final identity is `29,480` bytes / `612` LF / `0` CR / no BOM / SHA-256 `97FF4990066179AE779C9354D7A59859F3CA437826062DA0DA19C54776BC1620`.

`TestP5CAmbiguousOwnerMemberPreservesNoSelectionAndEveryBranch` at lines `271-358` models exactly the rejected path: barrel stars reach distinct owners A/B; only A owns `run`. It asserts:

- aggregate member outcome is `ambiguity`;
- the sole terminal member candidate remains as provenance;
- `definition()` and production `resolveImportedMember` select nothing;
- A contributes a terminal member proof and B contributes an owned missing-member proof;
- runtime resolution emits neither `CALLS` nor `ACCESSES` to A's member;
- both source sites remain unresolved.

Independent final-byte execution after an E-only free-lock check:

```text
go -C E:\Anvien test ./internal/resolution -run '^TestP5CAmbiguousOwnerMemberPreservesNoSelectionAndEveryBranch$' -count=1 -v
exit 0; package 0.227s

go -C E:\Anvien test ./internal/resolution -run '^TestP5C' -count=1 -v
exit 0; all seven top-level tests and four cycle/dedupe subtests; package 0.163s
```

## Invariant closure

1. Owner ambiguity remains non-terminal/no-selection: closed by `ownerAmbiguous` retention plus outcome enforcement after aggregation.
2. Every distinct owner branch retains terminal or missing-member provenance: closed by per-candidate/per-proof composition and `missingOwnedMemberProof`.
3. `definition()` and `resolveImportedMember` fail closed: closed by outcome gating at `export_resolution.go:84-92` and handled semantic return at `indexes.go:639-644`.
4. No `CALLS`/`ACCESSES` reaches the sole member branch: closed by the focused production-boundary fixture and zero resolved-call/access metrics.
5. No physical/global fallback or scope drift: closed by current consumer source, byte-identical cleared owners, exact two-file repair manifest, empty staging, and clean formatting/diff checks.

Residual unverified same-invariant surfaces: none.

## Build, output, and preservation evidence

- The canonical build was not repeated because there was no post-build source/test invalidation. The immutable report records the exact command `powershell -ExecutionPolicy Bypass -File E:\Anvien\scripts\full-build.ps1` from `E:\Anvien`, exit `0`, `1,966 / 740 / 0`, graph `115,902 / 159,581`.
- Current output bytes independently match that final run: `E:\Anvien\anvien\bin\anvien.exe` `71,478,272` bytes / SHA-256 `A55D9CE575EAD60EE07612B882448B69FC0272712B5AEE08EF2971C4577A4E62`; `lbug_shared.dll` SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7`; graph `462,380,776` bytes / SHA-256 `0A16E835901EB134023CC85C14F7900BBAC1B7DCD7E15E8F329FAB59227847B2`.
- Source/test last writes precede `anvien.exe` and `graph.json`; independent tests did not modify them.
- Accepted fixed-corpus artifact remains `9,386` bytes / SHA-256 `7687C5197AE76DC74294FB948AB9E25124BCEF26E444230EB29A2771EC373796` with `736` parsed code and `5,072 / 5,072 / 5,088` authority.
- Raw current graph inspection independently counted `5,177` `IMPORTS`; exact P5-B growth remains `8+12+20=40` and P5-C growth `8+21+20=49`, so `5,177-(40+49)=5,088`. The repair adds no import statement, path lookup, table construction, or emit change; fixed-corpus delta remains `0 / 0 / 0`.

## Evidence not run

- No repeated full build or analyze: current final-byte hashes and chronology were intact, and the user prohibited repetition absent concrete invalidation.
- No independent full `internal/resolution` package run: the immutable report records exit `0` on final bytes; the Supervisor independently ran the exact rejected fixture and complete P5-C matrix.
- No analyze/file-detail/impact restart, detect-changes, stage, commit, push/reset/checkout, cleanup, target/P5-D action, or C-worktree access.

## Boundary and final state

The only Supervisor write is this report. Source, tests, four ledgers, prior reports, target, and P5-D were not edited or opened. P5-C remains unchecked until Main records this acceptance and performs its separately authorized detect/commit transition.

State: `IDLE/PASS`.
