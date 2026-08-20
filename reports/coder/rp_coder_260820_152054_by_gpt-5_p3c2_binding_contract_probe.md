# Coder Report — Child 03 P3-C2 metadata-only binding contract probe

Final state: `READY_FOR_QA`

## Metadata

- Created: `2026-08-20 15:20:54 +07:00` (`Asia/Bangkok` / `SE Asia Standard Time`)
- Role: Coder lane only
- Repository/cwd: `E:\Anvien`
- Sole open slice: Child 03 `P3-C2`
- Accepted predecessor: P3-C commit `656a0445ff3e25b6225b994cdaf7cf1b35eb665c`, parent `a569b8674fefdaa757cf7fdf63f454caf7925215`
- QA handoff: `reports/QA/rp_qa_260820_143739_by_gpt-5_p3c2_target_binding_validation.md`, SHA-256 `2109A1C9451A4D910D9BDFEFD3991F8C9FB061621784D275468DA38E15C19487`
- Main authority: `reports/Investigation/rp_main_260820_144839_p3c2_blocked_handoff_verification.md`
- Next responsible person: Orchestration Main; Main independently verifies this candidate and resumes the existing QA task
- Non-claim: this report does not accept P3-C2, does not issue a Supervisor verdict, and does not open Pn-A or any later lane.

## Outcome

Created the authorized reusable metadata-only command with exactly one responsibility:

```text
caller-provided TypeScript file + ordered repeated local names
  -> production parser
  -> tsjs.Extract
  -> exact BindingLeafFact
  -> exact matching Variable DefinitionFact
  -> unique lexical owner with one OwnedDefID
  -> one owner-local BindingLocal
  -> deterministic metadata-only JSON stdout
```

The command never invokes analyze, never writes a report/output file, has no `-out`/`--out` option, does not copy a source body, and emits no JSON until every requested name passes the complete chain.

## Exact boundary

Created:

1. `cmd/binding-contract-probe/main.go`
2. `cmd/binding-contract-probe/main_test.go`
3. this timestamped report

Preserved:

- every existing production/test/governance file;
- accepted P3-C source, tests, reports, plan, roadmap, and all four living ledgers;
- existing graph probes and graphaccuracy owners;
- incoming untracked Main/QA reports;
- target repository and target source.

No `E:\cheapapp.org` command, read, stat, hash, scan, parser invocation, or write occurred in this coder lane. Final `rg` across both new cmd files returned exit `1` with zero matches for `cheapapp` and all six bounded target names.

## Invariant Family Map

| Field | Locked result |
| --- | --- |
| Family | direct metadata-only ScopeIR binding-contract observation |
| SSOT | graph-accuracy contract, Child 03 plan/ledgers, QA `BLOCKED` handoff, Main verification |
| Primary path | `-file` exactly once + ordered repeated `-name` -> parser -> `tsjs.Extract` -> validator -> JSON stdout |
| Sibling surfaces | parser, TSJS extractor, ScopeIR contracts are consume-only; analyzer, resolution, existing probes, graphaccuracy and target remain preserve-only |
| Forbidden fallback | no hard-coded names/path, name rescue, analyzer run, report/output file, source-body field, provider repair, or target access |
| Failure mode | missing, duplicate, mismatched, or ambiguous leaf/definition/owner/local-binding state returns an exact error and nonzero process exit before stdout JSON |
| Verify matrix | deterministic caller order, exact fact chain, missing name, duplicate leaf/definition/owner/OwnedDefID/local binding, source-body absence, forbidden `-out`, full build, parser/TSJS/ScopeIR regression, real CLI smoke |

Sibling surfaces checked: parser pool/registry/parse entrypoints, `tsjs.Extract`, ScopeIR `DefinitionFact`/`BindingLeafFact`/`ScopeFact`/`ScopeIR`, `BindingContextVariable`, `BindingLocal`, and scanner `TypeScript` constant.

Legacy fallback status: none added; validation matches exact range/selection/file/hash/DefID and never resolves by name alone.

Stale tests/helpers/plans updated: none; boundary forbids existing-file changes.

Residual unverified surfaces: `none` inside the coder asset boundary. Real-target execution remains deliberately unperformed and belongs to resumed P3-C2 QA.

## Authority and skill gate

Read completely before implementation:

- `AGENTS.md`
- `.agents/skills/working-rules/SKILL.md`
- `.agents/skills/planner/SKILL.md`
- `.agents/skills/coder/SKILL.md`
- `.agents/skills/backend-development/SKILL.md`
- relevant backend testing/code-quality references
- graph-accuracy roadmap
- complete Child 03 plan plus evidence, benchmark, and actual-status ledgers
- `docs/contracts/graph-accuracy-contract.md`
- the exact QA and Main reports above

Planner was used only to follow the existing Child 03 plan. Owner's preserve-only boundary prohibited a new or modified plan/ledger.

## Pre-edit Anvien evidence

`anvien --help` was the first Anvien command and exited `0`.

Fresh excluded graph command:

```powershell
anvien analyze . --force --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**"
```

Result: exit `0`; scanned/parsed/failed `1111/624/0`; graph `80204` nodes / `119147` relationships; indexed/current commit both `656a0445ff3e25b6225b994cdaf7cf1b35eb665c`; all consumed file-detail rows were `stale=false` and `changedSinceAnalyze=false`.

### Current file-detail and file impact

| Existing consume-only owner | File detail | Upstream file impact |
| --- | --- | --- |
| `internal/parser/pool.go` | `33` symbols / `114` inbound / `7` outbound / `51` local / `1` flow / `25` tests / HIGH | CRITICAL: `22` impacted symbols / `9` files / `1` flow |
| `internal/parser/registry.go` | `14/36/8/18`, `0` flows, `23` tests, HIGH | CRITICAL: `19` impacted symbols / `8` files / `1` flow |
| `internal/parser/parse.go` | `35/24/22/16`, `1` flow, `22` tests, HIGH | LOW: `3` impacted symbols / `1` file / `1` flow |
| `internal/providers/tsjs/extract.go` | `36/22/32/29`, `1` flow, `6` tests, HIGH | CRITICAL: `21` impacted symbols / `10` files / `1` flow |
| `internal/scopeir/facts.go` | `164/1057/22/148`, `97` tests, MEDIUM | CRITICAL: `328` impacted symbols / `75` files / `1` flow |
| `internal/scopeir/ir.go` | `47/1068/34/75`, `2` flows, `95` tests, HIGH | CRITICAL: `116` impacted symbols / `29` files / `1` flow |
| `internal/scanner/language.go` | `31/252/0/2`, `55` tests, MEDIUM | CRITICAL: `410` impacted symbols / `188` files / `1` flow |

### Exact canonical-UID symbol impact

| Symbol | Risk / impact |
| --- | --- |
| `parser.DefaultRegistry` | CRITICAL: `9` impacted symbols / `8` files / `13` process records |
| `parser.NewPool` | CRITICAL: `9` impacted symbols / `7` files / `23` process records |
| `(*parser.Pool).Parse` | LOW: `0/0/0` |
| `tsjs.Extract` | CRITICAL: `8` impacted symbols / `5` files / `20` process records |
| `scopeir.ScopeIR` | CRITICAL: `115` impacted symbols / `29` files / `74` process records |
| `scopeir.DefinitionFact` | CRITICAL: `188` impacted symbols / `22` modules / `68` process records |
| `scopeir.BindingLeafFact` | CRITICAL: `139` impacted symbols / `22` modules / `72` process records |
| `scopeir.ScopeFact` | CRITICAL: `178` impacted symbols / `22` modules / `68` process records |
| `scopeir.BindingContextVariable` | LOW: `0/0/0` |
| `scanner.TypeScript` | LOW: `0/0/0` |

All HIGH/CRITICAL results are blast-radius warnings. They justified a new isolated command and no edit to any existing owner.

No file-detail/impact command was attempted on either nonexistent new cmd path before creation.

## Implementation behavior

- CLI grammar: one required `-file`; one or more ordered repeatable `-name`; no positional arguments; duplicate requested names rejected.
- The path must resolve to one regular file recognized by the production scanner as TypeScript (`.ts`, `.tsx`, `.mts`, or `.cts`).
- Source is read once into memory, hashed with SHA-256, parsed through `parser.NewPool(...).Parse`, then passed to `tsjs.Extract` with the production TypeScript language contract.
- Syntax-tree errors fail closed.
- Results follow caller name order, not provider sort order.
- Each selected leaf requires variable context, valid exact range/selection, valid typed path segments, and complete provenance.
- Matching Definition requires exact name, `NodeVariable`, file/hash, range, selection, non-empty DefID, and unique DefID occurrence.
- Owner requires exactly one ScopeFact containing the DefID exactly once, same file/hash, lexical containment of leaf/definition full and selection ranges, and unique scope ID.
- Local binding requires exact `BindingLocal/name/DefID` with exact/same-name/same-def counts `1/1/1`.
- JSON is buffered and copied to stdout only after all requested names validate.
- Output includes schema version, normalized file path, language, hash, byte count, parser grammar/root kind, and only the required fact/owner metadata. It contains no source buffer/body field.

## Code-first sequence

1. Created only `main.go`.
2. `gofmt -w .\cmd\binding-contract-probe\main.go`.
3. `go test ./cmd/binding-contract-probe -run '^$' -count=1` exited `0` with `[no test files]`, proving production compile before test creation.
4. A repo-local `.tmp` synthetic smoke proved caller order `betaRows,alphaRows`, typed indexes `1,0`, Variable definitions, Function owners, owner/local counts `1/1`, and absent source-body sentinel. Intentional missing-name run returned process exit `1` and no JSON.
5. Only then created `main_test.go`.
6. Final code/test bytes were gofmt-clean before the holder gate and full build.

## Holder/lock gate

Before build:

- `anvien doctor locks --repo E:\Anvien --json`: exit `0`, status `free`, lock absent.
- `anvien doctor processes --json`: exit `0`; identified five installed `anvien.exe ... mcp` children and five `cmd.exe` parents owned by Codex.
- Exact child PIDs `12648,14032,9684,10088,12924` were revalidated by process name/executable path and terminated. Their five `cmd.exe` parents exited with them; the declared 10-PID holder set had `REMAINING=` empty.
- Second doctor process result: `[]`; second lock result: `free`.
- Exclusive-open checks passed `6/6` for repo/global `anvien-runtime.json`, `anvien.exe`, and `lbug_shared.dll`.

This proves the build replacement boundary was holder-clean before the unchanged full build started.

## Full build

Command:

```powershell
npm run full-build
```

Result: exit `0`.

- packaged Go runtime build PASS;
- npm audit `0 vulnerabilities`;
- global install PASS; blocked-install-script message was a non-failing npm warning;
- `anvien version` `1.2.8`;
- launcher/native build PASS;
- Web production build PASS with `2943` modules; dynamic-import/chunk-size messages were non-failing Vite warnings;
- script-final unexcluded analyze PASS at `1814/731/0`, graph `111430/153398`.

The script-final graph is build output only and is not Child 03 graph evidence. No Anvien graph command was run after the build.

## Post-build validation

### Focused command tests

```powershell
go test ./cmd/binding-contract-probe -count=1 -v
```

Exit `0`, package `PASS` in `0.206s`:

- `TestRunProducesDeterministicMetadataInCallerOrder` PASS;
- `TestRunFailsClosedForMissingName` PASS;
- `TestValidateBindingContractsFailsClosedOnDuplicateAndAmbiguousState` PASS with five PASS subcases: duplicate leaf, duplicate matching definition, ambiguous owner, duplicate owner DefID, duplicate owner-local binding;
- `TestRunDoesNotEmitSourceBody` PASS;
- `TestCommandOptionsRejectDuplicateNamesAndOutputFlag` PASS.

All filesystem test data was independently authored and created only under `E:\Anvien\.tmp\binding-contract-probe-tests`; every test cleanup passed.

### Nearest consume-only regression

```powershell
go test ./internal/parser ./internal/providers/tsjs ./internal/scopeir -count=1
```

Exit `0`:

- `internal/parser` PASS `0.481s`;
- `internal/providers/tsjs` PASS `0.246s`;
- `internal/scopeir` PASS `0.597s`.

### Static gate

```powershell
go vet ./cmd/binding-contract-probe
```

Exit `0`, no diagnostics.

### Real command boundary after full build

Actual `go run ./cmd/binding-contract-probe` against an independently authored repo-local `.tmp` TypeScript file exited `0` and returned:

- schema `binding-contract-probe.v1`;
- caller order `rightRows,leftRows`;
- typed indexes `1,0`;
- definition labels `Variable,Variable`;
- owner kinds `Function,Function`;
- OwnedDefID occurrences `1,1`;
- local Binding occurrences `1,1`;
- source-body sentinel present: `false`.

The direct missing-name command returned process exit `1` with precise error `expected exactly one variable-context BindingLeafFact, found 0` and emitted no JSON.

E2E Verification:

```text
[PASS] Compiled: code-first compile and canonical full build
[PASS] Runtime: real CLI invocation against repo-local synthetic TypeScript
[PASS] Happy path: argv -> file -> parser -> tsjs.Extract -> exact chain -> metadata JSON
[PASS] Edge cases: missing and duplicate/ambiguous facts/owners/local bindings fail closed
```

## Failure/classification ledger

| Event | Classification | Resolution |
| --- | --- | --- |
| Raw `file-detail internal/parser/parse.go --json` display was truncated by the output renderer after the required summary, canonical UID, and relationship counts were already present | read-only presentation truncation; command exit `0`; no graph/state defect | retained the complete required evidence already returned; did not repeat or broaden discovery |
| Batched `anvien impact file internal/scopeir/facts.go ... --json` did not return a final exit within the batch yield | read-only invocation/timing issue; no state change and no usable evidence claim | same exact command ran once separately: exit `0`, CRITICAL `328` impacted symbols / `75` files |
| First empty smoke-directory cleanup containing `Remove-Item` was rejected before PowerShell creation | execution-policy rejection; no state change | absolute path, repo-temp containment, and child count `0` were proven separately; exact non-recursive `.NET Directory.Delete` succeeded |
| Missing-name smoke commands exited `1` | intentional negative behavior evidence | precise errors and zero JSON confirmed; wrapper commands exited `0` after asserting expected nonzero child exit |

No source/test/build failure exposed a production binding defect. No repair outside the two-file boundary was needed.

## Cleanup and no-target proof

Final exact task-temp checks:

```text
E:\Anvien\.tmp\binding-contract-probe-smoke   = absent
E:\Anvien\.tmp\binding-contract-probe-boundary = absent
E:\Anvien\.tmp\binding-contract-probe-tests   = absent
```

The full build's repo-native Ladybug cache remains under `E:\Anvien\.tmp` and is not a task-created probe artifact.

Forbidden-target/name scan over both new cmd files: `rg` exit `1`, match count `0`.

Production `-out` registration scan: `rg` exit `1`, match count `0`; the only `-out` mention is the negative test proving rejection.

## Exact final bytes

| File | Bytes | Lines | SHA-256 |
| --- | ---: | ---: | --- |
| `cmd/binding-contract-probe/main.go` | `19,607` | `617` | `C68E520AF4F220AD2E9E4A9B71941EB9B816C3EAFC15D6F7B44212305BB7BBA7` |
| `cmd/binding-contract-probe/main_test.go` | `10,809` | `306` | `DA3B500FB161411A4D134DDF11935A83ECB1955F08C3087B5C78D4B4073C4F9A` |

Accepted preserved hashes:

| File | SHA-256 |
| --- | --- |
| QA report | `2109A1C9451A4D910D9BDFEFD3991F8C9FB061621784D275468DA38E15C19487` |
| Main verification report | `5F20CB434F13D7F47A9EABADD11FF3529B99FBF698D5859CAF8296F712BD50CE` |
| `internal/resolution/resolve.go` | `C1FF5C515D401ECAD4FBF93C271DF4AC19101B2F5410D4174F7F598502BBC96A` |
| `internal/resolution/p3c_binding_occurrence_test.go` | `9BAE2F63575C313B5F5F8EF4C265360BCB38F8D2F616CDCDC8942C57A080CE7F` |
| `internal/lbugload/p3c_binding_occurrence_persistence_test.go` | `C704B45FD350F2A1B064D79E78B4DC99F6378D44358B4E86149A09FA38D4A850` |

`gofmt -d` returned no output for either cmd file.

## Git boundary

- HEAD: `656a0445ff3e25b6225b994cdaf7cf1b35eb665c`
- parent: `a569b8674fefdaa757cf7fdf63f454caf7925215`
- tracked modifications: `0`
- staged paths: `0`
- new coder-owned paths: the exact two cmd files and this report only
- preserved incoming untracked paths:
  - `reports/Investigation/rp_main_260820_141241_orchestration_rotation_handoff.md`
  - `reports/Investigation/rp_main_260820_144839_p3c2_blocked_handoff_verification.md`
  - `reports/QA/rp_qa_260820_143739_by_gpt-5_p3c2_target_binding_validation.md`

No `detect-changes`, `git add`, commit, push, merge, reset, checkout, stash, amend, or self-acceptance action was performed.

## Benchmark classification

This validation asset does not change measured product/runtime performance, capacity, package/startup size, graph/DB throughput, or graph inventory semantics. It is non-benchmarkable. Build/test timings and graph inventories above are validation context only, not benchmark claims.

## Handoff

The exact two-file metadata-only probe candidate is ready for independent Main verification and resumed P3-C2 QA. QA may run the new command against its already authorized real-target boundary; this coder report does not do so and does not promote the retained persisted subset to complete P3-C2 acceptance.

READY_FOR_QA
