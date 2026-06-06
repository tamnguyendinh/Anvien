# UI-BE Binding Pre-Coding Checklist

Before coding, confirm:

- [ ] I identified `<active-plan-dir>` for this scope.
- [ ] I read `<active-plan-dir>/YYYY-MM-DD-ui-authority.md`.
- [ ] I read or created `<active-plan-dir>/YYYY-MM-DD-actual-wiring-status.md`.
- [ ] I read `<active-plan-dir>/YYYY-MM-DD-ui-slot-map.md`.
- [ ] I read `<active-plan-dir>/YYYY-MM-DD-state-map.md`.
- [ ] I read `<active-plan-dir>/YYYY-MM-DD-visible-text.lock.json`.
- [ ] I added the target surface to the actual wiring audit matrix.
- [ ] I classified the surface as `bound-correct`, `partial`, `wrong-fake-data`, `unbound`, `no-real-data`, or `blocked`.
- [ ] I identified the real production source for visible dynamic values.
- [ ] I searched for fake/demo/prototype/sample production display.
- [ ] I listed already-correct files/symbols as preserve-only, if any.
- [ ] I rewrote the implementation plan from the audit classification.
- [ ] I identified the exact page(s) in scope.
- [ ] I identified the exact backend endpoint(s) in scope.
- [ ] I identified the exact approved UI slot(s) in scope.
- [ ] I created a binding table.
- [ ] I confirmed no unapproved state is needed.
- [ ] I confirmed no new visible text is needed.
- [ ] I confirmed backend fields not listed in the slot map will not be rendered.
- [ ] I confirmed the implementation will use DTO → adapter → view model → approved slot.
