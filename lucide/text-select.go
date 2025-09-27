package lucide

import (
	html "github.com/plainkit/html"
)

// TextSelect creates a Text Select Lucide icon.
func TextSelect(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-text-select", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("M14 21h1"))),
		html.Child(html.SvgPath(html.AD("M14 3h1"))),
		html.Child(html.SvgPath(html.AD("M19 3a2 2 0 0 1 2 2"))),
		html.Child(html.SvgPath(html.AD("M21 14v1"))),
		html.Child(html.SvgPath(html.AD("M21 19a2 2 0 0 1-2 2"))),
		html.Child(html.SvgPath(html.AD("M21 9v1"))),
		html.Child(html.SvgPath(html.AD("M3 14v1"))),
		html.Child(html.SvgPath(html.AD("M3 9v1"))),
		html.Child(html.SvgPath(html.AD("M5 21a2 2 0 0 1-2-2"))),
		html.Child(html.SvgPath(html.AD("M5 3a2 2 0 0 0-2 2"))),
		html.Child(html.SvgPath(html.AD("M7 12h10"))),
		html.Child(html.SvgPath(html.AD("M7 16h6"))),
		html.Child(html.SvgPath(html.AD("M7 8h8"))),
		html.Child(html.SvgPath(html.AD("M9 21h1"))),
		html.Child(html.SvgPath(html.AD("M9 3h1"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
