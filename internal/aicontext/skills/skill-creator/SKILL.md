---
name: skill-creator
description: Use when the user asks to create or update a skill.
license: Complete terms in LICENSE.txt
version: 2.0.0
---

# Skill Creator

This skill provides guidance for creating effective skills.

## About Skills

Skills are modular, self-contained packages that extend Claude's capabilities by providing
specialized knowledge, workflows, and tools. Think of them as "onboarding guides" for specific
domains or tasks—they transform Claude from a general-purpose agent into a specialized agent
equipped with procedural knowledge that no model can fully possess.

**IMPORTANT:**
- Skills are not documentation, they are practical instructions for Claude Code to use the tools, packages, plugins or APIs to achieve the tasks.
- Each skill teaches Claude how to perform a specific development task, not what a tool does.
- Claude Code can activate multiple skills automatically to achieve the user's request.

### What Skills Provide

1. Specialized workflows - Multi-step procedures for specific domains
2. Tool integrations - Instructions for working with specific file formats or APIs
3. Domain expertise - Company-specific knowledge, schemas, business logic
4. Bundled resources - References and assets for complex and reusable task context

### Anatomy of a Skill

Every skill consists of a required SKILL.md file and optional bundled resources:

```
.claude/skills/
└── skill-name/
    ├── SKILL.md (required)
    │   ├── YAML frontmatter metadata (required)
    │   │   ├── name: (required)
    │   │   ├── description: (required)
    │   │   ├── license: (optional)
    │   │   └── version: (optional)
    │   └── Markdown instructions (required)
    └── Bundled Resources (optional)
        ├── references/       - Documentation intended to be loaded into context as needed
        └── assets/           - Files used in output (templates, icons, fonts, etc.)
```

#### Requirements (**IMPORTANT**)

- Skill should be combined into specific topics, for example: `cloudflare`, `cloudflare-r2`, `cloudflare-workers`, `docker`, `gcloud` should be combined into `devops`
- `SKILL.md` should be **less than 150 lines** and include the references of related markdown files and assets.
- Each referenced markdown file should be also **less than 150 lines**, remember that you can always split them into multiple files (**progressive disclosure** principle).
- Descriptions in metadata of `SKILL.md` files should be both concise (**less than 200 characters**) and still contains enough usecases of the references and assets, this will help skills can be activated automatically during the implementation process of Claude Code.
- **Referenced markdowns**:
  - Prefer concise wording, but never sacrifice precision, scope, or exceptions for brevity.
  - Can reference other markdown files or assets as needed.

**IMPORTANT:**
- Always keep in mind that `SKILL.md` and reference files should be token consumption efficient, so that **progressive disclosure** can be leveraged at best.
- `SKILL.md` should be **less than 150 lines**
- Referenced markdown files should be also **less than 150 lines**, remember that you can always split them into multiple files (**progressive disclosure** principle).

**Why?**
Better **context engineering**: leverage **progressive disclosure** technique of Agent Skills, when agent skills are activated, Claude Code will consider to load only relevant files into the context, instead of reading all long `SKILL.md` as before.

#### SKILL.md (required)

**File name:** `SKILL.md` (uppercase)
**File size:** Under 150 lines, if you need more, split it to multiple files (<150 lines each) in `references` folder.
`SKILL.md` is always short and concise, straight to the point, treat it as a quick reference guide.

**Metadata Quality:** The `name` and `description` (**MUST be under 200 characters**) in YAML frontmatter determine when Claude will use the skill. Be specific about what the skill does and when to use it, DO NOT sound generic, vague or educational. Use the third-person (e.g. "This skill should be used when..." instead of "Use this skill when...").

#### Instruction Precision Gate

Treat every sentence in `SKILL.md` as an operating instruction, not background prose.

Before keeping any sentence, check:
- Does this sentence make the agent choose the correct action?
- Is it true for every case covered by the skill?
- If not always true, is the exception explicit?
- Does it accidentally forbid valid user-requested work?
- Does it accidentally authorize risky work?
- Does it conflict with repo-first behavior, user intent, or another rule?

Remove or rewrite any sentence that is only mostly correct. A single wrong instruction can make the whole skill behave incorrectly.

#### Bundled Resources (optional)

##### References (`references/`)

Documentation and reference material intended to be loaded as needed into context to inform Claude's process and thinking.

- **When to include**: For documentation that Claude should reference while working
- **Examples**: `references/finance.md` for financial schemas, `references/mnda.md` for company NDA template, `references/policies.md` for company policies, `references/api_docs.md` for API specifications
- **Use cases**: Database schemas, best practices, common workflows, cheatsheets, tool instructions, API documentation, domain knowledge, company policies, detailed workflow guides
- **Benefits**: Keeps SKILL.md lean, loaded only when Claude determines it's needed, makes information discoverable without hogging the context window.
- **Best practice**: If files are large (>150 lines), split them into multiple files (<150 lines each) in `references` folder, include grep search patterns in `SKILL.md`
- **Avoid duplication**: Information should live in either `SKILL.md` or `references` files, not both. Prefer `references` files for detailed information unless it's truly core to the skill—this keeps `SKILL.md` lean while making information discoverable without hogging the context window. Keep only essential procedural instructions and workflow guidance in `SKILL.md`; move detailed reference material, schemas, and examples to `references` files.

**IMPORTANT:**
- Referenced markdown files should be also **less than 150 lines**, remember that you can always split them into multiple files (**progressive disclosure** principle).
- Referenced markdown files are practical instructions for Claude Code to use the tools, packages, plugins or APIs to achieve the tasks.
- Each skill teaches Claude how to perform a specific development task, not what a tool does.

##### Assets (`assets/`)

Files not intended to be loaded into context, but rather used within the output Claude produces.

- **When to include**: When the skill needs files that will be used in the final output
- **Examples**: `assets/logo.png` for brand assets, `assets/slides.pptx` for PowerPoint templates, `assets/frontend-template/` for HTML/React boilerplate, `assets/font.ttf` for typography
- **Use cases**: Templates, images, icons, boilerplate code, fonts, sample documents that get copied or modified
- **Benefits**: Separates output resources from documentation, enables Claude to use files without loading them into context

### Progressive Disclosure Design Principle

Skills use a three-level loading system to manage context efficiently:

1. **Metadata (name + description)** - Always in context (**less than 200 characters**)
2. **SKILL.md body** - When skill triggers (<5k words)
3. **Bundled resources** - References and assets loaded or used only when needed

## Skill Creation Process

To create a skill, follow the "Skill Creation Process" in order, skipping steps only if there is a clear reason why they are not applicable.

### Step 1: Understanding the Skill with Concrete Examples

Skip this step only when the skill's usage patterns are already clearly understood. It remains valuable even when working with an existing skill.

To create an effective skill, clearly understand concrete examples of how the skill will be used. This understanding can come from either direct user examples or generated examples that are validated with user feedback.

Use `AskUserQuestion` tool to gather user feedback and validate understanding.

For example, when building an image-editor skill, relevant questions include:

- "What functionality should the image-editor skill support? Editing, rotating, anything else?"
- "Can you give some examples of how this skill would be used?"
- "I can imagine users asking for things like 'Remove the red-eye from this image' or 'Rotate this image'. Are there other ways you imagine this skill being used?"
- "What would a user say that should trigger this skill?"

To avoid overwhelming users, avoid asking too many questions in a single message. Start with the most important questions and follow up as needed for better effectiveness.

Conclude this step when there is a clear sense of the functionality the skill should support.

### Step 2: Research on The Internet

An effective skill is a subset of real-life workflows from professional workflows and case studies, so it's important to research on the internet to understand the current state of the art and best practices.

Activate `/docs-seeker` skill to search for documentation if needed.

If you receive a lot of URLs or files, use multiple `WebFetch` tools and `Explore` subagents to explore them in parallel, then report back to main agent.

Activate `/research` skill to research for:
- Best practices & industry standards
- Existing CLI tools (executable via `npx`, `bunx` or `pipx`) and their usage patterns
- Workflows & success case studies
- Common patterns, use cases & examples
- Edge cases, potential pitfalls & avoidance strategies

Write down the reports from the research to be used in the next step.

### Step 3: Planning the Reusable Skill Contents

To turn concrete examples into an effective skill, analyze each example by:

1. Considering how to execute on the example from scratch
2. Prefer existing CLI tools' execution (via `npx`, `bunx` or `pipx`) over writing custom code
3. Identifying what references and assets would be helpful when executing these workflows repeatedly
4. Analyze the current skills catalog to avoid duplicating functionality, re-using existing skills where possible.

**Example:** When designing a `frontend-webapp-builder` skill for queries like "Build me a todo app" or "Build me a dashboard to track my steps," the analysis shows:

1. Writing a frontend webapp requires the same boilerplate HTML/React each time
2. An `assets/hello-world/` template containing the boilerplate HTML/React project files would be helpful to store in the skill

**Example:** When building a `big-query` skill to handle queries like "How many users have logged in today?" the analysis shows:

1. Querying BigQuery requires re-discovering the table schemas and relationships each time
2. A `references/schema.md` file documenting the table schemas would be helpful to store in the skill

To establish the skill's contents, analyze each concrete example to create a list of the reusable resources to include: references and assets.

### Step 4: Create or Update the Skill Files

At this point, create the skill folder and write the smallest useful file set.

Use `references/skill-md-template.md` as the starting point for `SKILL.md`.

Only create `references/` or `assets/` when the planned skill contents require them. Do not create placeholder files or example directories that are not needed by the final skill.

### Step 5: Edit the Skill

When editing the (newly-generated or existing) skill, remember that the skill is being created for another instance of Claude to use. Focus on including information that would be beneficial and non-obvious to Claude. Consider what procedural knowledge, domain-specific details, or reusable assets would help another Claude instance execute these tasks more effectively.

#### Start with Reusable Skill Contents

To begin implementation, start with the reusable resources identified above: `references/` and `assets/` files. Note that this step may require user input. For example, when implementing a `brand-guidelines` skill, the user may need to provide brand assets or templates to store in `assets/`, or documentation to store in `references/`.

Do not keep unused placeholder files. Every bundled resource should have a clear job in the skill workflow.

#### Update SKILL.md

**Writing Style:** Write the entire skill using **imperative/infinitive form** (verb-first instructions), not second person. Use objective, instructional language (e.g., "To accomplish X, do Y" rather than "You should do X" or "If you need to do X"). This maintains consistency and clarity for AI consumption.

To complete SKILL.md, answer the following questions:

1. What is the purpose of the skill, in a few sentences?
2. When should the skill be used?
3. In practice, how should Claude use the skill? All reusable skill contents developed above should be referenced so that Claude knows how to use them.

### Step 6: Validate and Package

Validate the skill manually against the criteria below before distribution. Do not treat a shallow structural check as proof that the skill is good.

If a distributable archive is needed, create it only after manual validation passes. The archive should include the skill folder contents and preserve relative paths.

### Step 7: Iterate

After testing the skill, users may request improvements. Often this happens right after using the skill, with fresh context of how the skill performed.

**Iteration workflow:**
1. Use the skill on real tasks
2. Notice struggles or inefficiencies
3. Notice token usage and performance
4. Identify how SKILL.md or bundled resources should be updated
5. Implement changes and test again

## Validation Criteria

Detailed validation criteria for evaluating skills:

- **Acceptance checklist**: `references/skill-acceptance-checklist.md`
- **Activation metadata**: `references/skill-activation-metadata.md`
- **Progressive disclosure and token budget**: `references/progressive-disclosure-token-budget.md`
- **Folder structure**: `references/skill-folder-structure.md`

## Plugin Marketplaces

For distributing skills as plugins via marketplaces, see:
- **Setup overview**: `references/plugin-marketplace-setup-overview.md`
- **JSON schema**: `references/plugin-marketplace-json-schema.md`
- **Source types**: `references/plugin-marketplace-source-types.md`
- **Hosting and distribution**: `references/plugin-marketplace-hosting-distribution.md`
- **Errors and fixes**: `references/plugin-marketplace-errors-and-fixes.md`

## References
- [Agent Skills](https://docs.claude.com/en/docs/claude-code/skills.md)
- [Agent Skills Spec](../agent_skills_spec.md)
- [Agent Skills Overview](https://docs.claude.com/en/docs/agents-and-tools/agent-skills/overview.md)
- [Best Practices](https://docs.claude.com/en/docs/agents-and-tools/agent-skills/best-practices.md)
- [Plugin Marketplaces](https://code.claude.com/docs/en/plugin-marketplaces.md)
