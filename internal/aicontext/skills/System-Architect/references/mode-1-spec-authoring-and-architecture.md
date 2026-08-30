# Mode 1 — SPEC Authoring And Architecture Design

> This file is part of System-Architect Skill v2. Read when: in Mode 1 to design system architecture, author/synchronize SPECs, draft ADRs, or evaluate architectural trade-offs.

---

## Hard Rule — Supreme Design Principle

**FORBIDDEN to build MVP.** Design must target production-ready from the start.
- All SPECs must be written for production — no temporary writing, no "will supplement after launch".
- Do not create SPEC authority that only works for a demo, pilot, single happy-path rollout, or short-lived stopgap and assumes a later redesign.
- If delivery must be phased, phase the implementation on top of production-grade architecture; do not phase the architecture down into MVP form.
- Every architecture decision must be safe for the full system lifecycle: rollout, growth, failure, recovery, operations, maintenance, and long-term ownership.
- `Security`, `error handling`, `monitoring`, and `logging` must be designed in from the start, not deferred as a later hardening pass.
- If there is insufficient information to design production-ready -> ask/research more, must not reduce scope to MVP.
- If production-safe coverage is not ready yet, remain in `Mode 1` and finish the authority. Do not jump to `Mode 2` just to produce a faster plan.

---

## Role & Responsibilities

- Design system architecture for new features
- Evaluate technical trade-offs
- Recommend patterns and best practices
- Identify scalability bottlenecks
- Plan for future growth
- Ensure consistency across codebase
- Clarify architecture direction only when existing SPEC/ADR authority is missing, contradictory, or needs a new standardized decision

---

## SPEC Boundary Rule

- SPEC is for architecture authority, boundaries, contracts, invariants, and forbidden patterns.
- SPEC should define what must be true, which layer owns the behavior, and which runtime contract must hold.
- SPEC should not drift into low-level coding prescription unless the detail itself is the contract surface.
- Avoid treating the following as SPEC authority by default:
  - function names
  - variable names
  - helper names
  - exact internal file splits
  - exact refactor choreography
- When producing architecture guidance, separate clearly:
  - `Architecture / SPEC rule`
  - `Implementation suggestion`
- If a recommendation is only one possible way to code the solution, label it as an implementation suggestion rather than architecture law.

### HARD RULE: FORBIDDEN to write SPEC containing specific function names or variable names.

---

## Terminology Resolution Rule

- Shared architecture prompts may use generic scope wording as cross-repo placeholders.
- Each repo must resolve generic scope wording to its domain-scoped entity from exact SPEC authority.
- For the target repo, use `AGENTS.md` and authoritative SPEC files for app type and scope identifier mapping.
- The target repo's owner identifier must follow `AGENTS.md` and authoritative SPEC files and must not be remapped outside that contract.
- Generic wording differences alone are not architectural conflict; conflict exists only when mapping is ambiguous or breaks ownership/isolation/runtime contracts.

---

## Architecture Review Process

### 1. Current State Analysis
- Review existing architecture
- Identify patterns and conventions
- Document technical debt
- Assess scalability limitations

### 2. Requirements Gathering
- Functional requirements
- Non-functional requirements (performance, security, scalability)
- Integration points
- Data flow requirements

### 3. Design Proposal
- High-level architecture diagram
- Component responsibilities
- Data models
- API contracts
- Integration patterns

### 4. Trade-Off Analysis
For each design decision, document:
- **Pros**: Benefits and advantages
- **Cons**: Drawbacks and limitations
- **Alternatives**: Other options considered
- **Decision**: Final choice and rationale

---

## Architectural Principles

### 1. Modularity & Separation of Concerns
- Single Responsibility Principle
- High cohesion, low coupling
- Clear interfaces between components
- Independent deployability

### 2. Scalability
- Horizontal scaling capability
- Stateless design where possible
- Efficient database queries
- Caching strategies
- Load balancing considerations

### 3. Maintainability
- Clear code organization
- Consistent patterns
- Comprehensive documentation
- Easy to test
- Simple to understand

### 4. Security
- Defense in depth
- Principle of least privilege
- Input validation at boundaries
- Secure by default
- Audit trail

### 5. Performance
- Efficient algorithms
- Minimal network requests
- Optimized database queries
- Appropriate caching
- Lazy loading

---

## Common Patterns

### Frontend Patterns
- **Component Composition**: Build complex UI from simple components
- **Container/Presenter**: Separate data logic from presentation
- **Custom Hooks**: Reusable stateful logic
- **Context for Global State**: Avoid prop drilling
- **Code Splitting**: Lazy load routes and heavy components

### Backend Patterns
- **Repository Pattern**: Abstract data access
- **Service Layer**: Business logic separation
- **Middleware Pattern**: Request/response processing
- **Event-Driven Architecture**: Async operations
- **CQRS**: Separate read and write operations

### Data Patterns
- **Normalized Database**: Reduce redundancy
- **Denormalized for Read Performance**: Optimize queries
- **Event Sourcing**: Audit trail and replayability
- **Caching Layers**: Redis, CDN
- **Eventual Consistency**: For distributed systems

---

## Architecture Decision Records (ADRs)

For significant architectural decisions, create ADRs:

```markdown
# ADR-001: Use Redis for Semantic Search Vector Storage

## Context
Need to store and query 1536-dimensional embeddings for semantic market search.

## Decision
Use Redis Stack with vector search capability.

## Consequences

### Positive
- Fast vector similarity search (<10ms)
- Built-in KNN algorithm
- Simple deployment
- Good performance up to 100K vectors

### Negative
- In-memory storage (expensive for large datasets)
- Single point of failure without clustering
- Limited to cosine similarity

### Alternatives Considered
- **PostgreSQL pgvector**: Slower, but persistent storage
- **Pinecone**: Managed service, higher cost
- **Weaviate**: More features, more complex setup

## Status
Accepted

## Date
YYYY-MM-DD
```

---

## System Design Checklist

When designing a new system or feature:

### Functional Requirements
- [ ] User stories documented
- [ ] API contracts defined
- [ ] Data models specified
- [ ] UI/UX flows mapped

### Non-Functional Requirements
- [ ] Performance targets defined (latency, throughput)
- [ ] Scalability requirements specified
- [ ] Security requirements identified
- [ ] Availability targets set (uptime %)

### Technical Design
- [ ] Architecture diagram created
- [ ] Component responsibilities defined
- [ ] Data flow documented
- [ ] Integration points identified
- [ ] Error handling strategy defined
- [ ] Testing strategy planned

### Operations
- [ ] Deployment strategy defined
- [ ] Monitoring and alerting planned
- [ ] Backup and recovery strategy
- [ ] Rollback plan documented

---

## Red Flags & Anti-Patterns

Watch for these architectural anti-patterns:
- **Big Ball of Mud**: No clear structure
- **Golden Hammer**: Using same solution for everything
- **Premature Optimization**: Optimizing too early
- **Not Invented Here**: Rejecting existing solutions
- **Analysis Paralysis**: Over-planning, under-building
- **Magic**: Unclear, undocumented behavior
- **Tight Coupling**: Components too dependent
- **God Object**: One class/component does everything
- **SPEC-as-code-style**: Architecture docs forcing exact low-level implementation where multiple compliant implementations are possible

---

## Output Definition — 10 SPEC Types

This lane must produce at least 10 output types (may be more depending on the actual situation). Each type is 1 file or 1 group of separate files.

### 1. Blueprint
Each file must contain:
- What the system is, who it serves
- Position in the ecosystem (relationship with other systems)
- Boundary between components
- Data flow — which direction, who pushes who pulls
- Connection rules (independent / dependent / one-way / bidirectional)
- Conditions: if system A goes down, what impact on B

### 2. DB SPEC
Each file must contain:
- Schema, tables, relationships between tables
- Index strategy
- Migration strategy (versioning, rollback)
- Naming convention for tables, columns, constraints
- Constraint rules (unique, foreign key, check)

### 3. Tech Stack SPEC
Each file must contain:
- Framework, language, runtime — with version
- Main libraries — with reason for selection
- "DO NOT USE" list — with reason for exclusion
- Library selection principles

### 4. Coding Patterns SPEC
Each file must contain:
- Rules by boundary (which layer owns what)
- Contract between layers
- Anti-pattern for each rule
- MUST NOT contain specific function names, variable names, file naming conventions

### 5. UI/UX SPEC
Reference `.agent/skills/ui-ux-pro-max-skill-main` when designing UI/UX SPEC.
Each file must contain:
- User flow
- Screen hierarchy
- Interaction rules (behavior on click, submit, error)
- Responsive rules
- Accessibility contract

### 6. Architecture Contract
Each file must contain:
- API contract definition between systems (endpoints, payload structure, auth method)
- Error codes and error handling contract
- Versioning strategy for API
- Backward compatibility rules

### 7. Security SPEC
Each file must contain:
- Auth flow (login, session, token lifecycle)
- TLS policy
- Rate limiting rules
- Data protection (PII handling, encryption at rest/in transit)
- Forbidden practices (prohibited list)

### 8. Infrastructure SPEC
Each file must contain:
- Deployment topology (services, ports, networking)
- Docker / container configuration
- Environment config (dev, staging, production)
- Backup / restore strategy
- Monitoring / alerting requirements

### 9. Testing Requirements
Each file must contain:
- Coverage targets per layer (unit, integration, E2E)
- Critical paths that must have tests
- Test strategy per type (auth, billing, distribution, admin, public content...)
- Performance benchmarks / acceptance criteria

### 10. Logging SPEC
Each file must contain:
- Log levels and when to use each level
- Structured log format
- Correlation ID strategy (cross-system tracing)
- Sensitive data rules (what is forbidden to log)
- Log retention / rotation policy

---

## Output Rules

- Each output type = 1 separate file or 1 group of separate files if the original file is too long and split into multiple smaller parts. Do not merge multiple types into 1 file.
- Must not contain specific function names, variable names, or file naming conventions in SPEC.
- Only contains boundary, contract, invariant, and forbidden patterns.
- Each SPEC must have a clear status: `IDEA` / `DRAFT` / `APPROVED`.
- When producing guidance, must clearly separate: `Architecture / SPEC rule` vs `Implementation suggestion`.
- Output is SPEC only — Mode 1 does not produce `AGENTS.md` or execution planning docs.

---

## SPEC File Splitting Rules

- Each SPEC file must not exceed **800 lines**.
- If exceeding 800 lines -> split into multiple Parts.
- Split by **content / functionality**, DO NOT cut across a document.
- Each Part must be **self-contained** — readable independently to understand the work content.
- Naming: `<SPEC-Name>-Part-<X>-<content>.md`

Example:
```text
Blueprint-Part-A-<content>.md
Blueprint-Part-B-<content>.md
TECH-STACK-SPEC-Part-A-<content>.md
TECH-STACK-SPEC-Part-B-<content>.md
```
