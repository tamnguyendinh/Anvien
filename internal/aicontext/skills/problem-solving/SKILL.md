---
name: problem-solving
description: Use when the user asks to solve a hard problem, break through stuck points, eliminate spiraling complexity, or find elegant solutions beyond conventional approaches.
---

# Problem-Solving Techniques

A structured toolkit of high-leverage cognitive and architectural techniques to break through stuck points, eliminate spiraling complexity, and discover elegant solutions.

## Core Philosophy

> **"One powerful abstraction > ten clever hacks"**  
> Rather than managing complexity through brute force, these techniques help you uncover the fundamental principles that make complexity unnecessary.

---

## Symptom Matcher & Dispatch Table

When encountering a difficult technical problem or blocker, match your symptom to the appropriate technique:

| Symptom / How You're Stuck | Recommended Technique | Target Reference |
|---|---|---|
| **Unsure which method to apply** | Stuck-Type Diagnostic Dispatch | `references/when-stuck-dispatch.md` |
| **Complexity spiraling / Same logic 5+ ways** | Simplification Cascades | `references/simplification-cascades.md` |
| **Conventional solutions inadequate / Need innovation** | Collision-Zone Thinking | `references/collision-zone-thinking.md` |
| **Rigid assumptions / "Only one way" mindset** | Inversion Exercise | `references/inversion-exercise.md` |
| **Recurring pattern across 3+ domains** | Meta-Pattern Recognition | `references/meta-pattern-recognition.md` |
| **Uncertain about production limits & extreme load** | Scale Game | `references/scale-game.md` |

---

## Problem-Solving Workflow

```text
                  [Technical Blocker / Stuck Point]
                                 │
                                 ▼
                     [Identify Stuck Symptom]
                                 │
         ┌───────────────────────┼───────────────────────┐
         │                       │                       │
         ▼                       ▼                       ▼
[Complexity Spiraling]  [Need Breakthrough]      [Rigid Assumptions]
references/simplification-  references/collision-  references/inversion-
cascades.md             zone-thinking.md         exercise.md
         │                       │                       │
         ├───────────────────────┼───────────────────────┤
         │                       │                       │
         ▼                       ▼                       ▼
[Recurring Pattern]     [Extreme Scale Limits]   [Unsure Where to Start]
references/meta-pattern- references/scale-        references/when-stuck-
recognition.md          game.md                  dispatch.md
                                 │
                                 ▼
                 [Apply Technique & Extract Invariant]
                                 │
                                 ▼
             [Synthesize Clean Architectural Solution]
```

---

## Quick Decision Tree

- Not sure which technique fits your blocker? → `references/when-stuck-dispatch.md`
- Implementing the same feature multiple ways or accumulating special cases? → `references/simplification-cascades.md`
- Out of ideas within the local domain and need cross-domain metaphor breakthroughs? → `references/collision-zone-thinking.md`
- Forcing a solution that feels wrong because of unquestioned assumptions? → `references/inversion-exercise.md`
- Solving a problem that feels structurally familiar to something in networking, OS, or biology? → `references/meta-pattern-recognition.md`
- Need to stress-test an architecture before writing production code? → `references/scale-game.md`

---

## Reference Index

| Technique | Core Mission | Reference File |
|---|---|---|
| **When Stuck Dispatch** | Symptom-to-technique routing matrix and combined methods | `references/when-stuck-dispatch.md` |
| **Simplification Cascades** | Finding unifying abstractions that eliminate 10 redundant components | `references/simplification-cascades.md` |
| **Collision-Zone Thinking** | Forcing unrelated concepts together to generate novel properties | `references/collision-zone-thinking.md` |
| **Inversion Exercise** | Flipping foundational assumptions to expose hidden constraints | `references/inversion-exercise.md` |
| **Meta-Pattern Recognition** | Extracting universal primitives from patterns found in 3+ domains | `references/meta-pattern-recognition.md` |
| **Scale Game** | Testing at 1000x extremes to discover breaking points and resilient cores | `references/scale-game.md` |
