# Child 03 P3-C — Exact owner-scope repair

Created: 2026-08-20 12:35:48 +07:00  
Role: Coder repair candidate; this report does not self-accept P3-C.  
Review history: `E3-P3C-REVIEW1` remains the immutable `REJECT`.  
Repository / branch / HEAD: `E:\Anvien` / `master` / `a36b42c8c71d2fa7de551b5bd00a8c026123f8c6`.

## Handoff summary

The sole REVIEW1 blocker is repaired inside the authorized two-file boundary. Binding-occurrence validation now proves, before graph selection or mutation, that the unique owner agrees with the leaf and definition file, that its range contains both declaration full ranges and every present selection range, and that the exact matching local Binding belongs to that validated owner scope.

The focused negative control moves both `OwnedDefID` and the matching local Binding together to a module scope in a different file. `ResolveInto` returns the expected file-drift error, returns no graph, and leaves a one-node sentinel graph byte-for-byte and count-for-count unchanged.

All well-formed and persistence invariants cleared by REVIEW1 remain green. P3-C2, target validation, detect, stage, commit, push, and ledger work were not opened.

## Invariant Family Map

| Item | Exact repair boundary |
| --- | --- |
| Family | accepted binding leaf -> exact Variable definition -> unique lexical owner/local Binding -> graph mutation |
| Authority / SSOT | active Child 03 plan/evidence/benchmark/actual-status, roadmap, graph-accuracy contract, and immutable `reports/Supervisor/rp_supervisor_260815_204754_by_gpt-5_e3_p3c_review1.md` |
| Rejected invariant | Exact ownership; gap integrity is its direct downstream consequence |
| Production owner | `internal/resolution/resolve.go` |
| Direct test owner | `internal/resolution/p3c_binding_occurrence_test.go` |
| Locked sibling | `internal/lbugload/p3c_binding_occurrence_persistence_test.go`, SHA-256 `C704B45FD350F2A1B064D79E78B4DC99F6378D44358B4E86149A09FA38D4A850` |
| Sibling surfaces checked | resolution, Ladybug load parity, graph, Ladybug schema, TSJS extraction, ScopeIR; assignment/import conservation and lexical shadowing remain covered by the focused P3-C matrix |
| Forbidden fallback | global/same-name rescue, inferred binding, owner/local-Binding coordination bypass, mutation before complete validation |
| Observable failure oracle | wrong-file owner plus matching local Binding -> error, nil result graph, sentinel JSON bytes and node/relationship counts unchanged |
| Residual unverified surfaces | none inside the REVIEW1 repair; P3-C2 is intentionally locked and out of scope |

## Source and impact evidence

The durable excluded-tree graph gate was preserved rather than rerun: `1105` scanned / `624` parsed / `0` failed; `80,118` nodes / `119,054` relationships.

The accepted exact impact gate was also preserved:

- `buildBindingOccurrenceIndex`: `CRITICAL`, 6 impacted symbols / 4 files / 33 processes.
- `validateBindingOccurrenceOwner`: `CRITICAL`, 4 impacted symbols / 2 files / 22 processes.
- `TestP3CBindingOccurrenceProjectionFailsClosedOnOrphanOrDrift`: `LOW`, 0 impacted symbols/files/processes.

HIGH/CRITICAL are blast-radius warnings. The repair stayed within the exact two-file boundary.

Final source anchors:

- `internal/resolution/resolve.go:52` — `ResolveBoundInto` builds and validates the occurrence index before selecting/creating the graph.
- `internal/resolution/resolve.go:134` — `buildBindingOccurrenceIndex` finds the unique definition/owner and then counts the local Binding only inside the validated owner scope.
- `internal/resolution/resolve.go:214` — `validateBindingOccurrenceOwner` checks normalized leaf/definition/owner-IR/owner-scope file equality and owner containment of full/selection ranges.
- `internal/resolution/p3c_binding_occurrence_test.go:125` — fail-closed matrix.
- `internal/resolution/p3c_binding_occurrence_test.go:198` — coordinated wrong-file owner/local-Binding sentinel control.

Final hashes:

| File | Lines | SHA-256 | State |
| --- | ---: | --- | --- |
| `internal/resolution/resolve.go` | 648 | `C1FF5C515D401ECAD4FBF93C271DF4AC19101B2F5410D4174F7F598502BBC96A` | repair production candidate |
| `internal/resolution/p3c_binding_occurrence_test.go` | 432 | `9BAE2F63575C313B5F5F8EF4C265360BCB38F8D2F616CDCDC8942C57A080CE7F` | focused P3-C candidate |
| `internal/lbugload/p3c_binding_occurrence_persistence_test.go` | 282 | `C704B45FD350F2A1B064D79E78B4DC99F6378D44358B4E86149A09FA38D4A850` | locked and unchanged |

`gofmt` left the focused test hash unchanged. `git diff --check` passed. Production diff relative to HEAD is `268` insertions / `11` deletions in the already authorized owner.

## Code-first and focused pre-build proof

Production last-write time precedes the test last-write time:

- production: `2026-08-15T14:27:07.3781243Z`
- focused test: `2026-08-15T14:46:47.9288287Z`

After resume and re-anchor:

```text
go test -count=1 -run '^$' ./internal/resolution
```

PASS, package compiled with no tests selected.

```text
go test -count=1 -v -run 'TestP3CBindingOccurrenceProjectionFailsClosedOnOrphanOrDrift/coordinated_wrong-file_owner_and_local_binding' ./internal/resolution
```

PASS. This confirmed the negative control before the canonical build.

## Holder-clean canonical full build

Pre-build evidence:

- analyze lock: `free`, lock file absent;
- one exact editor-owned holder tree was found: `cmd.exe` PID `12616` -> `anvien.exe mcp` PID `14560`;
- both PIDs were terminated and confirmed absent;
- recheck: `anvien doctor processes --json` returned `[]`; no Go/npm/node/esbuild/Vite build process remained.

With `TEMP`, `TMP`, `GOTMPDIR`, and `GOCACHE` rooted under `E:\Anvien\.tmp\p3c-owner-scope-repair`, the unchanged canonical command was run:

```text
powershell -ExecutionPolicy Bypass -File .\scripts\full-build.ps1
```

Result: PASS, exit `0`.

- packaged Go runtime build: PASS;
- npm install/audit: PASS, `0` vulnerabilities;
- global install and `anvien version`: PASS, `1.2.8`;
- launcher/native build: PASS;
- Web production build: PASS, `2,943` modules, Vite `38.72s`;
- canonical final analyze: PASS, `1806` scanned / `729` parsed / `0` failed.

The unmodified canonical script itself runs an unexcluded final analyze and surfaced forbidden-tree path names in stdout. No file content from those trees was opened, no path from that output was used as P3-C evidence, and no graph command was run afterward. The durable excluded-tree graph gate above remains the only graph/impact basis for this handoff.

The Swift scanner shift warning, npm allow-scripts warning, and Vite dynamic-import/chunk-size warnings were non-failing existing build output.

## Final-byte verification

Focused P3-C gate:

```text
go test -count=1 -run '^TestP3C' -v ./internal/resolution ./internal/lbugload
```

PASS:

- `TestP3CBindingOccurrencesProjectAndResolveLexically` — PASS.
- `TestP3CBindingOccurrenceProjectionFailsClosedOnOrphanOrDrift` — PASS.
- `missing_definition` — PASS.
- `missing_local_binding` — PASS.
- `duplicate_accepted_leaf` — PASS.
- `coordinated_wrong-file_owner_and_local_binding` — PASS.
- `TestP3CBindingOccurrencesPreserveGraphJSONAndLadybugLoadParity` — PASS.

Six-package regression:

```text
go test -count=1 ./internal/resolution ./internal/lbugload ./internal/graph ./internal/lbugschema ./internal/providers/tsjs ./internal/scopeir
```

PASS, all `6/6` packages.

Static gate:

```text
go vet ./internal/resolution ./internal/lbugload ./internal/graph ./internal/lbugschema ./internal/providers/tsjs ./internal/scopeir
```

PASS, exit `0`, no diagnostics.

Fresh hostile proof after vet:

```text
go test -count=1 -v -run 'TestP3CBindingOccurrenceProjectionFailsClosedOnOrphanOrDrift/coordinated_wrong-file_owner_and_local_binding' ./internal/resolution
```

PASS. The test constructs a sentinel graph with one node and metadata, marshals it before resolution, performs coordinated owner/local-Binding drift, and requires:

- a file-drift error;
- `result.Graph == nil`;
- unchanged node count;
- unchanged relationship count;
- byte-identical JSON after resolution.

## E2E Verification

```text
E2E Verification:
  [PASS] Compiled: canonical full-build command -> exit 0
  [PASS] Runtime boundary: ResolveInto hostile owner drift -> error before graph mutation
  [PASS] Happy path: well-formed 7/7 binding occurrences, lexical reads/write, shadowing/import conservation -> focused P3-C PASS
  [PASS] Edge case: coordinated wrong-file owner + local Binding -> sentinel byte/count unchanged
  [PASS] Persistence sibling: Graph JSON/Ladybug parity -> locked focused test PASS
```

E2E flow: malformed ScopeIR owner coordination -> `BuildCrossFileBinding` -> `ResolveBoundInto` -> `buildBindingOccurrenceIndex` -> exact owner validation error -> caller observes unchanged sentinel graph.

## Boundary, cleanup, and non-actions

- Files changed by the P3-C candidate remain only `internal/resolution/resolve.go` and `internal/resolution/p3c_binding_occurrence_test.go`; this report is the sole new repair artifact.
- The persistence test was read/hash-checked only and remains at the locked SHA-256.
- Governance/roadmap/four ledgers, prior reports, extraction/import/serialization/Ladybug production, target repository, P3-C2, and both forbidden skill trees were not edited.
- The scoped worktree also contains pre-existing Main-owned governance/report paths and the locked persistence test; they were not modified by this lane.
- `E:\Anvien\.tmp\p3c-owner-scope-repair` was validated as repo-local and removed after verification (`6,279` reproducible cache files and `264` cache directories); final existence check is false.
- Final analyze lock is free and final Anvien process inventory is `[]`.
- No benchmark is recorded because this repair does not change measured product performance/capacity.
- No `detect-changes`, `git add`, commit, push, merge, reset, checkout, stash, ledger edit, target access, or self-acceptance occurred.

## Sibling surfaces checked

- Well-formed occurrence conservation and seven distinct `DEFINES` endpoints: preserved by focused P3-C PASS.
- Lexical reads/write and inner/outer shadowing: preserved by focused P3-C PASS.
- Assignment non-declaration and import count conservation: preserved by focused P3-C PASS.
- Graph JSON/Ladybug load parity and zero persistence fallback: preserved by locked parity test PASS and locked hash.
- Resolution, Ladybug load, graph, Ladybug schema, TSJS, and ScopeIR packages: `6/6` regression PASS.

Legacy fallback status: none introduced; exact owner validation rejects coordinated drift before mutation and lexical resolution still uses the previously cleared validated occurrence set without name rescue.

Stale tests/helpers/plans updated: only the authorized focused fail-closed test gained the coordinated owner-drift negative control. Governance and plans are Main-owned and locked.

Residual unverified surfaces: none inside `E3-P3C-REVIEW1` repair scope. P3-C2 remains intentionally locked, not accepted, and not accessed.

Risks/open points: independent focused `E3-P3C-REVIEW2` is still required. This report makes no PASS/commit/P3-C2 claim.

READY_FOR_SUPERVISOR
