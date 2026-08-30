# Meta-Pattern Recognition

> use when: noticing the same problem/solution structure across 3+ different domains or experiencing architectural déjà vu.

---

## Overview

When the exact same pattern appears across 3+ distinct domains, it is a universal principle worth formalizing into a shared architectural primitive.

**Core Principle:** Discover patterns in how patterns emerge.

---

## Quick Reference

| Pattern In Context | Abstract Universal Form | Other High-Leverage Applications |
|---|---|---|
| CPU / DB / HTTP / DNS caching | Store frequently accessed data closer to compute | LLM prompt caching, local storage indexing |
| Network / Storage / OS layering | Separate concerns into strict abstraction tiers | Clean architecture, security perimeter defense |
| Message queues / Event streams | Decouple producer from consumer via buffer | Background task processing, UI state dispatch |
| Connection / Thread / Object pools | Amortize allocation cost of expensive resources | Memory management, API client lifecycle |

---

## Process

1. **Spot repetition**: Notice structural similarities appearing in 3+ unrelated areas.
2. **Extract abstract form**: Describe the mechanics completely independent of any specific domain or technology.
3. **Identify variation points**: What parameters, limits, or adapters vary per domain?
4. **Check applicability**: Where else in our system can this universal primitive resolve unmanaged complexity?

---

## Example

- **Pattern spotted:** API rate limiters, network traffic shaping, circuit breakers, server admission control.
- **Abstract form:** Bound resource consumption strictly at the boundary to prevent system-wide exhaustion.
- **Variation points:** Monitored resource type, consumption window, breach fallback action.
- **New application:** AI context window management — enforce a token budget governor to prevent LLM context overflows before calling models.

---

## Red Flags You're Missing a Meta-Pattern

- "Our problem is 100% unique and has never been solved before" (almost certainly false).
- Multiple engineers independently inventing custom solutions to the same underlying problem.
- Reinventing foundational architectural wheels.
- Sensation of déjà vu when designing subsystems.

---

## Key Invariant

> 3+ domains exhibiting the same pattern = universal, battle-tested principle.  
> Extracting the abstract form prevents reinventing bespoke, fragile mechanisms.
