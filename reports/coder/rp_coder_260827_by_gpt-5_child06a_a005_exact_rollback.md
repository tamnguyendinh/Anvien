# Child 06A A005 Exact Rollback

## Starting identities

- `internal/resolution/outcome.go`: `18203DFAB9A227B526F8F7478B516AE6673F635BABC02D9463975E428A3983AF`
- `internal/resolution/resolve.go`: `76B7B62A060B36EE2438E76689E858544358AD681DB42F6A4FC47D271F1749A1`
- `internal/resolution/p6c3_structured_outcome_test.go`: `E561280B3F8420D2288001179431A37FC39E6E2B8564863C9AD644EEE005A2E6`
- `internal/resolution/outcome_serialization_test.go`: `89B168B2764B9A9B8EACDAB37A071BE0E1C51A8F791FE158B111651E9640C957`

## Exact reversal

- `outcome.go`: removed the A005 collector sidecar and initialization, restored duplicate-path marshaling, removed record-time sidecar storage, removed the finalized bundle/coverage checks, restored `finalize` to return `[]ResolutionOutcome`, and restored projection-time marshaling with its local encoded map.
- `resolve.go`: restored direct `outcomes` carriage in both error and success results.
- `p6c3_structured_outcome_test.go`: restored direct finalized-slice access and the direct `[]ResolutionOutcome` projection call.
- Deleted the new A005-only `outcome_serialization_test.go`.

## Final identities and gates

- `internal/resolution/outcome.go`: `02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E`
- `internal/resolution/resolve.go`: `8CEEDBA1883314EE8883320D3647C25DEF6F19F043D57881A893FBA73BA210D9`
- `internal/resolution/p6c3_structured_outcome_test.go`: `6AB9F10B004FC5292C16F8CAECBC8673BC6DC20721BCA6C1A6C118DBD2DFD1FA`
- `internal/resolution/outcome_serialization_test.go`: absent.
- Four-path `git diff --exit-code --`: PASS, empty.
- Four-path `git diff --check --`: PASS, empty.
- Staged set: PASS, empty.

## Preserved boundaries and handoff

All other source/test bytes, A001-A004/WAL/P1 bytes, plans, ledgers, reports, target and `.anvien` state, frozen/measurement packets, and user/protected artifacts were preserved. No build, test, target analyze, benchmark, profile, graph refresh, impact, detect-changes, staging, commit, or cleanup was run.

Next owner: Main Orchestration to record `A005_ROLLBACK_COMPLETE` and open the fresh A006 Architect.
