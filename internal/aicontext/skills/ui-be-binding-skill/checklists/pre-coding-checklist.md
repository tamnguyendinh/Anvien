# UI-BE Binding Pre-Coding Checklist

Before coding, confirm:

- [ ] I read `docs/ui-authority.md`.
- [ ] I read `docs/ui-slot-map.md`.
- [ ] I read `docs/state-map.md`.
- [ ] I read `docs/visible-text.lock.json`.
- [ ] I identified the exact page(s) in scope.
- [ ] I identified the exact backend endpoint(s) in scope.
- [ ] I identified the exact approved UI slot(s) in scope.
- [ ] I created a binding table.
- [ ] I confirmed no unapproved state is needed.
- [ ] I confirmed no new visible text is needed.
- [ ] I confirmed backend fields not listed in the slot map will not be rendered.
- [ ] I confirmed the implementation will use DTO → adapter → view model → approved slot.
