---
name: repo-bootstrap
description: Prepare foundational documents and decisions before coding a new repository or project. Use when the user wants to bootstrap a repo, start a new project, write technical specs for an AI coding agent or dev team, define architecture, choose tooling, prepare handoff docs, or set up a repo before implementation. This skill is repo-type-aware: first classify the repository, then produce only the core and conditional sections that apply.
---

# Repo Bootstrap

Goal: create the minimum complete source-of-truth package that lets an AI agent
or developer start implementation without guessing. Do not force every repo into
an app/backend/SaaS shape. Classify the repository first, then select the right
lanes.

## Core Rule

Bootstrap is decision work before coding. It must make the next implementation
agent's choices explicit:

- what the repo is
- what it must and must not do
- which technologies are approved
- where boundaries are
- how to run, test, package, release, and verify it
- which decisions are still blocked

If a decision affects code, dependencies, public API, data shape, runtime
behavior, deployment, or QA, do not leave it implicit.

## Workflow

Run these phases in order.

```text
[0] Repository Type Classification
[1] Project Context
[2] Architecture / Repo Blueprint
[3] Toolchain & Dependency Policy
[4] Conditional Domain Specs
[5] Technical Standards
[6] Environment, Packaging, Release
[7] Test & QA Strategy
[8] Handoff Checklist
```

Do not skip phase 0. The old fixed sequence of Database -> API -> DevOps is
valid for many app repos, but wrong for many libraries, CLIs, docs repos,
infrastructure repos, MCP servers, firmware, games, and data/ML repos.

## Phase 0 - Repository Type Classification

Classify the repo before writing detailed specs.

Record:

```markdown
## Repository Classification

- Repository type:
- Primary users:
- Runtime target:
- Main deliverable:
- Interface surface:
- Data persistence:
- Deployment/distribution:
- Required conditional lanes:
- Explicitly not applicable lanes:
```

Use this guide:

| Repository Need | Select Lane |
| --- | --- |
| Web app, fullstack app, backend API, SaaS product | App / Service |
| Reusable package, SDK, framework, internal library | Library / SDK |
| Command-line tool or developer utility | CLI |
| Terraform/Pulumi/CloudFormation/Kubernetes/platform config | Infrastructure / IaC |
| Data pipeline, analytics, ML/model repo | Data / ML |
| Native or cross-platform mobile app | Mobile |
| Browser extension | Browser Extension |
| MCP server, plugin, agent tool integration | MCP / Plugin |
| Docs/wiki/content-only repo | Documentation |
| Embedded, hardware, firmware | Embedded / Firmware |
| Game, graphics, simulation, 3D app | Game / Graphics |
| Electron/Tauri/native desktop app | Desktop App |
| Monorepo with multiple packages/apps | Monorepo |
| Multi-tenant SaaS | SaaS Multi-Tenant |
| Frontend with complex UI system | Design System |
| Blockchain/Web3 | Blockchain / Web3 |

For detailed conditional outputs, read `references/repo-type-lanes.md`.
For desktop apps, also read `references/desktop-app.md`.
For SaaS multi-tenant systems, also read `references/saas-patterns.md`.
For complex UI systems, also read `references/design-system.md`.
For technical standards details, read `references/technical-standards.md`.

## Phase 1 - Project Context

Define the project boundary. Every later decision should point back here.

Output:

```markdown
## Project Context

- Project/repo name:
- One-sentence description:
- Target users:
- Primary use cases:
- MVP scope:
- Out of scope:
- Non-functional requirements:
- Main language/runtime:
- Team size and roles:
- Milestones/deadline:
- Success criteria:
```

## Phase 2 - Architecture / Repo Blueprint

Describe the repository shape and system boundaries at the right abstraction
level for the repo type.

Output:

```markdown
## Architecture / Repo Blueprint

### Overview
[Text diagram, Mermaid, or clear component/module map.]

### Module / Package Breakdown
| Module / package | Responsibility | Boundary / not responsible for |
| --- | --- | --- |

### Interface Surface
[HTTP routes, CLI commands, library public API, MCP tools, extension messages,
mobile screens, jobs, events, hardware IO, or docs publishing surface.]

### Data / State Flow
[Where data/state enters, transforms, persists, exits, or renders.]

### Integration Points
- Internal:
- External:

### Scalability / Reliability / Safety Notes
```

For monorepos, include workspace/package graph and ownership boundaries.

## Phase 3 - Toolchain & Dependency Policy

This is a binding decision section. Coders must not invent dependencies after
bootstrap unless the user approves the change.

Output:

```markdown
## Toolchain & Dependency Policy

### Runtime & Language
- Language:
- Runtime:
- Package manager:
- Build tool:

### Core Frameworks / Libraries
| Area | Approved choice | Version | Reason |
| --- | --- | --- | --- |

### Dependency Rules
- Approved dependencies:
- Banned dependencies:
- Version pinning policy:
- License constraints:
- Security scanning:
- Rule for adding new dependency:

### Dev Tooling
- Formatter:
- Linter:
- Type checker:
- Git hooks:
- Local commands:
```

Rule for AI agents: do not import a library outside this list without explicit
approval. If a new dependency is needed, stop and ask.

## Phase 4 - Conditional Domain Specs

Write only the sections that match the classified repo type. Mark non-applicable
sections as `N/A` only when that avoids ambiguity.

Common conditional specs:

| Lane | Use When | Required Output |
| --- | --- | --- |
| Database Spec | Repo stores durable structured data. | Schema, relationships, indexes, migrations, backup/restore. |
| API Contract | Repo exposes HTTP/RPC/GraphQL/gRPC/tRPC/MCP or IPC contracts. | Interface catalog, auth, envelope, errors, versioning. |
| Frontend / UX Spec | Repo has user-facing UI. | Routes/screens, states, design system, accessibility, responsive behavior. |
| Library / SDK Spec | Repo publishes reusable code. | Public API, semver, compatibility, examples, package release. |
| CLI Spec | Repo exposes commands. | Command catalog, flags, stdin/stdout/stderr, exit codes, config, completion. |
| Infrastructure Spec | Repo manages infrastructure. | Topology, state backend, environments, secrets, drift, rollback. |
| Data / ML Spec | Repo processes data or trains/serves models. | Datasets, pipeline DAG, metrics, reproducibility, model registry. |
| Mobile Spec | Repo ships mobile app. | Platform matrix, permissions, navigation, offline, store release. |
| Browser Extension Spec | Repo ships extension. | Manifest, permissions, content/background scripts, CSP, store release. |
| MCP / Plugin Spec | Repo exposes agent/plugin tools. | Tool schema, resources, prompts, transport, permissions, integration tests. |
| Documentation Spec | Repo is docs/content. | IA, source format, style guide, publishing, review workflow. |
| Embedded / Firmware Spec | Repo targets hardware. | Board/MCU, toolchain, flashing, IO, safety, test rigs. |
| Game / Graphics Spec | Repo is game/3D/simulation. | Engine, loop, asset pipeline, input, perf budget, platforms. |

Read `references/repo-type-lanes.md` for detailed templates.

## Phase 5 - Technical Standards

Define how the repo is maintained. Use `references/technical-standards.md` for
full patterns when needed.

Output:

```markdown
## Technical Standards

- Git workflow:
- Commit convention:
- Branch / PR rules:
- Error handling:
- Logging / tracing:
- Configuration:
- Secrets:
- Caching / invalidation, if applicable:
- Background jobs, if applicable:
- Concurrency / conflict resolution:
- Security baseline:
- Code organization:
- Generated code policy:
- Documentation policy:
```

Sign-off rule: if multiple roles are involved, architecture, backend, frontend,
DevOps, QA, and product owners should review the standards before coding starts.

## Phase 6 - Environment, Packaging, Release

Every repo needs a way to run or consume its deliverable. This is not always
deployment; for libraries it may be publishing, for docs it may be static-site
generation, for firmware it may be flashing.

Output:

```markdown
## Environment, Packaging, Release

### Environments / Targets
| Target | Purpose | Build/source | Notes |
| --- | --- | --- | --- |

### Configuration
- Environment variables:
- Config files:
- Secret handling:

### Local Setup
[Commands that work from clone to first successful run/build.]

### Build / Package / Publish
- Build command:
- Package artifact:
- Release channel:
- Versioning:

### CI/CD or Automation
- Trigger:
- Steps:
- Required checks:
- Rollback / restore:
```

## Phase 7 - Test & QA Strategy

Tests are verification. They do not replace correct behavior, runtime review, or
QA evidence.

Output:

```markdown
## Test & QA Strategy

- Test levels:
- Critical behavior to cover:
- Fixtures / data strategy:
- Integration or contract tests:
- E2E or runtime validation:
- Performance / capacity checks:
- Security checks:
- Accessibility / visual checks, if UI exists:
- CI gates:
- Manual QA requirements:
```

For UI/runtime projects, include visible user-flow QA requirements. If the repo
will use the Anvien QA skill, state which QA lanes are expected.

## Phase 8 - Handoff Checklist

Before coding starts, verify that all code-shaping decisions are either resolved
or explicitly blocked.

```markdown
## Handoff Checklist

- [ ] Repository classification completed.
- [ ] Project context and scope completed.
- [ ] Architecture / repo blueprint completed.
- [ ] Toolchain and dependency policy completed.
- [ ] Required conditional domain specs completed.
- [ ] Non-applicable lanes explicitly marked when needed.
- [ ] Technical standards completed.
- [ ] Environment/package/release path completed.
- [ ] Test and QA strategy completed.
- [ ] No code-shaping decision remains implicit.
- [ ] All blocking TBD items are listed with owner.

Handoff: <architect|supervisor|coder> - <reason>
```

Handoff roles:
- `architect`: unresolved architecture, source-of-truth, system-flow, or rule decisions.
- `supervisor`: needs priority, scope, acceptance, or coordination decision.
- `coder`: enough decisions are made to begin implementation.

## How To Work With The User

- Interview before writing. Do not fill code-shaping decisions from guesses.
- Ask for the smallest number of high-impact missing decisions at a time.
- If the user has not chosen a repository type, classify it with evidence from
  the request.
- If a lane is not applicable, say why.
- If a required decision is unknown, mark it as blocked instead of inventing it.
- Keep output as Markdown source-of-truth docs, usually `SPEC.md` or files under
  `docs/`.

## Red Flags

Stop and ask or mark blocked when:

- repository type is unknown and affects the workflow
- the skill is about to force DB/API/frontend sections onto a repo that does not need them
- dependencies are unapproved or version is "latest"
- public API/CLI/tool surface is unspecified
- data persistence exists but schema/migration policy is missing
- release/distribution target is missing
- tests are defined before expected behavior is clear
- the handoff would still require the coder to guess architecture or contracts
