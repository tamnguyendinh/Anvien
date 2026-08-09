# Anvien Graph Identity, Generation, and Migration Contract

Status: accepted for implementation by Owner instruction on 2026-08-09
Authority: `docs/plans/2026-07-26-anvien-graph-identity-resolution-v2-multi-plan/2026-07-26-anvien-graph-identity-resolution-v2-roadmap.md`
Source problem: `reports/problem/2026-07-26-tinh-chinh-cac-van-de-cua-anvien`

## Purpose

Define the v2 identity and generation invariants used by the seven-child execution campaign. This document is the P1-A authority boundary; implementation details belong to the child slice owners.

## Canonical model

```text
File --DECLARES--> Declaration --DECLARES_SYMBOL--> Symbol
Module --EXPORTS/REEXPORTS-----------------------> Symbol
SourceSite --ResolutionOutcome--> InternalSymbol | ExternalSymbol | Gap
```

`DeclarationID` identifies one source occurrence. `SymbolID` identifies one logical language symbol. They are distinct, opaque, versioned types and are never substituted for one another.

## Identity and range rules

- Identity version is explicit (`identity-v2`) and is included in every canonical tuple.
- Canonical tuples use repo-relative, normalized module/file identity; lexical owner chain; semantic name; meaning namespace; declaration role; provider merge key; and binding path. Absolute paths, worker order, random UUIDs, and raw ranges are not identity sources.
- `DeclarationID` additionally includes declaration role, binding path, selection anchor, and a deterministic discriminator.
- `SymbolID` excludes mutable body/comment text and source range. Body/comment edits preserve it; rename, owner/module move, language meaning, or unsupported merge evidence may change it.
- Ranges are zero-based, half-open byte offsets internally. Line/column are zero-based and use `utf-8-bytes` encoding in the graph contract; UTF-16 conversion occurs only at editor/API boundaries. `Range` covers the construct and `SelectionRange` covers the identifying token.
- Unicode, combining characters, emoji, tabs, CRLF, and LF must round-trip without changing byte boundaries.

## Mutation and integrity rules

- Canonical insertion with a conflicting payload returns a structured `IdentityCollision` error and aborts the candidate generation.
- Identical reinsertion is allowed only through an explicit idempotent operation that checks identity and provenance.
- Enrichment and replacement are separate explicit operations. A generic upsert, first-wins, last-wins, warning-only, or panic-only path is forbidden.
- Every relationship has a deterministic `RelationshipID` and retains all source-site IDs/provenance when aggregated.
- Validation rejects duplicate IDs, missing endpoints, mixed generations, invalid ranges, and incomplete closure before publication.

## Versioned reference and publication rules

A durable cross-boundary reference is:

```text
{repoKey, graphGeneration, graphSchemaVersion, identitySchemaVersion, symbolID}
```

The active manifest is the compatibility authority and must expose:

```text
protocolVersion
minReaderProtocol
minReaderBuild
graphSchemaVersion
identitySchemaVersion
scopeIrVersion
generation
analyzerVersion
positionEncoding
sourceFingerprint
configHash
catalogHash
```

Readers validate the typed handshake before opening graph bodies, caches, groups, embeddings, HTTP/Web streams, or Cypher rows. A mismatch returns `INDEX_VERSION_MISMATCH` and does not expose stale or mixed records.

Each analyze writes a fully staged, validated generation and publishes it with one atomic active-pointer/CAS update. Any failure leaves the prior generation queryable. The active v1 adapter remains byte/hash-preserving until Child 02's cutover slice; v2 shadow output is never consumed by v1 readers.

## Symbol merge and legacy behavior

Overloads or declaration merging share a `SymbolID` only with provider language-semantic evidence. Unverified candidates remain separate. A legacy ID with one candidate may redirect with a warning; a collided legacy ID returns `ambiguous` plus all candidates. No global same-name fallback may repair an explicit import or identity miss.

## Ownership and acceptance

P1-A owns this authority document only. P1-B owns range and identity types; P1-C0/C0A/C0B own occurrence, relationship, and decode validation; P1-C owns declaration-to-symbol mapping; P1-D and D1-D3 own strict mutation and producer adapters; P1-E owns shadow-v2 comparison. No owner may add resolver, binding, export, or ambient semantics to these files.

P1-A is accepted because the Owner explicitly authorized implementation on 2026-08-09, all proposed decisions are resolved above, the rejected shortcuts are recorded, and the downstream child dependency is unambiguous. A later architecture change requires a new contract revision and a refreshed child plan before editing.
