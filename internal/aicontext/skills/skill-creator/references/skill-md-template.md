# SKILL.md Template

Use this template when creating a new `SKILL.md`. Keep only sections that change the agent's behavior.

```md
---
name: skill-name
description: This skill should be used when [precise trigger under 200 characters].
---

# Skill Name

## Purpose

State what this skill helps the agent do.

State what the skill is not for if the boundary prevents wrong activation.

## Operating Rules

- Add rules that change agent behavior.
- Include exceptions when a rule is not absolute.
- Remove any sentence that is only mostly correct.

## Workflow

1. Identify the user's requested action.
2. Inspect the relevant source, files, tools, or context.
3. Choose the smallest workflow that satisfies the task.
4. Validate the result with evidence appropriate to the risk.

## Reference Navigation

- `references/example.md`: Load only when [specific condition].
```

## Writing Gate

Before keeping a sentence, verify:

- It makes the agent choose the correct action.
- It is true for every covered case, or the exception is explicit.
- It does not accidentally forbid valid user-requested work.
- It does not accidentally authorize risky work.
- It does not conflict with another rule.

Prefer concise wording, but never sacrifice precision, scope, or exceptions for brevity.
