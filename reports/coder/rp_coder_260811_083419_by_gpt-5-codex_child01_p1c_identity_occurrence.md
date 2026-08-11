# Coder Report: Child 01 P1-C Identity and Occurrence Conservation

Status: READY FOR ZERO-TRUST REVIEW; NOT ACCEPTED OR COMMITTED

## Metadata

- Report file: `reports/coder/rp_coder_260811_083419_by_gpt-5-codex_child01_p1c_identity_occurrence.md`
- Report time: `2026-08-11 08:34:19 +07:00`
- Coder: `gpt-5-codex`
- Repository: `E:\Anvien`
- Git base: `6b3a3fc4480c7f4c2291d8004207f50b52d91d21`
- Scope: Child 01 `P1-C` only — deterministic definition occurrence identity and occurrence conservation.
- Authority: `AGENTS.md`, the user work rules, the active graph-accuracy contract/roadmap, and the Child 01 plan/evidence/benchmark/actual-status ledgers.
- SPEC note: `Docs/SPEC` does not exist in this repository; no SPEC was inferred or created.

## Invariant Family Map

- Family: provider/ScopeIR definition occurrence -> lexical membership -> graph identity -> definition node and endpoints.
- Authority / source truth: provider `DefinitionFact.ID`, `ScopeFact.OwnedDefIDs`, production `defsByFile` denominator, `Label`, `OwnerID`, and the accepted P1-B inputs.
- Production owner changed: `internal/resolution/indexes.go`, limited to `buildWorkspace`, `graphIDForDef`, and the tuple encoder serving that identity responsibility.
- Sibling surfaces checked: all production definition providers, `emitDefinitionNodes`, graph mutation, relationship merge, standalone `BuildDefinitionIndex`, resolution callers, provider heritage parity, and the 61 product packages.
- Preserve-only siblings: provider extraction, `internal/resolution/emit.go`, `Graph.AddNode`, `AddRelationship`, persistence/readers, Child 03 binding patterns, and `E:\cheapapp.org`.
- Forbidden fallback: name-only/range-only identity, random or absolute-path inputs, traversal-order dependence, hash/topology/package invention, implicit declaration merge, test compatibility fallback, or broad producer rewrite.

## Implementation

- `internal/resolution/indexes.go:143` passes the existing `scopeByDef[DefinitionFact.ID]` lexical membership into identity construction.
- `internal/resolution/indexes.go:814` builds identity from normalized repository-relative file, semantic/qualified name, optional arity, provider occurrence ID, lexical scope ID, and owner ID; the graph label remains the meaning prefix.
- `internal/resolution/indexes.go:832` length-prefixes each field so tuple boundaries are unambiguous. Nil arity and zero arity remain distinct.
- No additional production file, runtime topology, persistence format, reader, relationship policy, or collision API was changed.

## QA and Stale Contract Alignment

- New direct behavior test: `internal/resolution/identity_occurrence_test.go`.
- New independently authored fixture: `internal/resolution/testdata/p1c_identity_repo/src/identity.ts`.
- Direct cases prove: two `time`, two `now`, same-lexeme Function/Interface meaning lanes, lexical/owner sensitivity, nil/zero arity separation, path normalization, definition-order determinism, no unsupported merge, `defsByFile` conservation, and endpoint existence.
- Existing relationship tests that encoded old literal graph IDs now resolve the intended semantic definition nodes. They still test their original resolution/heritage behavior; P1-C injectivity remains in the dedicated identity test.
- Updated QA files: `internal/resolution/resolution_test.go`, five legacy resolution test files, and `internal/providers/provider_parity_test.go`.

## Anvien Evidence and Blast Radius

- Fresh pre-edit production impact: `graphIDForDef` CRITICAL `23` symbols / `13` files / `3` modules / `18` processes; `buildWorkspace` CRITICAL `41/17/7/28`.
- Inspect-only projection impact: `emitDefinitionNodes` CRITICAL `24/9/6/35`; `emitter.emitNode` CRITICAL `6/5/2/19`.
- Production file-detail: `internal/resolution/indexes.go` has `46` related files; inspect-only `emit.go` has `42`.
- Resolution QA test functions changed after production were LOW with `0` upstream impact; `buildGraphSignature` was LOW `2` symbols / `2` files / `1` test module / `0` processes.
- Provider parity QA file-detail: `117` related files, low file risk; exact changed test function LOW with `0` upstream symbols/processes.
- CRITICAL/HIGH values were treated as blast-radius warnings and kept to the exact plan owner.

## Verification Outputs

- `[PASS]` `go test ./internal/resolution -run '^TestP1CIdentity' -count=1 -v` — four P1-C tests.
- `[PASS]` `go test ./internal/resolution -count=1` — complete resolution package.
- `[PASS]` `powershell -NoProfile -ExecutionPolicy Bypass -File .\scripts\full-build.ps1` — packaged Go/Ladybug runtime, web, launcher/server, CLI `1.2.8`, and built-CLI self-analyze.
- `[PASS]` built CLI fixture analyze: `1/1` parsed, `0` failed, `20` nodes, `19` relationships.
- `[PASS]` runtime graph oracle: `10/10` definitions, `10/10` unique definition IDs, `10/10` `DEFINES`, `time 2/2`, `now 2/2`, separate `Shared` meaning lanes, `0` missing endpoints.
- `[PASS]` tracked/runtime fixture SHA-256 equality: `327F6BD02FCB341078902FE4302CAB03A2077DDFE98B379164C17E96B547A078`.
- `[PASS]` provider heritage focused matrix: 12 representative language cases.
- `[PASS]` `go test ./internal/providers -count=1`.
- `[PASS]` product regression: `61/61` packages from `go list -e ./...`, excluding only the measured 11 non-standalone `*/anvien/test/fixtures/*` corpus packages.
- `[PASS]` `git diff --check`.

## Failed Attempts and Closure

- The first resolution-package regression exposed only stale literal-ID assertions. After fresh QA impact, they were aligned to semantic endpoints; the package passed.
- One unchanged package-test retry was terminated by a 30-second shell timeout and produced no assertion result; the same command with a sufficient timeout passed.
- The first 61-package product regression exposed one remaining stale literal-ID assertion in provider heritage parity while the relationships were present. After fresh file-detail/impact and QA-only correction, the focused provider matrix, provider package, and exact 61-package rerun passed.

## E2E Verification

E2E Verification:
  [PASS] Compiled: `scripts/full-build.ps1` -> full repository build and built self-analyze exited `0`.
  [PASS] Runtime: packaged `anvien\bin\anvien.exe analyze <repo-local fixture> --force --json` -> real graph artifact produced.
  [PASS] Happy path: ScopeIR definitions -> workspace identity -> graph nodes/`DEFINES` -> `10/10` conserved.
  [PASS] Edge case: duplicate names, different lexical owners/meanings, reordered facts, nil/zero arity, and unsupported merge -> distinct deterministic results.

## Sibling Surfaces Checked

- Provider occurrence IDs and normalized repository-relative file paths: preserved.
- Projection/endpoints: source unchanged; runtime and direct tests show no affected missing endpoints.
- Relationship merge: unchanged and regression-covered.
- `Graph.AddNode` collision policy: unchanged; P1-D remains closed.
- Persistence/readers: unchanged; owned by Child 02.
- Target repository: not accessed or changed; target validation belongs to P1-E.

## Legacy Fallback Status

No production or test compatibility fallback was added. Stale tests resolve actual semantic nodes and assert real relationships; they do not accept both old and new IDs.

## Residual Unverified Surfaces

None inside the P1-C behavior boundary. P1-D collision policy, P1-E target acceptance, and Child 02 persistence/readers are explicitly separate closed scopes, not residual P1-C work.

## Review Gate

- Required next action: independent zero-trust Supervisor review of authority, source, diff, runtime artifact, tests, ledgers, and this report.
- Not yet run: final Supervisor verdict, post-PASS Anvien detect-changes, and isolated P1-C commit.
- P1-C must not be marked complete and P1-D must not open before those gates pass.
