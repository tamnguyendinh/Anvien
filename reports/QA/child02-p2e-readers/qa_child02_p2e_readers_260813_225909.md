# Child 02 P2-E affected-reader command evidence

- Status: **PASS_PENDING_FRESH_UI**
- Denominator: `8`
- C09-C11 require the separate fresh mounted visible Playwright evidence before final 8/8 disposition.
- Remote VECTOR attempts: `0`

| Row | Scenario | Boundary | Result |
|---|---|---|---|
| C09 | exact opaque-ID resolution | built frontend unit + mounted Playwright | PASS_PENDING_FRESH_UI |
| C10 | unique grounding fail-closed | built frontend unit + mounted Playwright | PASS_PENDING_FRESH_UI |
| C11 | one-based lines / zero-based UTF-8 byte columns / exclusive ends | built frontend unit + mounted Playwright | PASS_PENDING_FRESH_UI |
| C12 | nodeRange | internal/filecontext | PASS |
| C13 | Definition MCP context | internal/mcp | PASS |
| C14 | detectChangedSymbols | internal/mcp | PASS |
| C15 | collectRenameChanges | internal/mcp | PASS |
| C16 | persisted explicit embedding label and semantic-search hydration | embeddings + HTTP + repo-local native Ladybug | PASS |

Full commands, exits, UTC intervals, stdout, stderr, durations, and peak working sets are in the paired JSON.
