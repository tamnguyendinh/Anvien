# Simplification Cascades

> use when: implementing the same concept multiple ways, accumulating special cases, or when system complexity is spiraling.

---

## Overview

Sometimes one insight eliminates 10 things. Look for the unifying principle that makes multiple components unnecessary.

**Core Principle:** *"Everything is a special case of..."* collapses complexity dramatically.

---

## Quick Reference

| Symptom | Likely Cascade |
|---|---|
| Same thing implemented 5+ ways | Abstract the common pattern into one unified handler |
| Growing special case list | Find the general case that naturally covers the edge cases |
| Complex rules with exceptions | Find the higher-order rule that has zero exceptions |
| Excessive config options | Find sensible defaults that eliminate 95% of switches |

---

## The Pattern

**Look for:**
- Multiple implementations of similar concepts
- Special case handling everywhere
- "We need to handle A, B, C, D differently..."
- Complex conditional rules with many exceptions

**Ask:** *"What if they're all the exact same thing underneath?"*

---

## Real-World Examples

### Cascade 1: Stream Abstraction
- **Before:** Separate custom handlers for batch, real-time, file, and network data.
- **Insight:** "All inputs are streams — just originating from different sources."
- **After:** One stream processor pipeline, multiple lightweight source adapters.
- **Eliminated:** 4 separate processing architectures.

### Cascade 2: Resource Governance
- **Before:** Session tracking, rate limiting, file upload validation, and connection pooling all handled separately.
- **Insight:** "All of these are simply per-entity resource limits."
- **After:** One unified `ResourceGovernor` managing different resource types.
- **Eliminated:** 4 custom enforcement subsystems.

### Cascade 3: Immutability
- **Before:** Defensive copying, locking mechanisms, complex cache invalidation, and temporal coupling.
- **Insight:** "Treat everything as immutable data + pure state transformations."
- **After:** Functional state pipelines.
- **Eliminated:** Entire classes of race conditions and synchronization bugs.

---

## Process

1. **List the variations**: What is currently implemented in multiple separate ways?
2. **Find the essence**: What is genuinely identical underneath the surface?
3. **Extract the abstraction**: What domain-independent pattern unifies them?
4. **Test the fit**: Do all current and future cases fit cleanly without hacks?
5. **Measure the cascade**: How many lines of code, classes, and tests become obsolete and can be deleted?

---

## Red Flags You Need a Simplification Cascade

- "We just need to add one more special case..." (repeating continuously).
- "These modules are all similar but different" (they are usually the same).
- Refactoring feels like whack-a-mole (fixing one case breaks another).
- Unbounded growth of configuration files.
- "Don't touch that module, it's complicated" (complexity hiding a missing abstraction).

---

## Key Invariant

> **Simplification cascades produce 10x wins, not 10% incremental tweaks.**  
> *One powerful abstraction > ten clever hacks.*  
> Measure success by: *"How many unnecessary components can we completely delete?"*
