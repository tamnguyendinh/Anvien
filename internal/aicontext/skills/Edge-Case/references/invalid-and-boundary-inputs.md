# Attack Guide: Invalid & Boundary Inputs

> This file is part of Edge-Case Review Skill v2. Use when: testing boundary values, null/nil handling, extreme payload sizes, special characters, or malformed data inputs.

---

## Attack Surface Overview

Unvalidated boundaries cause unhandled panics, memory overflows, database constraint crashes, and unexpected truncation. Robust systems enforce schema validation and sanitize inputs strictly at the boundary.

---

## Targeted Invariants

- **Boundary Validation**: Values outside defined numerical, string-length, or temporal ranges must be rejected before business logic execution.
- **Nil / Null Safety**: Missing fields, null pointers, and undefined values must never cause unhandled panics, server crashes (HTTP 500), or blank UI crashes.
- **Payload Size Limits**: Massive strings, giant arrays, or deep JSON nesting must be blocked by rate limiters and payload size guards.
- **Sanitized Encoding**: Unicode, emojis, control characters, and SQL/HTML injection payloads must be safely handled or rejected.

---

## Hostile Perturbations & Attack Scenarios

### 1. Extreme Numerical Boundaries
- Submit integer values: `0`, `-1`, `-2147483648`, `2147483647`, `9223372036854775807`, `1e308`, `NaN`, `Infinity`.
- *Check:* Does the system overflow, divide by zero, or store corrupted numbers?

### 2. String & Unicode Fuzzing
- Submit:
  - Empty string `""`, whitespace-only `"   "`.
  - 100,000-character long strings.
  - Multi-byte UTF-8, zero-width joiners, right-to-left override characters, emojis (e.g. `👨‍👩‍👧‍👦`).
  - SQL / Command injection strings (`' OR 1=1 --`, `$(rm -rf /)`).
- *Check:* Does the UI render cleanly without layout breaks? Does the database store strings without truncation errors?

### 3. Missing / Extra JSON Keys & Malformed Payloads
- Send JSON payloads with:
  - Required fields omitted (`{}`).
  - Explicit `null` for non-nullable keys (`{"name": null}`).
  - Extra unrecognized fields (`{"__proto__": {"admin": true}}`).
  - Truncated or malformed JSON (`{"key": `).
- *Check:* Does the parser fail closed with clean HTTP 400 Bad Request?

### 4. Deeply Nested Object Injection
- Send JSON with 500 levels of nested objects/arrays (`{"a":{"a":{"a":...}}}`).
- *Check:* Does the server crash with a stack overflow or parse error?

---

## Pass / Fail Checklist

- [ ] All inputs are strictly validated at the outer boundary.
- [ ] Null or missing values never trigger unhandled server panics or client blank screens.
- [ ] Payloads exceeding size or depth limits are rejected with HTTP 400/413.
- [ ] Special characters and multi-byte UTF-8 do not corrupt database records.
