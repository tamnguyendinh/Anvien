# Child 04 P4-C2 Source Oracle Authoring Report

- Oracle ID: p4c2-oracle-v1-a869876ab626-260821_110849+0700
- Evidence ID: E4-P4C2-ORACLE1
- Role: clean-context source-only Oracle Authoring / Data-Integrity lane
- Target basis: a869876ab6262dacde6cd5d432d099a91852a646 on master
- Authoring result: PASS_TO_SEAL

## Source basis

Only the three hash-pinned TypeScript files named in source-basis.json supplied semantic input. Their byte counts and SHA-256 identities matched before source reading and again after expected-value authoring. Target HEAD, branch, and the canonical tracked-status identity also remained unchanged. The seven tracked modifications recorded at entry were pre-existing and outside the three semantic-input files.

No target file was written. No target analyzer state, target .anvien content, Anvien implementation/test/golden material, QA output, historical report content, Child 05 state, or historical .tmp oracle material supplied an expected value.

## Derivation

The positive set contains exactly 21 independently anchored top-level direct exports:

- 15 named type-alias declarations with type meaning and explicit type-only state.
- 6 named function declarations with value meaning and non-type-only state.

Every positive declaration has one direct export fact, one File-to-Export relation, absent access state, isExported=true compatibility, zero export diagnostics, and empty Child 05-derived state. Local and exported names coincide because every positive occurrence is a named direct declaration rather than an alias or re-export.

The negative set contains exactly 11 independently anchored declarations in the email source:

- one non-exported top-level function;
- four same-name or ordinary constant bindings distinguished by their lexical owners, including the reducer callback nested in latestIso;
- six array-binding leaves owned by readEmailOperationsReport.

Every negative occurrence lacks direct module-export syntax, preserves absent access state, has zero Export facts and zero File-to-Export relations, retains isExported=false, and has empty Child 05-derived state.

## Coordinates and identity

All coordinates are one-based with exclusive end positions. Positive ranges cover the complete exported declaration and selection ranges cover the declaring identifier. Negative ranges cover the independently identified declaration occurrence and selection ranges cover its identifier. Matching is by source path, file SHA-256, exact range, lexical owner, and local name; no implementation-generated graph identifier is present.

## Fail-closed result

The expected-value file parses with declared and actual counts of 21+11. Every required source hash matched at both authoring checkpoints, and no supplied anchor was missing or duplicated. Provenance attestations record zero forbidden inputs, zero analyzer/QA observations before seal, zero target writes, and zero evidence-bearing .tmp artifacts.

The bundle is eligible for sealing only if the final validation repeats these invariants, confirms complete row schemas and uniqueness, recomputes the same three source hashes immediately before seal, and constructs the non-circular digest from the five non-seal files. seal.json is the authoritative final status and must be written last.
