# Child 02 P2-E repeated analyze evidence

- Overall: **PASS**
- Built CLI: `E:\Anvien\anvien\bin\anvien.exe`
- Built CLI SHA-256: `14DEB1820B58E4BBE68E5C8B542D09231CFAB49FA73521BC0163DA754588606B`
- Same isolated repository path: `E:\Anvien\.tmp\qa-child02-p2e\fixture`
- Isolated ANVIEN_HOME: `E:\Anvien\.tmp\qa-child02-p2e\anvien-home`
- Accepted fixture SHA-256: `327F6BD02FCB341078902FE4302CAB03A2077DDFE98B379164C17E96B547A078`

## Matrix

| ID | Scenario | Verdict | What it proves |
|---|---|---|---|
| M1 | two unchanged normal built analyze runs | PASS | same path + equal input manifests preserve accepted fact identity/ranges/endpoints; whole-artifact hashes are informational |
| M2 | changed source then current artifact/read | PASS | next successful artifact/read includes later, reduces now to one, and excludes the prior second-now identity |
| M3 | clear analyze-command failure | PASS | owned storage failure returns nonzero and no graph at the normal expected artifact path |
| M4 | C17 native availability with Graph JSON absent | PASS | successful current read used Ladybug, because graph.json was absent for the full MCP invocation |
| M5 | C17 no-readable-backend failure | PASS | native open failure is not ErrUnavailable fallback; with both backends absent MCP returns JSON-RPC non-success and no stale result |
| M6 | subsequent successful analyze/read recovery | PASS | restored baseline input produces the original fact signature and readable current rows after both faults |
| M7 | accepted corrected-fact semantics | PASS | exact IDs, construct/selection ranges, 10/10 DEFINES, and zero missing endpoints were compared rather than aggregate totals alone |

## Artifact and backend classification

- C04: Ladybug load completes before atomic Graph JSON temp+rename; errors prevent normal CLI registration/success output.
- C17: Ladybug is primary. Only the exact `ErrUnavailable` sentinel permits fallback to the same repository's `graph.json`. Other native open/query errors fail clearly.
- Native proof removed `graph.json` for the entire MCP call; current changed facts were still returned, proving Ladybug use.
- No-backend proof removed both `lbug` and `graph.json`; MCP returned a JSON-RPC error and no stale/substitute rows.
- The freshly built normal binary includes the `ladybugdb` tag. The supported fallback is classified from source but was not fabricated through a non-normal/dev build.
- Whole Graph JSON/Ladybug hashes are recorded only as artifact identity. Determinism verdict uses the accepted canonical fact signature, ranges, identities, and endpoints.

## Exact corrected facts

- Unchanged runs: definitions `10/10`, DEFINES `10/10`, missing endpoints `0/0`.
- Unchanged fact signature: `FCE1E085BFE0352FF219C2FC33E5B5DFBEE2D49A3B73EAFE89CF6490CFF1BAAD`.
- Changed run: time `2`, now `1`, later `1`, old second-now excluded `True`.
- Recovery fact signature equals baseline: `True`.

Full command output, exit codes, timings, peak working sets, artifact hashes/timestamps, fact rows, reads, fault injection/restoration, and clean-holder/full-build evidence are in the paired JSON.
