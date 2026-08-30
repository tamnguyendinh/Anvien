# When Stuck — Problem-Solving Dispatch

> use when: stuck on a problem and unsure which specific technique to apply for your symptom.

---

## Overview

Different stuck-types need different techniques. This guide helps you quickly identify and apply the right problem-solving method.

**Core Principle:** Match the stuck-symptom directly to the technique.

---

## Quick Dispatch Matrix

| Symptom / How You're Stuck | Root Cause | Recommended Technique | Reference File |
|---|---|---|---|
| **Complexity spiraling** | Same thing implemented 5+ ways, growing special cases | Simplification Cascades | `references/simplification-cascades.md` |
| **Need innovation** | Conventional solutions inadequate, can't find fitting approach | Collision-Zone Thinking | `references/collision-zone-thinking.md` |
| **Forced by assumptions** | "Must be done this way", can't question premise | Inversion Exercise | `references/inversion-exercise.md` |
| **Recurring patterns** | Same issue in different places, reinventing wheels | Meta-Pattern Recognition | `references/meta-pattern-recognition.md` |
| **Scale uncertainty** | Will it work in production? Edge cases unclear? | Scale Game | `references/scale-game.md` |
| **Code broken** | Wrong behavior, test failing, unexpected output | Systematic Debugging | `skills/debugging/SKILL.md` |

---

## Dispatch Decision Flow

```text
Are you stuck on a problem?
  │
  ├── [Complexity spiraling / 5+ ways / special cases?] ────► references/simplification-cascades.md
  │
  ├── [Conventional solutions inadequate / need breakthrough?] ► references/collision-zone-thinking.md
  │
  ├── [Forced into "only one way" / rigid assumptions?] ────► references/inversion-exercise.md
  │
  ├── [Same pattern across 3+ domains / feels familiar?] ───► references/meta-pattern-recognition.md
  │
  └── [Unsure of scale limits / 1000x volume / edge cases?] ─► references/scale-game.md
```

---

## Systematic Execution Process

1. **Identify stuck-type**: What symptom matches the matrix above?
2. **Open specific reference**: Read and follow that technique's step-by-step process.
3. **Apply technique**: Execute the thought experiment or restructuring.
4. **If still stuck**: Combine techniques or switch perspectives.

---

## Combining Techniques

Complex architectural or system challenges often benefit from combining two techniques:

- **Simplification + Meta-Pattern**: Find the universal pattern across domains, then simplify all local instances into one abstraction.
- **Collision + Inversion**: Force a cross-domain metaphor, then invert its core assumptions to find unexploited capabilities.
- **Scale + Simplification**: Push the system to 1000x extreme scale to instantly reveal which components are redundant and can be eliminated.
