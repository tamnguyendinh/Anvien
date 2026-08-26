# Child 06A A002 Cheapapp Corrected 17-Child Measurement

Status: `MEASUREMENT_INCOMPLETE_BUILD_FAILURE_MISSING_LBUG_HEADER`.

This is measurement data only for `E:\cheapapp.org`. It makes no `KEEP`, `REWORK`, `ROLLBACK`, Supervisor, disposition, checklist, streak, or baseline-promotion decision.

## Outcome

The single authorized build attempt failed before a candidate executable was produced. Per the no-retry stop rule, no second build and no analyze launch occurred. Consequently there is no corrected A002 runtime measurement, no D001 value, no parent/analyzer/process candidate value, and no 30/17 candidate packet to compare with accepted A001.

| Field | Result |
|---|---|
| Build attempt count | `1` |
| Build start | `2026-08-25T22:16:04.2598316+07:00` |
| Build end | `2026-08-25T22:18:55.9468223+07:00` |
| Build wall | `171.675682500 s` |
| Build exit | `1` |
| Candidate binary | Not produced |
| Analyze launch count | `0` |
| Benchmark JSON | Not produced |
| CPU profile | Not produced |
| Exact blocker | `internal\lbugnative\native_ladybugdb.go:6:10: fatal error: lbug.h: No such file or directory` |

Exact compiler output:

```text
# github.com/flamingoosesoftwareinc/tree-sitter-swift/bindings/go
In file included from vendor\github.com\flamingoosesoftwareinc\tree-sitter-swift\bindings\go\binding.go:6:
vendor\github.com\flamingoosesoftwareinc\tree-sitter-swift/src/scanner.c:131:9: warning: left shift count >= width of type [-Wshift-count-overflow]
  131 |     1UL << FAKE_TRY_BANG, // BANG,
      |         ^~
# github.com/tamnguyendinh/anvien/internal/lbugnative
internal\lbugnative\native_ladybugdb.go:6:10: fatal error: lbug.h: No such file or directory
    6 | #include "lbug.h"
      |          ^~~~~~~~
compilation terminated.
```

## Startup and identity evidence

All pre-build gates passed:

| Gate | Exact evidence |
|---|---|
| `anvien.exe --help` | PASS, exit `0` |
| Anvien HEAD | `cc420e7ad719d90dc4b2d9991be0249e8d648daa` |
| Staged path count | `0` |
| Cheapapp HEAD before and after build failure | `a869876ab6262dacde6cd5d432d099a91852a646` |
| Corrected root before creation | Absent |
| Competing Anvien analyze/benchmark process before build | `0` |
| Competing Go/compiler/link process before build | `0` |
| Go | `C:\Program Files\Go\bin\go.exe`; `go version go1.26.3 windows/amd64` |
| Overlay manifest | `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\optimized-overlay.json`; SHA-256 `7B138FBA06B41CBD4C6709E6DF00C80C03E4015B615D1E42F787DE2A65D378F7`; exactly two mappings |
| Overlay `resolve.go` | `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\optimized-overlay\internal\resolution\resolve.go`; SHA-256 `92A89227E1B1B1C159DE8BE77A1F060361C9058C4F83C78BFC32E9AB40DADEA9` |
| Overlay `types.go` | `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\optimized-overlay\internal\resolution\types.go`; SHA-256 `A8DDCB1C61C312D3D3B164B1DD02BFD91C640F1FEB8B6DB511B73FD923C802B8` |
| Current unmapped `diagnostics.go` | SHA-256 `AEABA8A541D1C293DFA4FE411253A85A26B6AC0DD18047AFA6F9151B4932BEE8` |
| Current unmapped `emit.go` | SHA-256 `73AE20E8A18AC25F8019E2A3223407621AE7ABC358DBAF4F058E1D9AC3136060` |
| Current unmapped `outcome.go` | SHA-256 `02092F9FE7DA2A4BDB49E13056FBA3C97DC24F416141E7F27866EE80F60C1F7E` |
| Current unmapped `diagnostics_test.go` | SHA-256 `58E8CAE2C4EBBF4672D2D338CC5E1419B7CECA9F7389B0E24F51BF2F975A7CEA` |
| Pinned native DLL authority | `E:\Anvien\third_party\ladybugdb\v0.19.1\windows-x86_64\lbug_shared.dll`; 20,230,656 bytes; SHA-256 `20CBD87840483A2053CFF3FC2DB23A86DD802B8915D86509D41A4B709624CDB7` |

The target pre/post HEAD and dirty status were unchanged. The build never invoked or wrote `E:\cheapapp.org`, so target source and its existing target-local graph were not mutated by this attempt.

## Exact build command and environment

Working directory: `E:\Anvien`.

```powershell
& 'C:\Program Files\Go\bin\go.exe' build -buildvcs=false -mod=vendor -tags ladybugdb -trimpath '-ldflags=-s -w' -overlay 'E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\optimized-overlay.json' -o 'E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\candidate\anvien-a002-cheapapp-17child.exe' .\cmd\anvien
```

The build wrapper explicitly set:

```text
HOME=E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\home
USERPROFILE=E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\home
TEMP=E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\temp
TMP=E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\temp
APPDATA=E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\appdata
LOCALAPPDATA=E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\localappdata
XDG_CACHE_HOME=E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\xdg-cache
XDG_CONFIG_HOME=E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\xdg-config
XDG_DATA_HOME=E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\xdg-data
GOCACHE=E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\gocache
GOMODCACHE=E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\gomodcache
GOPATH=E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\gopath
GOPROXY=off
GOSUMDB=off
CGO_ENABLED=1
PATH_PREFIX=E:\Anvien\third_party\ladybugdb\v0.19.1\windows-x86_64
```

The supplemental environment correction arrived after the build started and before it ended. The active build was intentionally not interrupted or restarted. During this sole build, `TMPDIR`, `GOENV`, `GOWORK`, `GOFLAGS`, `GOTELEMETRY`, `GOTELEMETRYDIR`, `GOTOOLCHAIN`, `GOPRIVATE`, `GONOPROXY`, `GONOSUMDB`, `GOINSECURE`, `GOVCS`, `CGO_CFLAGS`, and `CGO_LDFLAGS` were inherited as unset/null. The corrected runtime contract was therefore never reached.

All controlled build temp, cache, home, application-data, XDG, Go cache, module cache, and GOPATH locations were under the assigned `E:\Anvien\.tmp` root. Installed Go on `C:` was used only as the authorized executable.

## Corrected child contract not executed

Had the build passed, validation would have required the authoritative raw instrumentation order:

```text
build_binding_occurrence_index
runtime_setup
emit_definition_nodes
emit_unresolved_heritage_diagnostics
emit_import_edges
emit_heritage_compatibility_edges
resolve_calls
resolve_accesses
resolve_type_annotations
emit_method_dispatch_edges
finalize_typescript_authority_results
emit_typescript_external_symbols
finalize_resolution_outcomes
project_resolution_outcomes
finalize_resolution_metadata
assemble_resolution_result
binding_accumulator_dispose
```

Each row would also have required `childId`, `durationNs`, `denominators`, `denominatorExpression`, `sourceOwner`, `sourceStart`, `sourceEnd`, `callPath`, `intervalStart`, `intervalEnd`, and `overlapGroup=exclusive_resolution_parent`, with interval-sum equality, exactly 2,675 intervals, zero overlap, and nonnegative parent residual. None of these corrected candidate fields is claimed because analyze launch count is `0`.

## Raw evidence and handoff

| Artifact | Path |
|---|---|
| Corrected raw root | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child` |
| Machine-readable build failure | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\build-result.json` |
| Candidate directory | `E:\Anvien\.tmp\child06a_a002_cheapapp_benchmark_corrected_17child\candidate` |
| Corrected report | `E:\Anvien\reports\Investigation\rp_child06a_a002_cheapapp_benchmark_corrected_17child.md` |

Accepted A001 authority remains unchanged at `E:\Anvien\.tmp\child06a_a001_cheapapp_benchmark\owner-final-run\optimized`. No A001 measurement was rerun. No source, test, plan, ledger, plan-rules, existing report, target source, or Restaurant Manager artifact was edited. No retry, analyze, Supervisor, disposition, stage, commit, push, reset, checkout, stash, clean, install, or network action occurred.
