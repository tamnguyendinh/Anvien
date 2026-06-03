# AI Context Skill Script Problems

Date: 2026-06-03

Context: Imported skills under `internal/aicontext/skills` now share the `SKILL.md` directory shape, but some skills include scripts, references, assets, nested skill folders, and generated artifacts.

## Problems

1. Copying a skill folder does not mean the agent knows how to use the scripts inside it.

2. Script commands written in `SKILL.md` can use the wrong working directory after installation.

3. Relative script paths such as `scripts/foo.py` are ambiguous because they may be relative to the skill directory, not the repository root.

4. The agent cannot infer which scripts are safe and which scripts have side effects.

5. The agent cannot know whether a script supports `--dry-run` or `--validate-only` unless the skill documents it.

6. The agent cannot know a script input/output contract when the skill does not describe it.

7. The agent cannot know required runtime or dependency setup when the skill does not document it.

8. Some skills are nested deeper than `skills/<name>/SKILL.md`.

9. Some parent skill folders and child skill folders both contain `SKILL.md`, which creates ambiguity.

10. Some skills are missing frontmatter fields such as `name` and `description`.

11. A skill frontmatter name can differ from the source folder name.

12. Different folders can define duplicate or confusingly similar skill names.

13. Some files inside skill folders are generated artifacts, not skill assets.

14. Some skill folders include large assets that can significantly increase install or package size.

15. Some test or development files are mixed into skill folders without clear runtime relevance.

16. Some scripts may require external permissions, network access, or package managers.

17. Some scripts can modify the current repository if run from the wrong working directory.

18. Some scripts can write output to an unintended default location.

19. Some references or assets are present but not linked from `SKILL.md`, so the agent may not discover them.

20. Some scripts are only named in `SKILL.md` without enough workflow detail for reliable use.
