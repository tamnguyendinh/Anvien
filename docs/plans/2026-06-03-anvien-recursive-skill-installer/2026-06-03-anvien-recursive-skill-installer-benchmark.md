# Benchmark Ledger

Date: 2026-06-03

Status: Complete

Companion files:

- Plan: [2026-06-03-anvien-recursive-skill-installer-plan.md](2026-06-03-anvien-recursive-skill-installer-plan.md)
- Evidence ledger: [2026-06-03-anvien-recursive-skill-installer-evidence.md](2026-06-03-anvien-recursive-skill-installer-evidence.md)

## Benchmark Rules

1. Record measured inventory, size, package, startup, throughput, or runtime metrics here.
2. Build/test pass-fail belongs in the evidence ledger unless the timing/size/count is the measured target.
3. For this plan, skill package inventory, package payload size, script/resource payload counts, package/fallback source size, and generated install counts are benchmarkable.
4. Record before/after numbers as each implementation phase changes the installer.

## B0 - Planning Baseline Inventory

Status: Complete

### Graph Inventory

Source: `anvien analyze --force` on 2026-06-03.

| Metric | Baseline |
|---|---:|
| Files scanned | 1352 |
| Parsed code files | 702 |
| Failed parses | 0 |
| Indexed documents | 405 |
| Indexed metadata files | 114 |
| Indexed scripts | 7 |
| Indexed static files | 3 |
| Graph nodes | 83838 |
| Graph relationships | 122361 |

### Skill Package Source Inventory

Source: filesystem scan under `internal/aicontext/skills`.

| Metric | Baseline |
|---|---:|
| Source skill packages | 35 |
| `SKILL.md` entry files | 48 |
| Packages with multiple `SKILL.md` entries | 3 |
| Empty top-level package folders | 0 |
| Maximum observed `SKILL.md` entry depth within a package | 2 |
| Packages with files under `scripts/` | 16 |
| Total package payload files | 607 |
| Total package payload bytes | 11311115 |
| Total package script files | 150 |

Depth examples:

| Skill path | Depth |
|---|---:|
| `document-skills/docx/SKILL.md` | 2 |
| `document-skills/pdf/SKILL.md` | 2 |
| `problem-solving/when-stuck/SKILL.md` | 2 |
| `debugging/root-cause-tracing/SKILL.md` | 2 |

### Top-Level Package Payload Inventory

Source: filesystem scan under each top-level `internal/aicontext/skills/<top>` directory.

| Top-level package | Skill entries | Files | Bytes |
|---|---:|---:|---:|
| `aesthetic` | 1 | 7 | 20611 |
| `ai-multimodal` | 1 | 15 | 194154 |
| `anvien-api-surface` | 1 | 1 | 3318 |
| `anvien-debugging` | 1 | 1 | 3507 |
| `anvien-planner` | 1 | 1 | 4418 |
| `anvien-refactoring` | 1 | 1 | 2611 |
| `backend-development` | 1 | 12 | 127123 |
| `better-auth` | 1 | 10 | 185154 |
| `bunny` | 1 | 6 | 38608 |
| `chrome-devtools` | 1 | 26 | 173146 |
| `code-review` | 1 | 4 | 18505 |
| `context-engineering` | 1 | 17 | 50704 |
| `databases` | 1 | 19 | 269268 |
| `debugging` | 5 | 11 | 39034 |
| `devops` | 1 | 29 | 134129 |
| `docs-seeker` | 1 | 8 | 91323 |
| `document-skills` | 4 | 130 | 2498732 |
| `frontend-design` | 1 | 2 | 15459 |
| `frontend-development` | 1 | 11 | 131925 |
| `google-adk-python` | 1 | 1 | 6777 |
| `mcp-builder` | 1 | 10 | 147050 |
| `mcp-management` | 1 | 15 | 127927 |
| `media-processing` | 1 | 16 | 190901 |
| `mermaidjs-v11` | 1 | 6 | 30064 |
| `payment-integration` | 1 | 43 | 248066 |
| `problem-solving` | 7 | 8 | 19841 |
| `repomix` | 1 | 9 | 106939 |
| `sequential-thinking` | 1 | 4 | 17193 |
| `shopify` | 1 | 10 | 172844 |
| `skill-creator` | 1 | 16 | 72202 |
| `template-skill` | 1 | 1 | 140 |
| `threejs` | 1 | 21 | 147247 |
| `ui-styling` | 1 | 98 | 5750719 |
| `web-frameworks` | 1 | 18 | 229991 |
| `web-testing` | 1 | 20 | 41485 |

### Script-Heavy Package Inventory

Source: filesystem scan under each top-level skill package root. These rows are useful because they prove the installer must copy package payloads, not only `SKILL.md`.

| Package install root | Max entry depth | Files | Bytes | Script files |
|---|---:|---:|---:|---:|
| `chrome-devtools` | 1 | 26 | 173146 | 21 |
| `document-skills` | 2 | 130 | 2498732 | 37 |
| `databases` | 1 | 19 | 269268 | 10 |
| `web-frameworks` | 1 | 18 | 229991 | 9 |
| `media-processing` | 1 | 16 | 190901 | 9 |
| `ai-multimodal` | 1 | 15 | 194154 | 9 |
| `mcp-management` | 1 | 15 | 127927 | 9 |
| `ui-styling` | 1 | 98 | 5750719 | 8 |
| `payment-integration` | 1 | 43 | 248066 | 6 |
| `repomix` | 1 | 9 | 106939 | 6 |
| `devops` | 1 | 29 | 134129 | 6 |
| `better-auth` | 1 | 10 | 185154 | 5 |
| `shopify` | 1 | 10 | 172844 | 5 |
| `skill-creator` | 1 | 16 | 72202 | 4 |
| `mcp-builder` | 1 | 10 | 147050 | 4 |
| `context-engineering` | 1 | 17 | 50704 | 2 |

## B1 - Implementation Benchmarks

Status: Complete

### Skill Package Catalog And Install Counts

Source: package catalog and installer tests after implementation.

| Phase | Metric | Baseline | Latest | Final | Unit | Notes |
|---|---|---:|---:|---:|---|---|
| P1-A | Discovered top-level packages | 35 | 35 | 35 | count | Catalog enumerates immediate children of `internal/aicontext/skills`. |
| P1-A | Discovered `SKILL.md` entries | 48 | 48 | 48 | count | Nested entries remain inside their top-level package. |
| P2-A | Manifest file path | 0 | 1 | 1 | file | `.claude/skills/anvien/.anvien-skill-manifest.json`. |
| P2-A | First install package IDs | 4 old base skills | 35 | 35 | count | `SkillInstallResult.Installed=35` in test. |
| P2-A | Same-hash skipped managed packages | 0 | 34 | 34 | count | Second install test tampers one managed package, leaving 34 skipped. |
| P2-B | Preserved unmanifested target packages | 0 | 1 | 1 | count | Test fixture preserves `custom-skill`. |
| P2-B | Rejected unmanaged same-name collisions | 0 | 1 | 1 | count | Test fixture rejects unmanifested `ui-styling` collision. |
| P2-C | Adopted legacy Anvien package | 0 | 1 | 1 | count | Test fixture adopts old `anvien-planner` output by frontmatter name. |
| P2-C | Stale managed package entries marked | 0 | 1 | 1 | count | Test fixture marks `old-managed` stale without deleting payload. |
| P3-A | Installed package payload files | 4 old `SKILL.md` files | 607 | 607 | files | Full source package payload is copied, including dotfiles/scripts/assets. |
| P3-A | Installed package payload bytes | 13854 | 11311115 | 11311115 | bytes | Full source package payload is hashed/copied without child-name exclusions. |
| P3-A | Installed package script files | 0 | 150 | 150 | files | Scripts are copied but not executed. |
| P4-A | Setup/analyze package report counters | 1 count only | 8 counters | 8 counters | fields | `discovered`, `installed`, `updated`, `skipped`, `adopted`, `stale`, `preserved`, `collisions`. |

### Representative Payload Checks

| Package | Payload checked | Purpose |
|---|---|---|
| `debugging` | `root-cause-tracing/find-polluter.sh` | Proves nested child skill script travels with package. |
| `document-skills` | `docx/scripts/document.py` | Proves multi-entry document package scripts are included. |
| `ui-styling` | `scripts/shadcn_add.py` | Proves root package scripts are included. |
| `ui-styling` | `canvas-fonts/ArsenalSC-Regular.ttf` | Proves non-script binary assets are included. |
