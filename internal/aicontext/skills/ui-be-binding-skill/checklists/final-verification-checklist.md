# UI-BE Binding Final Verification Checklist

Before handoff, confirm:

- [ ] Full build passed.
- [ ] Unit/integration tests passed.
- [ ] Playwright tests passed.
- [ ] Actual wiring status was completed or updated for the target surface.
- [ ] Final surface classification and evidence IDs are recorded.
- [ ] Preserve-only files/symbols were not rewritten.
- [ ] Fake/prototype/demo leakage was removed or explicitly blocked by authority.
- [ ] Screenshot comparison passed or approved diff is documented.
- [ ] Visible text snapshot passed.
- [ ] Forbidden text guard passed.
- [ ] No unapproved visible text was added.
- [ ] No unapproved loading state was added.
- [ ] No unapproved empty state was added.
- [ ] No unapproved error state was added.
- [ ] No layout/spacing/color/typography drift was introduced.
- [ ] Header/footer remained unchanged unless explicitly approved.
- [ ] Backend-only/internal fields are not rendered.
- [ ] Final report includes changed files and evidence.
