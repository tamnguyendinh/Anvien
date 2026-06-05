# Notes Decisions Log - 2026-06-05

## File Role Classification Gap Post-Review Fix
- Scope: close Supervisor reject `reports/Supervisor/rp_supervisor_260605_232104_by_gpt-5-codex_file-role-classification-gap-review.md`.
- Decision: implement the fix, not just rewrite evidence, because the recorded e2e command and graph File-node `fileGroup` property were real closure gaps.
- Report: `reports/coder/rp_coder_260605_234312_by_gpt-5-codex_file_group_post_review_fix.md`.
- Evidence: full build passed, version remained `1.2.5`, focused Go tests passed, Web unit tests passed, Playwright e2e passed with `webServer`, `file-context` and graph inspection show `fileGroup=backend_support_model_helper`, and `detect-changes` passed before commit.
- Benchmark: current `backend_support_model_helper` group count is `47` files and `2531` unresolved source sites in the current file projection; historical E11 closure count remains `42`.
- Commit: `e64aaab fix(file-group): close post-review graph and e2e gaps`.
