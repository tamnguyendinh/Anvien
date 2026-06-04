# Repo Type Lanes Reference

Use this reference after `Repository Type Classification`. Load only the lanes
that apply to the repo being bootstrapped.

## App / Service

Use for web apps, backend services, fullstack products, SaaS products, and
runtime systems.

Required outputs:

- route/page/service map
- API or RPC contract, if external or frontend consumers exist
- database/read-model spec, if durable data exists
- auth/session/permission model
- background job/event model, if async work exists
- deployment/runtime target
- runtime QA strategy

## Library / SDK

Use for reusable packages, frameworks, internal libraries, and SDKs.

Required outputs:

```markdown
## Library / SDK Spec

- Package name:
- Public API surface:
- Supported languages/runtimes:
- Compatibility matrix:
- Semver policy:
- Breaking-change policy:
- Error model:
- Examples and quickstart:
- Documentation requirements:
- Package publishing target:
- Test matrix:
- Deprecation policy:
```

Do not require DB, HTTP API, or deployment sections unless the library actually
includes them.

## CLI

Use for command-line tools and developer utilities.

Required outputs:

```markdown
## CLI Spec

- Binary name:
- Command catalog:
- Flags and arguments:
- stdin/stdout/stderr behavior:
- Exit codes:
- Config files and precedence:
- Environment variables:
- Shell completion:
- Interactive vs non-interactive behavior:
- Logging and verbosity:
- Packaging/distribution:
- Golden tests / snapshot tests:
```

Every command should have examples and error behavior.

## Infrastructure / IaC

Use for Terraform, Pulumi, CloudFormation, Kubernetes, Helm, platform, and ops
repos.

Required outputs:

```markdown
## Infrastructure Spec

- Cloud/provider scope:
- Environment topology:
- State backend:
- Module/resource boundaries:
- Secrets management:
- Network/security model:
- Policy-as-code:
- Drift detection:
- Plan/apply workflow:
- Rollback/restore:
- Cost controls:
- CI/CD gates:
```

Do not treat infrastructure repos like app repos unless they also ship an app.

## Data / ML

Use for analytics pipelines, ETL/ELT, ML training, model serving, and data
science repos.

Required outputs:

```markdown
## Data / ML Spec

- Data sources:
- Dataset/versioning strategy:
- Pipeline DAG:
- Feature definitions:
- Training/evaluation split:
- Metrics and acceptance thresholds:
- Reproducibility requirements:
- Experiment tracking:
- Model registry:
- Serving interface:
- Privacy/PII rules:
- Data quality checks:
```

If the repo only analyzes data and does not deploy a service, API/deployment can
be N/A.

## Mobile

Use for iOS, Android, React Native, Flutter, Kotlin Multiplatform, or other
mobile-first apps.

Required outputs:

```markdown
## Mobile Spec

- Platform matrix:
- Navigation map:
- Permission model:
- Offline/cache behavior:
- Push notifications:
- Device/storage/security constraints:
- Accessibility:
- App store/release path:
- Crash reporting:
- Device test matrix:
```

## Browser Extension

Use for Chrome/Edge/Firefox extensions.

Required outputs:

```markdown
## Browser Extension Spec

- Manifest version:
- Permissions:
- Content scripts:
- Background/service worker:
- Popup/options pages:
- Message passing:
- CSP/security constraints:
- Host permissions:
- Storage model:
- Browser compatibility:
- Store release path:
```

## MCP / Plugin

Use for MCP servers, Codex/Claude plugins, tool integrations, agent capabilities,
and connector repos.

Required outputs:

```markdown
## MCP / Plugin Spec

- Tool/resource/prompt catalog:
- Input/output schemas:
- Transport:
- Auth/permission model:
- Rate limits and safety rules:
- Error model:
- Host integration:
- Local setup:
- Contract tests:
- Example prompts:
- Versioning and compatibility:
```

For MCP tools, schema clarity is the API contract.

## Documentation

Use for docs-only, wiki, playbook, standards, and content repos.

Required outputs:

```markdown
## Documentation Spec

- Audience:
- Information architecture:
- Source format:
- Style guide:
- Terminology:
- Review workflow:
- Publishing target:
- Link/check automation:
- Versioning:
- Ownership:
```

Do not require DB/API/build runtime unless the docs site has an app layer.

## Embedded / Firmware

Use for firmware, microcontroller, hardware, robotics, or device repos.

Required outputs:

```markdown
## Embedded / Firmware Spec

- Target board/MCU:
- Toolchain:
- Flashing/debugging:
- Hardware IO map:
- Timing and power constraints:
- Safety constraints:
- Serial/logging protocol:
- Update mechanism:
- Test rig:
- Hardware-in-loop strategy:
```

## Game / Graphics

Use for games, Three.js/WebGL/WebGPU apps, simulations, 3D tools, and graphics
repos.

Required outputs:

```markdown
## Game / Graphics Spec

- Engine/framework:
- Game loop or render loop:
- Scene/entity architecture:
- Asset pipeline:
- Input model:
- Physics/collision:
- Save/state model:
- Performance budget:
- Platform targets:
- Visual QA strategy:
```

## Desktop App

Use `references/desktop-app.md` for detailed desktop-specific architecture,
security, packaging, update, and test guidance.

Minimum outputs:

- framework choice
- main/core process responsibilities
- renderer/UI responsibilities
- IPC contract
- local data strategy
- OS integration
- signing/update/distribution
- desktop test strategy

## Monorepo

Use when the repo contains multiple packages/apps/services.

Required outputs:

```markdown
## Monorepo Spec

- Workspace tool:
- Package graph:
- Ownership boundaries:
- Shared contracts:
- Build orchestration:
- Dependency rules:
- Versioning policy:
- Release strategy:
- Cross-package test strategy:
- Generated code policy:
```

## SaaS Multi-Tenant

Use `references/saas-patterns.md` for tenant isolation, entitlement, billing, and
admin governance.

Minimum outputs:

- tenant/account model
- ownership tuple
- entitlement model
- billing lifecycle
- tenant isolation
- audit logs

## Design System

Use `references/design-system.md` when the repo has complex UI or reusable
frontend components.

Minimum outputs:

- tokens
- components
- accessibility rules
- responsive breakpoints
- theming
- visual QA expectations

## Blockchain / Web3

Use when smart contracts, wallet flows, chain data, or on-chain/off-chain
systems are in scope.

Required outputs:

```markdown
## Blockchain / Web3 Spec

- Chain/network targets:
- Smart contract standards:
- On-chain vs off-chain data decision:
- Wallet integration:
- Transaction lifecycle:
- Gas/performance constraints:
- Contract audit requirements:
- Testnet/mainnet release path:
- Fork/mainnet test strategy:
```
