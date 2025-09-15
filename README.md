# Plain Lucide Icons

Type‑safe Lucide SVG icons for Go, built on top of `github.com/plainkit/blox`. Each icon is a simple Go function that returns a `blox.Component`, so you compose icons directly in your views with zero runtime templating.

## Install

```
go get github.com/plainkit/blox
go get github.com/plainkit/lucide
```

## Quick Start

```go
package main

import (
    "fmt"
    x "github.com/plainkit/blox"
    "github.com/plainkit/lucide"
)

func main() {
    btn := x.Button(
        x.Class("inline-flex items-center gap-2 px-3 py-2 border rounded"),
        x.Child(lucide.Check(
            x.Class("size-4"),          // Tailwind size utility (sets width/height)
            lucide.StrokeWidth("1.75"), // Thinner stroke
        )),
        x.Child(x.Text("Save")),
    )

    fmt.Println(x.Render(btn))
}
```

## Icon API

- One function per icon: `lucide.Activity(...)`, `lucide.CirclePower(...)`, `lucide.AArrowDown(...)`, etc.
- Signature: `func <IconName>(args ...x.SvgArg) x.Component`
- Works anywhere a `blox.Component` is accepted (e.g., as `Child(...)`).

### Arguments you can pass

Use any `blox` SVG option as arguments:

- `x.Class("...")`: CSS classes (e.g., Tailwind utilities)
- `x.SvgWidth("24")`, `x.SvgHeight("24")`: size in px, `em`, etc.
- `x.Stroke("currentColor"|"#000"|"..." )`: stroke color
- `x.Fill("none"|"currentColor"|"..." )`: fill color
- `x.Role("img")`, `x.Aria("label", "...")`, `x.Title("...")`: accessibility

Convenience helpers in this package:

- `lucide.Size("24")`: sets both width and height
- `lucide.StrokeWidth("1.5")`: shorthand for `x.StrokeWidth(...)`

### Sensible defaults

Every icon renders with these defaults (you override any of them by passing args):

- `xmlns="http://www.w3.org/2000/svg"`
- `width="24"`, `height="24"`
- `viewBox="0 0 24 24"`
- `fill="none"`, `stroke="currentColor"`, `stroke-width="2"`
- `stroke-linecap="round"`, `stroke-linejoin="round"`
- `class="lucide lucide-<kebab-name>"`

Tip: color comes from `currentColor`, so setting CSS `color` on the element or parent will color the icon.

## Naming & Discovery

- Function names are PascalCase versions of Lucide names:
  - `circle-power` → `lucide.CirclePower`
  - `a-arrow-down` → `lucide.AArrowDown`
- Use IDE autocomplete on `lucide.` to browse all available icons.

## Common Recipes

- Icon only, decorative:

  ```go
  lucide.Heart(
      x.Class("size-5 text-rose-500"),
      x.Aria("hidden", "true"),
  )
  ```

- Icon with accessible label:

  ```go
  lucide.AlertTriangle(
      x.Class("size-4 text-amber-600"),
      x.Role("img"),
      x.Aria("label", "Warning"),
  )
  ```

- Precise sizing without CSS classes:

  ```go
  lucide.Activity(
      lucide.Size("16"),      // width=16, height=16
      lucide.StrokeWidth("1.5"),
  )
  ```

- In a button with text:
  ```go
  x.Button(
      x.Class("inline-flex items-center gap-2"),
      x.Child(lucide.LogIn(x.Class("size-4"))),
      x.Child(x.Text("Sign in")),
  )
  ```

## Styling Guidance

- Size: prefer CSS (`.size-4`, `w-4 h-4`, etc.) or `lucide.Size("...")`.
- Color: inherits from `currentColor`. Set a `color` class or pass `x.Stroke("...")`.
- Thickness: use `lucide.StrokeWidth("1.5")` or `x.StrokeWidth("...")`.

## Accessibility

- Decorative icons: `x.Aria("hidden", "true")`.
- Informative icons: set `x.Role("img")` and either `x.Aria("label", "...")` or `x.Title("...")`.

## Important note about classes

Passing any global attribute (including `x.Role`, `x.Aria`, etc.) counts as a class override and disables the default `class="lucide lucide-<name>"`. If you rely on those defaults, include them explicitly when you also pass globals:

```go
lucide.Activity(
    x.Class("lucide lucide-activity size-4 text-slate-600"),
    x.Role("img"), // any Global arg means you should set class yourself if you need it
)
```

If you don’t depend on the `lucide` classes for styling, you can ignore this.

## Server Rendering

`blox.Render(component)` returns the final HTML string. Icons compose the same as any other `blox` component and render efficiently without reflection or template parsing.

## What’s Included

- Hundreds of Lucide icons as first‑class Go functions
- Type‑safe, chainable SVG options via `blox`
- Small helpers: `Size`, `StrokeWidth`
- Linker only keeps icons you actually reference

## License

- Go code in this repository: MIT © 2025 Plain Contributors
- Icon assets and path data: Lucide (ISC); portions derived from Feather (MIT)

See `LICENSE` for full texts and attribution details.
