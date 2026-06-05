# Plan

Title: AI Context Rule 1 Help Wording
Date: 2026-06-06
Status: Completed

## Goal

Change generated `# AGENTS Rules` rule 1 to:

```text
1. How to use anvien: run command "anvien --help".
```

## Scope

- `internal/aicontext/aicontext.go`
- `internal/aicontext/aicontext_test.go`
- Generated managed output for `AGENTS.md` and `CLAUDE.md`

## Checklist

- [x] P0-A: Run Anvien impact for the generator symbol before editing.
- [x] P1-A: Change generated rule 1 wording only.
- [x] P2-A: Update test assertion after code behavior is changed.
- [x] P3-A: Run full build, targeted tests, detect-changes, record evidence, and commit.
