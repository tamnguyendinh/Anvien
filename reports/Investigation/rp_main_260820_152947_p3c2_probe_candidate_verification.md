# Main Verification — P3-C2 binding-contract probe candidate

Verdict: `PASS` for `READY_FOR_QA` handoff only. This does not accept P3-C2, does not issue `E3-P3C2-REVIEW1`, and does not open Pn-A.

## Metadata

- Review time: `2026-08-20 15:29:47 +07:00` (`Asia/Bangkok`)
- Reviewer role: Orchestration Main using the zero-trust Supervisor review guard; review/route only
- Repository / HEAD: `E:\Anvien` / `656a0445ff3e25b6225b994cdaf7cf1b35eb665c`
- Open slice: Child 03 `P3-C2` only
- Claim reviewed: the isolated metadata-only probe candidate is safe and complete enough for the existing P3-C2 QA lane to execute against its already authorized target boundary
- Coder report: `reports/coder/rp_coder_260820_152054_by_gpt-5_p3c2_binding_contract_probe.md`
- Next responsible person: existing visible P3-C2 QA task `01a01df0-aaec-76e0-8cdc-6108f793fd7f`

## Exact candidate and artifact integrity

| Artifact | Bytes | Lines | SHA-256 | Result |
| --- | ---: | ---: | --- | --- |
| `cmd/binding-contract-probe/main.go` | `19,607` | `617` | `C68E520AF4F220AD2E9E4A9B71941EB9B816C3EAFC15D6F7B44212305BB7BBA7` | exact |
| `cmd/binding-contract-probe/main_test.go` | `10,809` | `306` | `DA3B500FB161411A4D134DDF11935A83ECB1955F08C3087B5C78D4B4073C4F9A` | exact |
| coder report | `16,942` | `317` | `FC25935961463336A554503F9D5045B3F117BE5DC7ED7E61FEBF88F062927A5F` | exact |
| QA `BLOCKED` report | `19,701` | `290` | `2109A1C9451A4D910D9BDFEFD3991F8C9FB061621784D275468DA38E15C19487` | preserved |
| prior Main blocker verification | current | current | `5F20CB434F13D7F47A9EABADD11FF3529B99FBF698D5859CAF8296F712BD50CE` | preserved |

`gofmt -d` returned no output for both cmd files. No tracked or staged modification exists. The new coder-owned boundary is exactly the two cmd files plus the coder report; prior Main/QA reports remain separate untracked artifacts.

## Source-level clearance

- `main.go:29-65`: clear. `-file` is single-use and non-empty; `-name` is repeatable, ordered, trimmed, and later duplicate-checked.
- `main.go:138-159`: clear. JSON is encoded into an in-memory buffer and copied to stdout only after `probeFile` validates every requested name; failures emit no partial JSON.
- `main.go:162-192`: clear. No positional arguments or duplicate names are accepted; no `-out` flag is registered.
- `main.go:194-242`: clear. One caller-provided regular TypeScript file is read once; source is hashed in memory; output contains selected metadata only.
- `main.go:244-274`: clear. Parsing uses the production parser pool and `scanner.TypeScript`, then calls `tsjs.Extract`; syntax-tree errors fail closed.
- `main.go:276-450`: clear. Caller name order is preserved. Each result requires exactly one variable-context BindingLeaf, one exact matching Variable Definition/unique DefID, one unique owner with one OwnedDefID occurrence, and one exact owner-local BindingLocal with `1/1/1` name/DefID conservation.
- `main.go:452-617`: clear. Range/path/provenance, file/hash equality, scope containment, unique scope ID, and typed path-kind metadata are validated without global/same-name rescue.
- `main_test.go`: clear. Tests use independent synthetic TypeScript under repo-local temp, cover deterministic caller order, exact ranges/indexes/owner chain, missing name, duplicate leaf/definition/owner/OwnedDefID/local binding, source-body absence, duplicate names, and forbidden `-out`.
- Forbidden target/path scan across both cmd files returned `rg` exit `1` with zero matches for `cheapapp` and all six bounded names.

No existing provider, ScopeIR, parser, analyzer, resolution, graphaccuracy, probe, plan, ledger, or report byte was edited by the coder lane.

## Fresh Anvien evidence and blast radius

Because the canonical full build overwrote the graph with an unexcluded analyze, Main refreshed again with both exclusions on exact final bytes:

```text
anvien analyze . --force --exclude "internal/aicontext/skills/**" --exclude ".claude/skills/**"
files: scanned=1114 parsed_code=626 failed=0
graph: nodes=80707 relationships=119966
indexed/current commit: 656a0445ff3e25b6225b994cdaf7cf1b35eb665c
```

Current final-byte results:

| File | File detail | Upstream impact | Clearance |
| --- | --- | --- | --- |
| `cmd/binding-contract-probe/main.go` | `161` symbols, `18` inbound, `100` outbound, `95` local, `16` linked flows, `1` linked test, HIGH, current/non-stale | HIGH: `22` impacted symbols / `1` affected file / `1` internal flow | blast radius is confined to the new command itself; no downstream repo file |
| `cmd/binding-contract-probe/main_test.go` | `78` symbols, `0` inbound, `38` outbound, `5` local, LOW, current/non-stale | LOW: `0` impacted items/files/flows | isolated focused test owner |

HIGH is a scope warning, not a prohibition. The candidate remains the exact isolated two-file responsibility authorized by Main.

## Validation evidence checked

Passed in the coder report on the exact hashes above:

- code-first compile before test creation;
- holder-clean canonical `npm run full-build`, exit `0`;
- focused package: `5` top-level tests plus `5` hostile subcases, exit `0`;
- uncached `internal/parser`, `internal/providers/tsjs`, and `internal/scopeir` regression, exit `0`;
- `go vet ./cmd/binding-contract-probe`, exit `0`;
- real post-build CLI smoke: caller order, typed indexes, Variable Definition, Function owner, `1/1` owner/local counts, and source sentinel absent;
- missing-name real CLI run: child exit `1`, precise error, zero JSON.

Main did not rerun the accepted full build or focused tests because final source/test hashes are exact and no byte changed after those gates. This avoids duplicating worker validation. Resumed QA supplies the next independent real-target boundary.

Failed/classified in the coder report:

- one read-only renderer truncation;
- one batched impact timing issue resolved by the same exact standalone invocation;
- one blocked empty-directory cleanup invocation, followed by verified literal repo-local empty-directory cleanup;
- intentional negative probe invocations returning nonzero.

None indicates a code, target, or boundary failure.

## Git, cleanup, and non-actions

- HEAD / parent remain `656a0445ff3e25b6225b994cdaf7cf1b35eb665c` / `a569b8674fefdaa757cf7fdf63f454caf7925215`.
- Tracked modifications: `0`; staged paths: `0`.
- Probe temp paths `.tmp/binding-contract-probe-smoke`, `.tmp/binding-contract-probe-boundary`, and `.tmp/binding-contract-probe-tests` are absent.
- No target access, detect-changes, plan/ledger edit, Git staging/commit/push, Supervisor verdict, or later-slice action occurred.

## Invariant closure and route

- Closed inside the coder boundary: a reusable command can directly observe and fail-closed validate the required BindingLeaf/Definition/owner/local-binding facts without editing production or emitting source bodies.
- Residual P3-C2 surface: execute the command on the already authorized six-site target file, compare direct facts to retained persisted occurrences/endpoints/gaps, recheck target boundary, and update the single durable QA report.
- Route decision: resume the existing QA task only. Do not open a duplicate QA lane or Supervisor yet.

## Overall evaluation

The candidate is `READY_FOR_QA`. Its exact two-file scope, source behavior, final-byte blast radius, build/test evidence, cleanup, and Git boundary support resuming P3-C2 validation. This review does not pre-accept the probe as committed product or accept the P3-C2 slice; those decisions remain behind resumed QA and a separate visible Supervisor gate.
