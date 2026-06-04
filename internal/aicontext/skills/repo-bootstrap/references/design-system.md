# Design System Reference

> Đọc file này khi dự án có FE phức tạp cần design system riêng.

## Token Structure

```css
:root {
  /* Color */
  --color-primary: #...;
  --color-primary-hover: #...;
  --color-surface: #...;
  --color-surface-raised: #...;
  --color-text: #...;
  --color-text-muted: #...;
  --color-border: #...;
  --color-error: #...;
  --color-success: #...;

  /* Spacing (4px base) */
  --space-1: 4px;
  --space-2: 8px;
  --space-3: 12px;
  --space-4: 16px;
  --space-6: 24px;
  --space-8: 32px;
  --space-12: 48px;
  --space-16: 64px;

  /* Typography */
  --font-sans: ...;
  --font-mono: ...;
  --text-xs: 11px;
  --text-sm: 13px;
  --text-base: 15px;
  --text-lg: 18px;
  --text-xl: 22px;
  --text-2xl: 28px;
  --text-3xl: 36px;
  --leading-tight: 1.2;
  --leading-normal: 1.5;

  /* Border */
  --radius-sm: 4px;
  --radius-md: 8px;
  --radius-lg: 12px;
  --radius-full: 9999px;

  /* Shadow */
  --shadow-sm: 0 1px 2px rgba(0,0,0,.05);
  --shadow-md: 0 4px 12px rgba(0,0,0,.1);
  --shadow-lg: 0 8px 24px rgba(0,0,0,.15);

  /* Motion */
  --duration-fast: 100ms;
  --duration-normal: 200ms;
  --duration-slow: 350ms;
  --ease-default: cubic-bezier(0.4, 0, 0.2, 1);
}
```

## Component Inventory

### Atomic Components
| Component | Props cần định nghĩa | Variants |
|-----------|----------------------|----------|
| Button | size, variant, disabled, loading | primary / secondary / ghost / danger |
| Input | type, error, disabled, prefix/suffix | default / error / disabled |
| Badge | variant, size | success / warning / error / info |
| Avatar | src, fallback, size | sm / md / lg |
| Spinner | size, color | - |
| Tooltip | content, position | - |
| Checkbox | checked, indeterminate | - |
| Switch | checked | - |

### Composite Components
| Component | Gồm những atomic nào |
|-----------|----------------------|
| FormField | Label + Input + ErrorMessage |
| SearchBar | Input + Icon + clear Button |
| DataTable | Header + Row + Pagination |
| Modal | Overlay + Dialog + Header + Footer |
| Toast | Icon + Message + dismiss Button |
| Dropdown | Trigger + Menu + MenuItem |

## Responsive Breakpoints
```
mobile:  < 640px
tablet:  640px – 1024px
desktop: > 1024px
```

## Accessibility Requirements
- Mọi interactive element có `aria-label` hoặc visible label
- Focus ring: không bỏ outline, custom style bằng `focus-visible`
- Color contrast: minimum WCAG AA (4.5:1 text, 3:1 UI)
- Keyboard navigable: Tab order logic, Escape đóng modal/dropdown