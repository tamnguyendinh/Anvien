# Child 04 / P4-A Coder Candidate Handoff

## Candidate State

- State: `READY_FOR_SUPERVISOR`.
- This report is Coder evidence, not an acceptance verdict and not a claim that P4-A is complete.
- Next recipient: Main/orchestration, to open one independent visible Supervisor review.
- `E4-P4A-REVIEW1`, `E4-P4A-DETECT1`, and `E4-P4A-COMMIT1` remain reserved and pending.
- No detect-final, stage, commit, push, P4-B, P4-B1, P4-C, P4-C2, or Child 05 work was performed.

## Baseline and Exact Boundary

- Repository / resolved cwd: `E:\Anvien`.
- Baseline/current HEAD: `ff2467bb92f94a9c53c4de030685686700051a98`.
- Accepted predecessor: Child 03 Pn-C commit `3e25f9ca75c9cb7d59bf228e6cc9aa0b81d738b4`.
- P4-A plan gate: Child 04 P0-A commit `ff2467bb92f94a9c53c4de030685686700051a98`.
- Production editable boundary used exactly:
  - `internal/scopeir/facts.go`
  - `internal/scopeir/kinds.go`
  - `internal/scopeir/ir.go`
  - `internal/scopeir/sort_keys.go`
- Tests were changed only after production behavior was implemented:
  - `internal/scopeir/scopeir_test.go`
  - `internal/scopeir/testdata/scopeir.golden.json`
- Main-owned concurrent untracked file `reports/Investigation/rp_main_260820_211734_orchestration_rotation_handoff.md` was inspected for handoff context and preserved unchanged.
- `E:\cheapapp.org`, `internal/aicontext/skills/**`, and `.claude/skills/**` were not accessed or used as evidence.

## Invariant Family Map

| Family surface | Authority / SSOT | P4-A state | Evidence / verify target |
|---|---|---|---|
| One export syntax fact per binding/specifier | Child 04 plan Requirements and P4-A acceptance; `docs/contracts/graph-accuracy-contract.md` canonical distinctions | `ExportFact` is a source-site record; no statement-level aggregation or compatibility boolean was added | `TestExportKindsAndMeaningsRepresentSourceForms`; golden/round-trip |
| Export source form and names | Child 04 Requirements: direct/default/alias/re-export/star/namespace, source names/ranges | `ExportKind` has seven non-terminal source-form kinds; `ExportedName`, `LocalName`, `TargetRaw`, and source-side `TargetExportedName` are representable | source diff; kind matrix test |
| Meaning and type-only lane | Child 04 Requirements: value/type/namespace and dual meaning without guessing | `[]ExportMeaning` normalizes to a sorted, deduplicated canonical set; `TypeOnly` remains explicit | dual value/type and type-only matrix; deep-copy test |
| Range and source provenance | Child 04 Requirements: source range, selection range, containing statement provenance | `Range`, optional `SelectionRange`, `ExportProvenance.StatementRange`, and `SiteKind` are serialized | golden/round-trip and deep-copy test |
| Unsupported/malformed export diagnostics | Child 04 hidden-fallback rule and P4-A behavior test | `ExportDiagnosticFact` plus two stable codes are structured, sortable, serializable, and countable | diagnostic count/order assertions; golden |
| Deterministic ScopeIR boundary | P4-A `ScopeIR` collection/copy/normalize/JSON acceptance | `Exports`/`ExportDiagnostics` are cloned, normalized, total-ordered, deterministically marshaled, and normalized on unmarshal | permutation test; nil-vs-empty `TargetRaw`; golden; package tests |
| Access visibility separation | Child 04 Requirement: access visibility and module export are independent | `DefinitionFact.Visibility` was not edited; export normalization test keeps `private` unchanged | zero Visibility diff; focused assertion |
| Child 05 terminal-resolution boundary | Child 04/05 ownership rule | no `TargetFile`, `TargetDefID`, `TargetModuleScope`, `LinkStatus`, `TransitiveVia`, barrel reachability, ambiguity, cycles, or public-API state in export JSON | forbidden-field assertions; preserve-only diff count `0` |

Sibling runtime surfaces checked and preserved:

- `internal/providers/tsjs/{definitions.go,imports.go,extract.go}`: no AST extraction or result wiring was added in P4-A.
- `internal/resolution/{emit.go,indexes.go}`: no graph projection or terminal resolution edit; `DefinitionFact.Visibility` projection is unchanged.
- `internal/lbugload/csv.go` and `internal/lbugschema/schema.go`: no persistence/compatibility-field edit.
- `cmd/binding-contract-probe/main.go`: normalization consumer preserved; nearest product regression passed.
- Preserve-only production diff across those named paths: `0` lines.

Forbidden legacy/alternate paths:

- no `Visibility="exported"` writer;
- no `isExported` or other compatibility-field writer;
- no legacy `ImportFact` rewrite;
- no AST extraction;
- no terminal re-export target or barrel/public-API state.

Stale artifacts aligned:

- `internal/scopeir/scopeir_test.go` now covers export collections, diagnostics, meaning-set canonicalization, deep copy, ordering, round trip, access separation, and zero terminal state.
- `internal/scopeir/testdata/scopeir.golden.json` now contains representative direct/re-export facts and one structured unsupported diagnostic.
- TSJS provider, resolution, graph, and persistence tests remain unchanged because those behaviors belong to later slices.

## Fresh Anvien Evidence — `E4-P4A-IMPACT1`

Fresh graph command before implementation:

```text
anvien analyze . --force --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**"
```

- Exit `0`.
- Scanned / parsed code / failed: `1,126 / 626 / 0`.
- Graph: `81,034` nodes / `120,392` relationships.
- Indexed/current commit: `ff2467bb92f94a9c53c4de030685686700051a98`; stale `false`.

Fresh `file-detail` summaries:

| File | Symbols | Inbound / outbound / local | Linked flows / tests | Risk |
|---|---:|---:|---:|---|
| `internal/scopeir/facts.go` | 192 | 1,123 / 27 / 174 | 0 / 98 | MEDIUM |
| `internal/scopeir/kinds.go` | 99 | 616 / 0 / 10 | 0 / 97 | LOW |
| `internal/scopeir/ir.go` | 54 | 1,094 / 38 / 89 | 3 / 96 | HIGH |
| `internal/scopeir/sort_keys.go` | 122 | 252 / 128 / 55 | 3 / 95 | HIGH |

Fresh upstream file impact:

| File | Risk | Impacted / direct / files / flows |
|---|---|---:|
| `internal/scopeir/facts.go` | CRITICAL | `347 / 110 / 76 / 1` |
| `internal/scopeir/kinds.go` | CRITICAL | `1,358 / 145 / 182 / 1` |
| `internal/scopeir/ir.go` | CRITICAL | `124 / 55 / 30 / 1` |
| `internal/scopeir/sort_keys.go` | HIGH | `27 / 24 / 3 / 1` |

Exact symbol impacts used to constrain implementation:

- `ExportFact`: CRITICAL, impacted `121`, direct `2`, modules `24`, processes `81`.
- `ExportDiagnosticFact`: CRITICAL, impacted `121`, direct `2`, modules `24`, processes `80`.
- `ExportProvenance`: CRITICAL, impacted `60`, direct `3`, modules `20`, processes `29`.
- `ExportDiagnosticCode`: CRITICAL, impacted `57`, direct `1`, modules `20`, processes `28`.
- `ExportKind`: CRITICAL, impacted `58`, direct `2`, modules `20`, processes `29`.
- `ExportMeaning`: CRITICAL, impacted `59`, direct `3`, modules `20`, processes `29`.
- `ScopeIR.Normalized`: CRITICAL, impacted `5`, direct `3`, modules `3`, processes `15`.
- `ScopeIR.NormalizeInPlace`: LOW, impacted/direct/modules/processes `1/1/1/0`.
- `ScopeIR.NormalizeOwned`: LOW, `0/0/0/0`.
- `ScopeIR.MarshalDeterministic`: LOW, `0/0/0/0`.
- `cloneExportCollections`: CRITICAL, impacted `6`, direct `2`, modules `3`, processes `14`.
- `cloneStringPointer`: CRITICAL, impacted `6`, direct `1`, modules `2`, processes `7`.
- `compareExport`: LOW, impacted `2`, direct `1`, modules `1`, processes `2`.
- `compareExportDiagnostic`: LOW, impacted `2`, direct `1`, modules `1`, processes `2`.
- `compareExportMeanings`: HIGH, impacted `3`, direct `1`, modules `1`, processes `3`.
- `compareExportProvenance`: HIGH, impacted `4`, direct `2`, modules `1`, processes `3`.

CRITICAL/HIGH results were treated as blast-radius warnings. They justified full build, focused serialization tests, and named product regression while keeping edits inside the approved six-file boundary.

## Production Source — `E4-P4A-SRC1`

`internal/scopeir/facts.go`:

- Adds `ExportProvenance` with containing statement range and exact source-site kind.
- Adds one immutable-value `ExportFact` shape per binding/specifier with file/hash, kind, names, optional local DefID, raw source module text, source-side target name, canonical meaning set, `TypeOnly`, range, selection range, and provenance.
- Comments explicitly forbid interpreting `LocalDefID`, `TargetRaw`, or `TargetExportedName` as terminal re-export targets.
- Adds `ExportDiagnosticCode`, `ExportDiagnosticUnsupportedSyntax`, `ExportDiagnosticMalformedSyntax`, and `ExportDiagnosticFact`.

`internal/scopeir/kinds.go`:

- Adds source-form kinds `direct`, `named`, `default`, `alias`, `reexport`, `star`, and `namespace` with non-overlapping source-form comments.
- Adds `value`, `type`, and `namespace` meaning lanes; dual meaning is represented by one canonical set.

`internal/scopeir/ir.go`:

- Adds `Exports` and `ExportDiagnostics` to `ScopeIR` JSON.
- `Normalized()` and `NormalizeOwned()` clone the new outer collections plus nested meanings, selection-range pointer, and raw-module string pointer.
- `NormalizeInPlace()` sorts and deduplicates meanings, then sorts exports and diagnostics.
- Existing `MarshalDeterministic()` and `Unmarshal()` now carry the new collections through the normalized boundary.

`internal/scopeir/sort_keys.go`:

- Adds total ordering over every serialized export/diagnostic field.
- Distinguishes `TargetRaw=nil` from `TargetRaw=&""`; those encode different provenance and different JSON.
- Distinguishes nil/non-nil meaning slices before lexicographic content comparison.

No production owner outside these four files changed.

## Tests and Golden — `E4-P4A-TEST1`

Direct new tests in `internal/scopeir/scopeir_test.go`:

- line 213: `TestExportCollectionsAreDeepCopiedByNormalization` — canonical meaning set; deep-copy of meanings/selection/raw target for `Normalized` and `NormalizeOwned`; structured diagnostic retained; `DefinitionFact.Visibility="private"` unchanged.
- line 291: `TestExportCollectionsMarshalDeterministically` — reversed input equivalence; diagnostic ordering/count; nil-versus-empty module provenance; no source mutation; JSON round trip; absence of terminal/resolution/barrel/public-API fields.
- line 380: `TestExportKindsAndMeaningsRepresentSourceForms` — seven kind values, exactly one fact per supplied site, dual value/type meaning, explicit type-only re-export state.

Existing boundary tests now cover the new fields:

- line 13: `TestMarshalDeterministicMatchesGolden`.
- line 27: `TestUnmarshalNormalizesScopeIR`.
- line 66: `TestNormalizeInPlaceMatchesNormalized`.
- line 77: `TestNormalizeOwnedMatchesNormalized`.

Focused command after full build:

```text
go test -v ./internal/scopeir -run '^(TestMarshalDeterministicMatchesGolden|TestUnmarshalNormalizesScopeIR|TestNormalizeInPlaceMatchesNormalized|TestNormalizeOwnedMatchesNormalized|TestExportCollectionsAreDeepCopiedByNormalization|TestExportCollectionsMarshalDeterministically|TestExportKindsAndMeaningsRepresentSourceForms)$' -count=1
```

- Exit `0`.
- `7/7` named tests PASS.

Full package command after full build:

```text
go test ./internal/scopeir -count=1
```

- Exit `0`; package PASS in `0.557s`.

## Pre-Build Holder Gate and Full Build — `E4-P4A-BUILD1`

Before the canonical build:

- `anvien doctor locks --repo E:\Anvien --json`: `status=free`, `exists=false`, `alive=false`.
- Four editor-owned global `anvien.exe ... mcp` processes held the packaged executable; all four were stopped.
- Independent recount: `buildRelatedHolderCount=0`; no `node`, `npm`, `pnpm`, `go`, `esbuild`, `tsc`, `vite`, `webpack`, `rollup`, `full-build.ps1`, launcher build, or global `anvien.exe` holder remained.

Canonical final-byte build:

```text
npm run full-build
```

- Exit `0`.
- Packaged/global CLI version: `1.2.8`.
- Launcher/Web production build: PASS; `2,943` modules transformed; Vite build `53.72s`.
- npm allow-scripts, mixed dynamic/static import, and large-chunk messages were non-failing warnings.
- The script's final unexcluded analyze exited `0`, but its graph counts are deliberately not used as Child 04 semantic evidence because that canonical script includes `internal/aicontext/skills/**`. Only the full-build exit/result is used here.
- Post-build holder count: `0`.
- The interrupted outgoing-Main build is not reused as evidence.

## Nearest Product Regression — `E4-P4A-BOUNDARY1`

Clean nearest product command after full build:

```text
go test ./internal/scopeir ./internal/providers ./internal/providers/tsjs ./internal/resolution ./internal/analyze ./cmd/binding-contract-probe -count=1
```

- Exit `0`.
- PASS: `internal/scopeir`, `internal/providers`, `internal/providers/tsjs`, `internal/resolution`, `internal/analyze`, `cmd/binding-contract-probe`.
- This verifies serialization/normalization plus the inspect-only provider, resolver/index, analyze, and existing normalization consumer boundaries without adding extraction/projection behavior.

Retained non-PASS history:

```text
go test ./... -count=1
```

- Exit `1`; not used as PASS evidence.
- Compile/setup-negative fixture packages: `anvien/test/fixtures/lang-resolution/go-map-range`, `go-method-enrichment`, `sample-code`, `go-make-builtin`, and `go-type-assertion`.
- Out-of-slice parity mismatches: `internal/providers/csharp::TestResolveCSharpGraphParityCounts` expected `ACCESSES:2` but got none; `internal/providers/dart::TestResolveDartGraphParityCounts` expected `ACCESSES:2` but got `1`.
- In that failed command, `internal/providers/tsjs`, `internal/resolution`, and `internal/scopeir` passed, but those partial results are not used as closure proof; the clean nearest product command above is the evidence.

## Real Boundary Result and Measurements

Observable boundary:

```text
source-form ExportFact values
  -> ScopeIR Normalized / NormalizeOwned
  -> deterministic JSON
  -> Unmarshal normalization
  -> same fact/diagnostic counts and canonical fields
```

- Kind matrix represented: `7/7` (`direct`, `named`, `default`, `alias`, `reexport`, `star`, `namespace`).
- Meaning lanes represented: `3/3` (`value`, `type`, `namespace`).
- Dual value/type fact: `1/1` retained as canonical two-member set.
- Type-only state: `1/1` retained in the focused matrix.
- One-fact-per-input-site matrix: `7/7`; ratio `1.0`.
- Structured diagnostic classes represented: `2/2` stable codes; deterministic test carries `2/2` diagnostic instances; golden carries `1/1` unsupported instance.
- Nested mutable export values tested: `3/3` (`Meanings`, `SelectionRange`, `TargetRaw`) deep-copied by both owned clone paths.
- Access visibility fields changed by export contract creation: `0`.
- Terminal/resolution/barrel/public-API fields present in export JSON: `0/7` forbidden names.
- This slice changes no measured product/runtime performance, capacity, graph/DB throughput, or package/startup size. Build/test timings are validation evidence, not benchmark results.

## Exact Files and Hashes

| Path | Git blob | SHA-256 | Lines | Bytes |
|---|---|---|---:|---:|
| `internal/scopeir/facts.go` | `472c84c171f5ded85315ed3923f307d82334feba` | `7F2B51D878F1541995AA884C438E18F1B1E6C72E20597E2C171FB36A59BFCA6A` | 272 | 11,518 |
| `internal/scopeir/kinds.go` | `91ed38e36ee257b16014705021ce72fa2a568c68` | `2D80ECABB63D07BAFDED161E000BFC1A412BC5B8CB242B29726271688E5D2878` | 153 | 5,167 |
| `internal/scopeir/ir.go` | `d062b6d3a68ad92628165a2122a4e4c44835450c` | `732EE7F8959F077FED5550962A5369A020F6C9EC5ABDAF4540054C00E46E728C` | 237 | 9,259 |
| `internal/scopeir/sort_keys.go` | `db16c7565490259fd00a5893fd92f4b8359063e5` | `5C155B4C151D8E11833015376C26979C50928425A169CABAB475A65F52A52DB5` | 446 | 12,528 |
| `internal/scopeir/scopeir_test.go` | `ea4cde217e8ddb98aa3b23ad13c542284e087f5b` | `ADC375EBB590F28FF72347A9B16EF0A09B971941D8C22F3F0789E877E58E6D04` | 788 | 27,670 |
| `internal/scopeir/testdata/scopeir.golden.json` | `1163a34804efad92710bfc2ea487151424c6bde8` | `C92682219EB9BD6F518B28913C64276A2E6C44B1162215DB959EF3CD5254B306` | 307 | 6,740 |

`git diff --check -- internal/scopeir`: exit `0`.

## Residual Unverified Surfaces and Handoff

- Residual unverified surfaces inside the assigned P4-A invariant family: `none`.
- Intentionally unimplemented/locked surfaces are not residual P4-A gaps: AST extraction (P4-B/P4-B1), graph/persistence/compatibility projection (P4-C), real-target `21/21` validation (P4-C2), and terminal re-export resolution (Child 05).
- Known repository-wide unrelated failures are retained above and were not changed or hidden.
- Coder-side blocker: none known inside the authorized P4-A boundary.
- Required next action: Main opens independent Supervisor review against this exact report/worktree, then owns ledger refresh, excluded graph/change detection, exact staging, isolated commit, and no-push closure only after Supervisor PASS.
