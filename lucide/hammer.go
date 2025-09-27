package lucide

import (
	html "github.com/plainkit/html"
)

// Hammer creates a Hammer Lucide icon.
func Hammer(args ...html.SvgArg) html.Node {
	svgArgs := withLucideDefaults("lucide lucide-hammer", args)
	children := []html.SvgArg{
		html.Child(html.SvgPath(html.AD("m15 12-9.373 9.373a1 1 0 0 1-3.001-3L12 9"))),
		html.Child(html.SvgPath(html.AD("m18 15 4-4"))),
		html.Child(html.SvgPath(html.AD("m21.5 11.5-1.914-1.914A2 2 0 0 1 19 8.172v-.344a2 2 0 0 0-.586-1.414l-1.657-1.657A6 6 0 0 0 12.516 3H9l1.243 1.243A6 6 0 0 1 12 8.485V10l2 2h1.172a2 2 0 0 1 1.414.586L18.5 14.5"))),
	}
	return html.Svg(append(svgArgs, children...)...)
}
