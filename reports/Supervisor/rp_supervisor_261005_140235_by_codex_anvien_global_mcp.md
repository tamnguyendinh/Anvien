# Anvien global MCP review

Verdict: PASS

Authority: user requested Anvien as a global MCP for this Codex environment.
Claim: global Codex configuration successfully launches Anvien from outside the repository.

## Completed checklist

- [x] Inspect official OpenAI MCP documentation and existing global registration.
- [x] Replace the unresolved PATH command with the absolute runtime path.
- [x] Independently re-read global TOML and verify registration from the user profile directory.
- [x] Verify MCP initialize and tools/list from outside the repository.

## Evidence

- Official documentation: https://developers.openai.com/codex/mcp/
- Configuration: C:\Users\TAM NGUYEN\.codex\config.toml, section mcp_servers.anvien.
- Command: E:\Anvien\anvien\bin\anvien.exe; arguments: mcp; enabled: true; cwd: unset.
- codex mcp add reported Added global MCP server.
- codex mcp get anvien --json from C:\Users\TAM NGUYEN verified the enabled global stdio registration.
- Direct MCP initialize returned anvien 1.3.0 and protocol 2024-11-05.
- Direct tools/list returned 16 tools: list_repos, query, cypher, context, detect_changes, rename, impact, route_map, tool_map, shape_check, api_impact, group_list, group_sync, group_contracts, group_query, group_status.
- Verification process was terminated after inspection.

## Invariant closure

The global server resolves outside the repository and completes initialization and tool discovery. No residual unverified surfaces remain within this claim. Existing running chats have not reloaded their tools; a new chat or Codex restart may be required. The configured executable must remain at its current path.

This was a user configuration change using an existing runtime. Product source, graph, API, and project MCP settings were not edited; product rebuild, graph refresh, blast-radius analysis, benchmarks, and an implementation commit were not performed.
Review time: 2026-10-05T14:02:35.3128575+07:00 (Asia/Bangkok)
