#!/usr/bin/env bash
set -euo pipefail
# Generate Go icon components from SVGs using the local `claude` CLI.
# - Iterates over SVGs in icons directory (default: ./icons)
# - For each SVG, builds a rich, repo-specific prompt and feeds it to `claude`
# - Waits until the corresponding .go file appears and has size > 0, or 30s timeout
# - Kills the claude process between files via `pkill -f "claude"`
# - Skips icons that already have a non-empty .go file in this directory

ICONS_DIR=${1:-icons}
TIMEOUT_SEC=${TIMEOUT_SEC:-30}

if ! command -v claude >/dev/null 2>&1; then
  echo "Error: 'claude' CLI not found in PATH" >&2
  exit 1
fi

if [[ ! -d "$ICONS_DIR" ]]; then
  echo "Error: icons directory not found: $ICONS_DIR" >&2
  exit 1
fi

# Static prompt header describing how to convert an SVG to a Go icon function
PROMPT_HEADER=$(cat << PROMPT
You are converting a user-provided SVG into a Go icon component in the repository `github.com/bloxui/lucide` using the `github.com/bloxui/blox` API. Follow these exact rules:

- Goal: Convert a user-provided SVG into a Go icon component within package lucide, using
  github.com/bloxui/blox APIs for SVG elements and attributes.
  - File: Name is the SVG’s kebab-case base name with .go suffix (e.g., a-arrow-down.go).
  - Package/import:
      - package lucide
      - import x "github.com/bloxui/blox"
  - Function:
      - Name: PascalCase of the kebab name, exported (e.g., a-arrow-down → AArrowDown).
      - Doc: // <PascalName> creates a <Title Case> Lucide icon.
      - Signature: func <PascalName>(args ...x.SvgArg) x.Component
  - Root defaults: Use buildLucideArgs("lucide lucide-<kebab-name>", args) to set and allow
  overrides:
      - xmlns="http://www.w3.org/2000/svg"
      - width="24", height="24"
      - viewBox="0 0 24 24"
      - fill="none", stroke="currentColor", stroke-width="2"
      - stroke-linecap="round", stroke-linejoin="round"
      - class="lucide lucide-<kebab-name>"
  - Children translation: Convert each child element inside <svg> to x.Child(...) in order.
  Append via a single append call.
      - Return x.Svg(svgArgs...).

  Global Attributes

  - Use on any element: x.Class(v), x.Id(v), x.Title(v), x.Role(v), x.Lang(v),
  x.Dir(v), x.Slot(v), x.Part(v), x.Popover(v), x.Nonce(v), x.IsAttr(v), x.AccessKey(v),
  x.ContentEditable(v), x.InputMode(v), x.EnterKeyHint(v), x.ExportParts(v),
  x.ItemScope(b), x.ItemType(v), x.ItemId(v), x.ItemProp(v), x.ItemRef(v), x.XMLLang(v),
  x.XMLBase(v), x.VirtualKeyboardPolicy(v), x.Hidden(), x.Inert(), x.Autofocus(),
  x.TabIndex(n), x.Spellcheck(b), x.Translate(b), x.Draggable(b), x.WritingSuggestions(b),
  x.StyleKV(k, v), x.Aria(k, v), x.Data(k, v), x.On(ev, handler).

  Root <svg>

  - Constructor: x.Svg(...) via svgArgs built with buildLucideArgs.
  - Root attribute helpers (usable at root or via buildLucideArgs defaults):
      - Size/position: x.SvgWidth(v), x.SvgHeight(v), x.ViewBox(v),
  x.PreserveAspectRatio(v), x.SvgX(v), x.SvgY(v)
      - Namespacing/meta: x.Xmlns(v), x.Version(v), x.BaseProfile(v),
  x.ContentScriptType(v), x.ContentStyleType(v), x.ZoomAndPan(v)
      - Painting/transform/opacity: x.Fill(v), x.FillOpacity(v), x.FillRule(v),
  x.Stroke(v), x.StrokeWidth(v), x.StrokeDasharray(v), x.StrokeDashoffset(v),
  x.StrokeLinecap(v), x.StrokeLinejoin(v), x.StrokeOpacity(v), x.StrokeMiterlimit(v),
  x.Transform(v), x.Opacity(v)

  Element Mappings

  - path:
      - x.Path(x.D(v)[, paint/transform/opacity...])
      - Extra: x.PathLength(v) for pathLength
      - Common: x.Fill, x.FillOpacity, x.FillRule, x.Stroke, x.StrokeWidth,
  x.StrokeDasharray, x.StrokeDashoffset, x.StrokeLinecap, x.StrokeLinejoin,
  x.StrokeOpacity, x.StrokeMiterlimit, x.Transform, x.Opacity
  - line:
      - x.Line(x.X1(v), x.Y1(v), x.X2(v), x.Y2(v)[, PathLength/paint/transform/opacity...])
      - Extra: x.PathLength(v)
      - Common painting/transform/opacity as above
  - rect:
      - x.Rect(x.X(v), x.Y(v), x.RectWidth(v), x.RectHeight(v)[, x.Rx(v)][, x.Ry(v)][,
  PathLength/paint/transform/opacity...])
      - Notes: For rect, use x.X/x.Y; for root <svg> it’s x.SvgX/x.SvgY.
      - Extra: x.PathLength(v)
      - Common painting/transform/opacity as above
  - circle:
      - x.Circle(x.Cx(v), x.Cy(v), x.R(v)[, PathLength/paint/transform/opacity...])
      - Extra: x.PathLength(v)
      - Common painting/transform/opacity as above
  - ellipse:
      - x.Ellipse(x.EllipseCx(v), x.EllipseCy(v), x.EllipseRx(v), x.EllipseRy(v)[,
  PathLength/paint/transform/opacity...])
      - Extra: x.PathLength(v)
      - Common painting/transform/opacity as above
  - polyline:
      - x.Polyline(x.Points(v)[, PathLength/paint/transform/opacity...]) where v is "x1,y1
  x2,y2 ..."
      - Extra: x.PathLength(v)
      - Common painting/transform/opacity as above
  - polygon:
      - x.Polygon(x.Points(v)[, PathLength/paint/transform/opacity...])
      - Extra: x.PathLength(v)
      - Common painting/transform/opacity as above
  - g (group):
      - x.G([paint/transform/opacity/global...], x.Child(...), ...)
      - Common painting/transform/opacity as above
  - text (if present; uncommon in icons):
      - x.SvgText([x.SvgTextX(v)][, x.SvgTextY(v)][, x.Dx(v)][, x.Dy(v)][, x.Rotate(v)]
  [, x.TextLength(v)][, x.LengthAdjust(v)][, x.DominantBaseline(v)][, x.TextAnchor(v)][,
  x.FontFamily(v)][, x.FontSize(v)][, x.FontWeight(v)][, x.FontStyle(v)][, paint/transform/
  opacity...], x.Text("..."))

  Attribute Handling Rules

  - Preserve per-element attributes and map to the exact helpers above.
  - Omit attributes already set by buildLucideArgs (xmlns, width/height, viewBox,
  fill=none, stroke=currentColor, stroke-width=2, linecap=round, linejoin=round, class)
  unless overriding is intended.
  - Keep numeric/length values as strings (e.g., "24", "0 0 24 24", "1.5", "2", "12",
  "12.5").
  - Maintain path d and points exactly (remove trivial whitespace only).
  - Unsupported attributes (e.g., clip-path, mask, marker-*, filter, vector-effect,
  shape-rendering, etc.) are not directly supported by helpers:
      - If crucial, approximate via x.StyleKV("clip-path", "url(#id)") on the element.
      - Otherwise omit. Do not invent unknown helper names.

  Structure and Style

  - Build svgArgs := buildLucideArgs("lucide lucide-<kebab>", args).
  - Append children in a single append with trailing commas:
      - svgArgs = append(svgArgs, x.Child(x.Path(...)), x.Child(x.Line(...)), ...)
  - Return x.Svg(svgArgs...).
  - Naming:
      - File: <kebab>.go
      - Function: <Pascal>(args ...x.SvgArg) x.Component
      - Doc: // <Pascal> creates a <Title Case> Lucide icon

  Validation Checklist

  - File/function/doc/class names match the SVG input.
  - Uses buildLucideArgs("lucide lucide-<kebab>", args).
  - All SVG child elements are represented in order.
  - No duplicate defaults; overrides allowed via user args.
  - Compiles with go build in module github.com/bloxui/lucide.

Do not create tasks, do not invetigate current codebase, do not build the code after completion
PROMPT
)

shopt -s nullglob
found_any=false
for svg in "$ICONS_DIR"/*.svg; do
  found_any=true
  base=$(basename "$svg" .svg)
  out_file="${base}.go"

  if [[ -s "$out_file" ]]; then
    echo "[skip] $out_file already exists and is non-empty"
    continue
  fi

  tmpfile=$(mktemp 2>/dev/null || mktemp -t lucidegen)
  trap 'rm -f "$tmpfile" >/dev/null 2>&1 || true' EXIT

  # Read SVG content
  svg_content=$(cat "$svg")

  # Build the full prompt for this icon
  {
    printf '%s\n\n' "$PROMPT_HEADER"
    printf 'SVG filename: %s\n' "$base.svg"
    printf 'Output Go filename: %s\n' "$out_file"
    printf 'SVG content:\n%s\n' "$svg_content"
  } >"$tmpfile"

  # Collapse to a single line and escape quotes
  prompt_oneline=$(tr -d '\r\n' <"$tmpfile" | sed 's/"/\\"/g')

  echo "[gen ] $svg -> $out_file"

  # Attach to the controlling TTY. This keeps a TTY available so Ink
  # won't attempt raw mode on a non-TTY. Note: reading from TTY in background
  # may still be restricted by the shell, but claude shouldn't need to read input
  # for our non-interactive usage.
  claude --dangerously-skip-permissions "$prompt_oneline" </dev/tty >/dev/tty 2>&1 &
  claude_pid=$!

  # Wait until the .go file appears and is non-empty or timeout
  start_ts=$(date +%s)
  generated=0
  while :; do
    if [[ -s "$out_file" ]]; then
      generated=1
      break
    fi
    now_ts=$(date +%s)
    if (( now_ts - start_ts >= TIMEOUT_SEC )); then
      break
    fi
    sleep 1
  done

  # Kill claude process(es) between files as requested
  if ps -p "$claude_pid" >/dev/null 2>&1; then
    pkill -f "claude" >/dev/null 2>&1 || true
    # Give it a brief moment to terminate
    sleep 0.5
  fi

  if [[ "$generated" -eq 1 ]]; then
    echo "[ok  ] generated $out_file"
  else
    echo "[warn] timed out waiting for $out_file (>${TIMEOUT_SEC}s)"
  fi

  # Cleanup tmp between icons (trap also cleans on exit)
  rm -f "$tmpfile" || true
  trap - EXIT
done

if [[ "$found_any" = false ]]; then
  echo "No SVG files found in $ICONS_DIR" >&2
fi
