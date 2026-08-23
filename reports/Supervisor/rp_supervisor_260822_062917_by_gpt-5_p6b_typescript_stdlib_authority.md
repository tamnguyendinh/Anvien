# Supervisor Report: P6-B TypeScript Standard-Library Authority

Verdict: REJECT

## Metadata

- Report file: `reports/Supervisor/rp_supervisor_260822_062917_by_gpt-5_p6b_typescript_stdlib_authority.md`
- Review time: `2026-08-22 06:29:17 +07:00`
- Reviewer: `gpt-5`
- Repo/project: `E:\Anvien`
- Scope reviewed: only Child 06 P6-B “TypeScript standard-library authority”; P6-C1/P6-C2/P6-C3/P6-D were not opened or reviewed.
- Claim reviewed: the exact 25-path P6-B candidate is complete, conforms to the committed P6-A authority, has no remaining same-scope work, and is ready for independent acceptance before Main stages or commits it.
- Candidate anchor: branch `master`, HEAD `b98131e44932a7bcac17b487ecb2914535927d01`, parent `ec765debff335540c77d409ebb2c9f45e4a0a77d`.
- Candidate manifest before this report: 8 tracked modifications + 16 new P6-B source/test/fixture assets + immutable coder report; index empty; `git diff --check` clean.
- Protected boundary: exactly 19 `reports/Investigation/rp_main_*` handoffs, including `rp_main_260822_061831_orchestration_rotation_handoff.md`. Their pathnames alone were observed through status; none was read, modified, removed, staged, or committed. They are not candidate drift.
- Authority used: Owner review instruction; full `AGENTS.md` and `CLAUDE.md`; full `working-rules` and `supervisor` skills; all four Child 06 ledgers; immutable P6-A architecture report; both P6-A Supervisor reports; current source/diff/runtime/graph reality.
- Related artifact: immutable coder handoff `reports/coder/rp_coder_260822_055455_by_gpt-5_p6b_typescript_stdlib_authority.md`.
- Report identity protocol: this file is written once and externally sealed after creation. Its final filesystem createdAt, bytes, LF/CR counts, strict UTF-8/BOM state, and SHA-256 are supplied in the final handoff; a self-hash is intentionally not embedded because that would change the identified bytes.

## Executive Summary

- Problem: decide whether P6-B truthfully implements the accepted TypeScript standard-library authority at its real resolver boundary, not merely whether focused commands or aggregate counters are green.
- Decision: reject. Catalog provenance, profile closure, fail-closed loading/config, TypeScript-only injection, package isolation, artifact identity, and cleanup are substantially correct. Four required P6-B invariants remain open: handled declaration results are silently discarded at source-site level; catalog IDs violate the accepted canonical identity contract; local/repository member receivers can fall through to the external catalog; and the ledger’s `10/10` compiler-vector claim is broader than the durable fixtures.
- Required outcome: repair only the four findings below, add the exact re-review evidence, rerun proportionate affected gates, refresh the four ledgers truthfully, and return the same P6-B slice for independent re-review. Do not open P6-C1/C2/C3/D to hide or defer these P6-B defects.

## Blocking Findings

### [CRITICAL] Handled declaration results are discarded, leaving no observable source-site outcome

File: `internal/resolution/resolve.go:488`, `internal/resolution/resolve.go:615`, `internal/resolution/resolve.go:806`

Issue: the resolver obtains a rich `tsstdlib.LookupResult`, calls `recordTypeScriptLookup`, increments one aggregate metric, and immediately returns. It emits neither a reference nor a source-site record nor a reasoned unavailable record. The file/range/requested meaning/status/reason/symbol/provenance contained by `LookupResult` is lost. Aggregate counts and the absence of a generic gap cannot identify which site resolved, which site was unavailable, what target/reason applied, or whether a site was silently swallowed.

Evidence:

- P6-A’s immutable contract requires `resolved external target or structured unavailable/unresolved outcome` at report line 133 and requires every affected site to retain status, stage, requested name/meaning, language, target-or-reason, proof, and authority/catalog/config hashes at line 145.
- `resolveCall` invokes member/global lookup at lines 488-498 and returns when `recordTypeScriptLookup` reports handled. `resolveAccess` repeats the pattern at lines 615-626; `resolveTypeAnnotation` repeats it at lines 689-695.
- `recordTypeScriptLookup` at lines 806-823 only increments `ResolvedExternalDeclarations`, `ExternalCapabilityUnavailable`, `ExternalProfileExcluded`, or `ExternalMeaningMismatches`. It does not retain any field of the lookup result.
- `resolution.Result` and `ReferenceIndex` in `internal/resolution/types.go` contain no P6-B site-level authority-result collection. The four new fields at lines 26-29 are counters only.
- `internal/resolution/p6b_tsstdlib_test.go:61-73`, `:149-150`, and `:163-164` assert counters, gap absence, and no `ExternalSymbol`; they do not assert an observable per-site lookup result. `internal/analyze/p6b_tsstdlib_test.go:34-48` likewise infers success from aggregate metrics and missing diagnostics for selected target text.
- The fresh packaged analyze proves the runtime can execute, but its graph/metrics cannot reconstruct a handled lookup site from the discarded result. This is exactly the review boundary where aggregate metrics are insufficient.

Why this blocks acceptance: deferring P6-C2 external-node materialization and P6-C3 final shared outcome DTO is legal only if P6-B preserves the lookup result truthfully for those later consumers. Current code destroys that truth at the P6-B boundary. Later slices would have to redo/refactor P6-B lookup handling rather than consume an accepted result, and unavailable sites are silent today.

Fix direction: preserve one immutable site-level authority result for each handled TypeScript global/type/member lookup, including file/range or source-site ID, requested name/meaning, status/reason, resolved symbol/owner identity, declaration ranges, and catalog/profile/config hashes. Keep C2 node materialization and C3 final cross-stage DTO locked, but expose a lossless P6-B result that those slices can consume. Capability-unavailable/profile-excluded/meaning-mismatch must remain non-generic gaps without becoming invisible.

Re-review evidence required: direct resolver and built-runtime fixtures showing exact observable records for resolved global/type/member sites and each unavailable class; equality between site count and aggregate counters; no resolved/unavailable overlap; no inference from gap absence alone; current source proving the lookup payload is not discarded.

### [HIGH] Generated symbol identities violate the accepted canonical identity contract

File: `anvien-web/scripts/generate-tsstdlib-catalog.mjs:174`, `anvien-web/scripts/generate-tsstdlib-catalog.mjs:271`, `internal/tsstdlib/catalog.go:180`

Issue: generator IDs are only `global:<name>` and `global:<owner>.<member>`. They omit authority kind, TypeScript version, catalog hash, declaration library/range, semantic owner path, and meaning. The loader accepts any non-empty ID and does not enforce uniqueness. `LookupGlobal` and `LookupMember` expose these values as `SymbolID`/`OwnerID`, despite P6-A naming those fields as the canonical external identity input.

Evidence:

- P6-A line 104 requires identity to be deterministic from authority kind, TypeScript version, catalog hash, canonical declaration library/source range, semantic owner path, name, and meaning.
- Generator line 174 emits `global:${name}`; line 271 emits `global:${owner.getName()}.${memberName}`.
- Loader validation at `internal/tsstdlib/catalog.go:602` and `:623` checks only non-empty IDs and name ordering. It has no ID uniqueness or required-component validation.
- `LookupGlobal` assigns the name-only ID at line 180; `LookupMember` assigns name-only owner/member IDs at lines 208-209.
- Independent parsing of the exact checked-in artifact found 13,832 non-empty IDs, 424 duplicate-ID groups, zero format mismatches from the two name-only templates, zero IDs containing the catalog hash, and zero IDs containing a meaning token. Samples include `global:Console x2`, `global:CookieStore x2`, `global:Crypto x2`, and multiple `global:CSS.<member> x2` groups.
- The artifact itself is otherwise internally consistent: 2,030 symbols + 11,802 direct member rows = 13,832 ID-bearing rows. The duplicates therefore are semantic identity collisions, not empty/malformed rows.

Why this blocks acceptance: canonical identity is an explicit P6-B catalog contract, not a P6-C2 presentation choice. C2 cannot materialize deterministic `ExternalSymbol` facts from colliding `SymbolID` values without re-deriving/replacing P6-B identity semantics. Meaning lanes that share a name are not interchangeable targets.

Fix direction: define a deterministic two-phase identity scheme that avoids hash self-reference: compute a canonical semantic payload/catalog identity over the required authority/version/input/range/owner/name/meaning facts, then derive unique IDs from that stable identity and keep the final artifact hash as a separately validated envelope hash. Regenerate the catalog, validate all required ID components and global uniqueness in the loader, and return lane-specific canonical IDs.

Re-review evidence required: zero duplicate semantic IDs across all symbol/member meaning rows; machine-checkable proof that every returned ID incorporates or cryptographically commits to every P6-A-required component; loader rejection for duplicate/component drift; two clean byte-equal generations with equal logical/artifact hashes; updated artifact inventory and package-size measurement.

### [CRITICAL] Local/repository member receivers can be rescued by the external catalog

File: `internal/resolution/resolve.go:394`, `internal/resolution/resolve.go:488`, `internal/resolution/resolve.go:503`, `internal/resolution/resolve.go:615`

Issue: `resolveCall` records that a local variable receiver was resolved (`bindingReceiverResolved = true` at line 394), but it still performs external member lookup at lines 488-498 before the local-receiver guard at line 503. `lookupTypeScriptMember` at lines 771-784 uses a repository type binding when available, otherwise falls back to the receiver’s literal name. Thus a locally bound `Math` without a type binding, or a repository-owned colliding type whose requested member is absent, can resolve `Math.max` from the standard library. `resolveAccess` has the same external fallback at lines 615-626 and no equivalent resolved-binding terminal guard.

Evidence:

- P6-A line 136 requires member lookup only after an eligible external receiver has resolved; repository/P5 resolution is first and authoritative.
- `bindingReceiverResolved` only affects the later unresolved-return branch. It does not set `externalBlocked` and does not prevent the earlier catalog call.
- `lookupTypeScriptMember` selects `ownerName = receiver`/value meaning unless `resolveReceiverType` succeeds; a repository receiver is not proven external before catalog lookup.
- The current repository precedence test `internal/resolution/p6b_tsstdlib_test.go:76-103` covers only a free global `parseInt` shadow. It has no local `Math` receiver, repository-typed receiver with a missing/colliding member, or access-form shadow case.
- Current Anvien impact for each of `resolveCall`, `resolveAccess`, and `resolveTypeAnnotation` is CRITICAL: 29 impacted symbols, 12 affected files, 6 modules, and 32 processes. This is a real resolver invariant, not a narrow test concern.

Why this blocks acceptance: the candidate violates the immutable repository-first/member-eligibility contract and can convert a repository member miss into a false external success. It also combines with the first finding to suppress any trace that this wrong fallback occurred.

Fix direction: make external member lookup require an explicitly resolved external receiver result. A local/repository/import receiver claim must be terminal for external receiver fallback even when the member itself is missing. Preserve explicit-import terminality for call/access/type paths. Do not special-case `Math` or any target name.

Re-review evidence required: source and direct call/access fixtures for (1) a local `Math` value without type metadata, (2) a repository-owned colliding receiver type with a missing member, (3) an explicit-import receiver failure, and (4) a genuinely external receiver/member. Repository cases must not increment external counters or suppress their truthful repository miss; external cases must retain the site-level result required by finding one.

### [HIGH] The durable fixtures do not reproduce all ten accepted compiler vectors

File: `internal/tsstdlib/catalog_test.go:29`, `internal/tsstdlib/catalog_test.go:47`, `internal/tsstdlib/catalog_test.go:58`, `internal/tsstdlib/catalog_test.go:68`, `internal/tsstdlib/catalog_test.go:110`

Issue: P6-A requires all ten compiler vectors as independently authored general fixtures, but current tests cover several vectors only partially while both benchmark rows claim `10/10` reproduction.

Evidence:

- P6-A lines 65-74 define the exact ten-vector matrix; line 244 explicitly requires all ten as independent P6-B fixtures.
- The default-profile checks at `catalog_test.go:29-36` cover Promise type/value and Math/max/min but do not reproduce the default value vector containing Promise/Map/Set/Symbol together.
- ES2022 at lines 47-55 covers all four values.
- Explicit ES5 at lines 58-65 checks Promise type/value and Math global only; it omits Map/Set/Symbol value failures and does not verify Math member behavior under that profile.
- Explicit ES2015 at lines 68-74 checks Promise and Map values but omits Set and Symbol.
- ES5+Promise at lines 77-84 covers the four values.
- `noLib` at lines 110-128 checks only Promise type; it does not reproduce the accepted Promise-and-Math type/member vector.
- The benchmark ledger lines 35 and 37 state `10/10 reproduced by P6-B behavior fixtures` and `10/10 coder-final candidate`. The actual fixture matrix is narrower, so those rows are not truthful evidence of the exact requirement.

Why this blocks acceptance: complete durable reproduction is an explicit P6-B acceptance gate. The missing combinations are material to profile and meaning-lane behavior and overlap the repository/member defects above; helper checks for a subset cannot be counted as an entire compiler vector.

Fix direction: add a table-driven, independently authored production-behavior matrix matching each P6-A vector exactly, including all named values/types/members and expected resolved/excluded/unavailable statuses. Record vector-level results rather than deriving `10/10` from category overlap or aggregate counts.

Re-review evidence required: a ten-row fixture/result artifact with one exact row per P6-A vector; focused test output identifying every row; updated ledger language/counts that match the actual denominator; exact site-level outcomes for each requested symbol/meaning.

## Source-Level Clearance Notes

- `anvien-web/scripts/generate-tsstdlib-catalog.mjs` and `internal/tsstdlib/catalog.v1.json`: blocked only on canonical identity. Clear on developer-only/offline placement, exact locked TypeScript `5.9.3` and integrity checks, official sorted `lib*.d.ts` inputs, compiler-API declaration/member/base/value-owner extraction, deterministic ordering, and absence of target-name branches.
- `internal/tsstdlib/catalog.go`: blocked on ID contract. Clear on `go:embed`, strict DTO decoding, trailing-token rejection, schema/version/integrity/generator validation, logical hash validation, input/reference/lane ordering checks, meaning separation, active-profile declaration filtering, and inherited-member traversal.
- `internal/tsstdlib/profile.go`: clear. Zero/one root `tsconfig.json`, JSONC data-only parsing, `target`/`lib`/`noLib`, explicit-lib closure, noLib disablement, invalid/unreadable/unsupported topology, nested/multiple configs, and `jsconfig.json` all follow the narrow fail-closed P6-A boundary. Independent closure calculation reproduced ES2022 `63 / 2,389,540`, ES2015 `13 / 310,653`, and ES5 `3 / 232,957`.
- `internal/analyze/analyze.go`: clear. Authority injection occurs only when scanned inventory contains TypeScript (`:295-299`, `:468-477`) and passes the complete scan path inventory to profile selection. No scanner-corpus mutation or target-name allowlist was added.
- `internal/resolution/types.go` and `internal/resolution/indexes.go`: wiring is clear, but the result surface is blocked by finding one. The authority pointer remains separate from repository ScopeIR; repository workspace ownership is not polluted. The four new metrics cannot substitute for site-level truth.
- `internal/resolution/resolve.go`: global/type repository precedence and the tested explicit-import constructor failure are clear. Member receiver precedence and handled-result observability are blocked by findings one and three.
- Package/runtime surface: clear. `anvien-web/package.json`, `anvien-web/package-lock.json`, and `anvien/package.json` have zero diff. Go production source contains no `os/exec`, `exec.Command`, Node/`tsc`/npm discovery, HTTP/network request, install, or package-script execution path. The only Go `node` string is static generator provenance validation; Node and TypeScript imports exist only in the offline generator.
- Tests/fixtures: blocked by findings one, three, and four. They establish many catalog/profile behaviors but do not close site observability, member-shadow precedence, canonical ID uniqueness, or the exact ten-vector denominator.
- Four living ledgers: clear on P6-B remaining unchecked/uncommitted, P6-C1/C2/C3/D locks, package/build limitations, artifact sizes, cleanup, and historical command classifications. Blocked on the `10/10` fixture claim. Historical detect counts are retained as historical input but are not current graph truth.

## Current Anvien Evidence

- Repo-local CLI help was read first. Exactly one fresh `analyze 'E:\Anvien' --force --json` ran before graph queries; it succeeded with 2,007 scanned, 750 parsed code, 0 failed, 111,540 nodes, and 156,275 relationships at `E:\Anvien\.anvien\graph.json` on indexed/current commit `b98131e...`. The `+1 file / +16 nodes / +16 relationships` versus the coder snapshot coincides with the now-sealed coder report being present; no causal assumption is needed for the verdict. The later protected Main handoff #19 is excluded by explicit Owner authority, and analyze was not restarted.
- Fresh file-detail summaries are current/stale=false and high risk for all parsed production owners: analyze 332 symbols / 190 related files / 20 flows / 8 tests; resolution types 69 / 36 / 1 / 19; indexes 343 / 60 / 18 / 32; resolve 157 / 67 / 22 / 35; catalog 239 / 7 / 1 / 2; profile 91 / 6 / 1 / 1. The generator is parsed JavaScript, high file risk, 109 symbols, no related file/flow/test. The JSON catalog is metadata and has no graph file node.
- Current upstream file impacts with tests are CRITICAL for analyze `102 symbols / 24 files`, resolution types `65 / 15`, indexes `242 / 36`, resolve `120 / 44`, catalog `247 / 62`, and profile `36 / 11`; generator file impact is MEDIUM `11 / 1`.
- Exact current symbol impacts: `resolveCall`, `resolveAccess`, and `resolveTypeAnnotation` each CRITICAL `29 symbols / 12 files / 6 modules / 32 processes`; `recordTypeScriptLookup` CRITICAL `8 / 4 / 2 / 34`; `loadCatalog` CRITICAL `13 / 4 / 3 / 12`; `selectProfile` CRITICAL `33 / 10 / 7 / 21`; `typeScriptDeclarationInventory` CRITICAL `24 / 8 / 6 / 31`. These are blast-radius warnings, not edit prohibitions.
- Current explicit-path `detect-changes --repo E:\Anvien --scope all --json` exited 0 with CRITICAL risk: affected count 23, affected files 8, changed count 148, changed files 8; affected app layers backend 19 / mixed 4; affected functional areas analyzer 8 / mixed 6 / resolution 9. This supersedes the coder’s historical raw `23/7/143/7` snapshot for current-review reporting. Candidate Git inventory remains exactly eight tracked modifications, so the larger current graph count is not hidden or normalized away.

## Artifact, Build, Runtime, Regression, and Benchmark Disposition

Passed/current independent checks:

- Immutable coder handoff seal exactly reproduced: 16,002 bytes, 197 LF, 0 CR, strict UTF-8, no BOM, SHA-256 `3326E2658522F28824C16978CB485D34B74CD9F7CA4F03C994812E81D13D072A`, filesystem createdAt `2026-08-22T05:56:38.8110838+07:00`.
- Catalog artifact independently reproduced: 1,388,709 bytes, SHA-256 `78C9F6F49896D2F11E7E45B58C9D7CB90E1CA20413F06FFE9FAC95310C1AF054`; stored and recomputed logical hash both `22ab8f6986bfa8efd01addffac64b4265fff09244391ecb5d385b80dfa187119`.
- Every one of the 100 E-local TypeScript declaration inputs was hashed against the manifest: 3,141,835 bytes total, 0 mismatches. Catalog counts are 2,030 symbols and 11,802 direct member rows.
- Final packaged artifact independently reproduced: `anvien/bin/anvien.exe`, 72,961,536 bytes, SHA-256 `E6EE2F57BAA0EF57531EA9B45B288174FE23AE7D9591CCDDE40EE27231FEFE41`, Go `1.26.3`, `ladybugdb`, VCS basis `b98131e...`, modified=true. Against the accepted P5-D baseline 71,509,504 bytes / SHA `BC060FDDF46394B72859B86930409B1B7E91C996BA304A86F0A633C37E043532`, the recorded delta is +1,452,032 bytes (+2.03%).
- The same packaged binary successfully served help and completed the one fresh authoritative self-analyze. This is a normal-environment runtime check, not a sanitized-PATH/no-Node proof.
- All six declared dead/debug paths are absent: `.tmp/p6b-graph-evidence`, both P6-B benchmark JSON captures, fixture `.anvien`, and both launcher wrapper executables.

Historical coder evidence retained but not promoted to fresh acceptance proof:

- The official lifecycle build was prohibited and was not run because it executes installation/package scripts. This remains an explicit not-run boundary.
- Broad `go build ./...` failed on intentional non-buildable fixtures. It is not a successful full-build gate and is not presented as one.
- `go build ./cmd/... ./internal/...`, both shipping launcher Go-module builds, and direct package-runtime build were reported successful on the sealed candidate. They were not rerun after source inspection established rejection blockers.
- Full `go test ./cmd/... ./internal/... -count=1` was reported to fail exactly the two accepted out-of-scope C#/Dart baselines: C# expected `ACCESSES=2`, got none; Dart expected `ACCESSES=2`, got `1`. Those failures are preserved, not relabeled.
- The E-only sanitized PATH probe exited at native loader failure `-1073741515` before CLI startup. It is not a no-Node success and supplies no acceptance proof.
- Prior generator determinism remains the historical `2/2` gate and was not repeated because generator/catalog bytes were unchanged and the review instruction explicitly prohibited mechanical reruns. Microbenchmark ranges were not rerun; no Owner threshold exists. Neither historical determinism nor benchmark performance cures the identity and resolver defects.

Not run in this review:

- No network, install, package script, target access, alternate checkout/worktree, stage, commit, push, code repair, ledger repair, or subagent/lane action.
- No full/focused test or build rerun after the source gate failed. Repeating green subsets cannot prove the missing invariants.
- No `E:\cheapapp.org` access. Target remains locked to P6-D.

## Invariant Closure

- Affected invariant family: deterministic declaration authority + config profile + repository-first eligibility + source-site observability + canonical identity + exact compiler-vector proof.
- Closed surfaces: locked TypeScript provenance/input manifest; logical/artifact hashes; schema/version/trailing-data fail-closed loader; value/type/namespace lanes; direct/inherited member lookup mechanics; profile closures; JSONC and fail-closed topology; TypeScript-only analyze injection; no target-name branch; no runtime Node/network/install/script path; package binary/artifact/cleanup boundary; global/type and tested explicit-import precedence.
- Open same-invariant surfaces: lossless site-level result retention, canonical unique semantic IDs, local/repository member receiver terminality, and exact ten-vector fixture coverage.
- P6-C1/C2/C3/D disposition: unchanged and locked. C2/C3 deferral does not authorize P6-B to discard lookup truth or ship colliding identity inputs.

## Required Fix List For Resubmission

1. Preserve a lossless per-source-site P6-B lookup result for resolved and unavailable outcomes; prove it at direct and built boundaries without aggregate-only inference.
2. Replace name-only/colliding IDs with the complete canonical semantic identity required by P6-A; regenerate and validate uniqueness/components/determinism.
3. Make external member lookup contingent on an actually external receiver; add local/repository/import receiver shadow regressions for both call and access paths.
4. Implement and record the exact ten-row P6-A compiler-vector fixture matrix, then correct the four ledgers to the real denominator/results.
5. After those source fixes: perform required fresh graph/file-detail/impact, proportionate build/runtime/regression/benchmark validation, explicit-path detect-changes with full CRITICAL/HIGH counts, exact cleanup/boundary inventory, and a new immutable coder handoff for independent re-review.

## Overall Evaluation

The candidate establishes a strong offline catalog and packaging foundation, and its provenance/profile/artifact evidence is credible. Acceptance nevertheless requires the actual resolver invariant, not a favorable aggregate. Current code can silently swallow handled sites, expose non-canonical colliding IDs, and let a repository receiver fall through to external member lookup; the required ten-vector proof is also incomplete. These are P6-B defects and evidence gaps, so no same-scope acceptance is possible until they are repaired and independently re-verified.
